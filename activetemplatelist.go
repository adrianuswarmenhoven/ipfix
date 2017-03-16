package ipfix

import (
	"fmt"
	"sync"
	"time"
)

// ActiveTemplates is a list of currently active templates.
// These can be used in a session or when testing the marshalling/unmarshalling of the complex types
type ActiveTemplates struct {
	templates map[uint16]*activeTemplate //Using a map here instead of an array for memory reasons (512K per session might be excessive otherwise)

	sync.Mutex
}

//activeTemplate is the structure that holds the data for a template record
type activeTemplate struct {
	Record *TemplateRecord

	Added        time.Time //So we do not remove it if it just has been very recently added
	LastAccessed time.Time //To implement clean-up routine
	NofAccess    uint64    //Counting the number of times this template is accessed
}

//NewActiveTemplateList returns a new empty templates list
func NewActiveTemplateList() *ActiveTemplates {
	return &ActiveTemplates{templates: make(map[uint16]*activeTemplate)}
}

//Set adds or replaces a templates in the list
func (at *ActiveTemplates) Set(id uint16, tpl *TemplateRecord) error {
	if id < 256 {
		return NewError(fmt.Sprintf("Invalid templates id. Must be >=256 but got %d", id), ErrCritical)
	}
	if tpl == nil {
		return NewError("Got nil pointer to templates", ErrCritical)
	}
	at.Lock()
	defer at.Unlock()

	if tmpl, found := at.templates[id]; found {
		isequal := true
		for idx, fsp := range tpl.FieldSpecifiers {
			if fsp.EnterpriseNumber != tmpl.Record.FieldSpecifiers[idx].EnterpriseNumber ||
				fsp.InformationElementIdentifier != tmpl.Record.FieldSpecifiers[idx].InformationElementIdentifier ||
				fsp.FieldLength != tmpl.Record.FieldSpecifiers[idx].FieldLength {
				isequal = false
			}
		}
		if !isequal {
			at.templates[id] = &activeTemplate{
				Record:       tpl,
				Added:        time.Now(),
				LastAccessed: time.Now(),
			}

		} else {
			tmpl.LastAccessed = time.Now()
		}
	} else {
		at.templates[id] = &activeTemplate{
			Record:       tpl,
			Added:        time.Now(),
			LastAccessed: time.Now(),
		}
	}
	return nil
}

//Get returns the templates record for the id or an error if not found
func (at *ActiveTemplates) Get(id uint16) (*TemplateRecord, error) {
	if at == nil {
		return nil, NewError("No active templates available", ErrCritical)
	}
	at.Lock()
	defer at.Unlock()

	var tmpl *activeTemplate
	var found bool
	if tmpl, found = at.templates[id]; !found {
		return nil, NewError(fmt.Sprintf("No such templates (%d) in list.", id), ErrFailure) //Not necessarily a fatal error. May hold back until we get a new one
	}
	tmpl.LastAccessed = time.Now()
	tmpl.NofAccess++
	return tmpl.Record, nil

}
