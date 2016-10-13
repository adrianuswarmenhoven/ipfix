package ipfixmessage

import "fmt"

// SubTemplateMultiList represents a list of zero or more instances of a structured data type, where the data type of each list element can be different and corresponds with different Template definitions.
// Examples include a structured data type composed of multiple access-list entries, where entries can be composed of different criteria types.
type SubTemplateMultiList struct {
	Semantic            uint8              //one of: SemanticsNoneOf, ExactlyOneOf, OneOrMoreOf, AllOf, Ordered or Undefined
	AssociatedTemplates *ActiveTemplates   //We can only begin to marshal/unmarshal the records when we have the whole template belonging to the TemplateID
	SubTemplates        []*SubTemplateData //The main difference with a SubTemplateList is that this type only has the semantic once
}

// NewSubTemplateMultiList returns a new SubTemplateMultiList.
func NewSubTemplateMultiList(semantic uint8) (*SubTemplateMultiList, error) {
	if semantic >= 0x05 && semantic <= 0xFE {
		return nil, fmt.Errorf("Semantic undefined: %d", semantic)
	}
	return &SubTemplateMultiList{
		Semantic:     semantic,
		SubTemplates: make([]*SubTemplateData, 0, 0),
	}, nil
}

// Len returns the length of the field specifier, in octets.
func (stml *SubTemplateMultiList) Len() uint16 {
	stmllen := uint16(1) //Semantic
	for _, listitem := range stml.SubTemplates {
		stmllen += listitem.Len()
	}
	return stmllen
}

// SubTemplateData represents a list of zero or more instances of a structured data type, where the data type of each list element is the same and corresponds with a single Template Record.
type SubTemplateData struct {
	TemplateID          uint16           //Each of the newly generated Template Records is given a unique Template ID.  This uniqueness is local to the Transport Session and Observation Domain that generated the Template ID. Template IDs 0-255 are reserved for Template Sets, Options Template Sets, and other reserved Sets yet to be created.  Template IDs of Data Sets are numbered from 256 to 65535.  There are no constraints regarding the order of the Template ID allocation.
	AssociatedTemplates *ActiveTemplates //We can only begin to marshal/unmarshal the records when we have the whole template belonging to the TemplateID
	Records             []Record         //The list of Records
}

// NewSubTemplateData returns a new SubTemplateData record.
func NewSubTemplateData(templateid uint16) (*SubTemplateData, error) {
	if templateid < 256 {
		return nil, fmt.Errorf("Can not have a template id <256, but got %d", templateid)
	}
	return &SubTemplateData{
		TemplateID: templateid,
		Records:    make([]Record, 0, 0),
	}, nil
}

//Len returns the total length of the Data Records encoding for the Template ID previously specified, including the two bytes for the Template ID and the two bytes for the Data Records Length field itself.
func (stld SubTemplateData) Len() uint16 {
	stldlen := uint16(4) //two for id, two for length
	for _, listitem := range stld.Records {
		stldlen += listitem.Len()
	}
	return stldlen
}

// AssociateTemplates sets the template to be used marshalling/unmarshalling this SubTemplateData
func (stld *SubTemplateData) AssociateTemplates(at *ActiveTemplates) error {
	if at == nil {
		return fmt.Errorf("Can not use nil as Template List")
	}
	stld.AssociatedTemplates = at
	return nil
}
