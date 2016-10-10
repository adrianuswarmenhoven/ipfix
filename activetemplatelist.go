package ipfixmessage

import (
	"fmt"
	"sync"
	"time"
)

// ActiveTemplates is a list of currently active templates.
// These can be used in a session or when testing the marshalling/unmarshalling of the complex types
type ActiveTemplates struct {
	Template map[uint16]*ActiveTemplate

	sync.Mutex
}

//ActiveTemplate is the structure that holds the data for a template record
type ActiveTemplate struct {
	Record *TemplateRecord

	Added        time.Time //So we do not remove it if it just has been very recently added
	LastAccessed time.Time //To implement clean-up routine
	NofAccess    uint64    //Counting the number of times this template is accessed
}

//NewActiveTemplateList returns a new empty template list
func NewActiveTemplateList() *ActiveTemplates {
	return &ActiveTemplates{Template: make(map[uint16]*ActiveTemplate)}
}

//Set adds or replaces a template in the list
func (at *ActiveTemplates) Set(id uint16, tpl *TemplateRecord) error {
	if id < 256 {
		return fmt.Errorf("Invalid template id. Must be >=256 but got %d", id)
	}
	if tpl == nil {
		return fmt.Errorf("Got nil pointer to template")
	}
	at.Lock()
	defer at.Unlock()

	if tmpl, found := at.Template[id]; found {
		isequal := true
		for idx, fsp := range tpl.FieldSpecifiers {
			if fsp.EnterpriseNumber != tmpl.Record.FieldSpecifiers[idx].EnterpriseNumber ||
				fsp.InformationElementIdentifier != tmpl.Record.FieldSpecifiers[idx].InformationElementIdentifier ||
				fsp.FieldLength != tmpl.Record.FieldSpecifiers[idx].FieldLength {
				isequal = false
			}
		}
		if !isequal {
			at.Template[id] = &ActiveTemplate{
				Record:       tpl,
				Added:        time.Now(),
				LastAccessed: time.Now(),
			}

		} else {
			tmpl.LastAccessed = time.Now()
		}
	} else {
		at.Template[id] = &ActiveTemplate{
			Record:       tpl,
			Added:        time.Now(),
			LastAccessed: time.Now(),
		}
	}
	return nil
}

//Get returns the template record for the id or an error if not found
func (at *ActiveTemplates) Get(id uint16) (*TemplateRecord, error) {
	if at == nil {
		return nil, fmt.Errorf("WTF")
	}
	at.Lock()
	defer at.Unlock()

	var tmpl *ActiveTemplate
	var found bool
	if tmpl, found = at.Template[id]; !found {
		return nil, fmt.Errorf("No such template (%d) in list.", id) //Not necessarily a fatal error.
	}
	tmpl.LastAccessed = time.Now()
	tmpl.NofAccess++
	return tmpl.Record, nil

}
