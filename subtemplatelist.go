package ipfixmessage

import "fmt"

// SubTemplateList represents a list of zero or more instances of a structured data type, where the data type of each list element is the same and corresponds with a single Template Record.
// Examples include a structured data type composed of multiple pairs of ("MPLS label stack entry position", "MPLS label stack value"), a structured data type composed of performance metrics, and a structured data type composed of multiple pairs of IP address, etc.
type SubTemplateList struct {
	Semantic    uint8        //one of: SemanticsNoneOf, ExactlyOneOf, OneOrMoreOf, AllOf, Ordered or Undefined
	TemplateID  uint16       //Each of the newly generated Template Records is given a unique Template ID.  This uniqueness is local to the Transport Session and Observation Domain that generated the Template ID. Template IDs 0-255 are reserved for Template Sets, Options Template Sets, and other reserved Sets yet to be created.  Template IDs of Data Sets are numbered from 256 to 65535.  There are no constraints regarding the order of the Template ID allocation.
	FieldValues []FieldValue //The list of Field Values
}

// NewSubTemplateList returns a new SubTemplateList.
func NewSubTemplateList(semantic uint8, templateid uint16) (*SubTemplateList, error) {
	if templateid < 256 {
		return nil, fmt.Errorf("Can not have a template id <256, but got %d", templateid)
	}
	if semantic >= 0x05 && semantic <= 0xFE {
		return nil, fmt.Errorf("Semantic undefined: %d", semantic)
	}
	return &SubTemplateList{
		Semantic:    semantic,
		TemplateID:  templateid,
		FieldValues: make([]FieldValue, 0, 0),
	}, nil
}

// Len returns the length of the field specifier, in octets.
func (stl *SubTemplateList) Len() uint16 {
	stllen := uint16(3)
	for _, listitem := range stl.FieldValues {
		stllen += listitem.Len()
	}
	return stllen
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (stl *SubTemplateList) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (stl *SubTemplateList) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("Can not unmarshal, invalid data. %#v", data)
	}

	return fmt.Errorf("Not yet implemented!")
}
