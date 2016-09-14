package ipfixmessage

import (
	"fmt"
)

/*
One of the essential elements in the IPFIX record format is the Template Record.
Templates greatly enhance the flexibility of the record format because they allow the Collecting Process to process IPFIX Messages without necessarily knowing the interpretation of all Data Records.
A Template Record contains any combination of IANA-assigned and/or enterprise-specific Information Element identifiers.

The format of the Template Record consists of a Template Record Header and one or more Field Specifiers.
*/

// TemplateRecord defines the structure of a Template Record
type TemplateRecord struct {
	Header          TemplateRecordHeader
	FieldSpecifiers []FieldSpecifier
}

// TemplateRecordHeader defines the structure of a Template Record Header
type TemplateRecordHeader struct {
	TemplateID uint16 //Each of the newly generated Template Records is given a unique Template ID.  This uniqueness is local to the Transport Session and Observation Domain that generated the Template ID. Template IDs 0-255 are reserved for Template Sets, Options Template Sets, and other reserved Sets yet to be created.  Template IDs of Data Sets are numbered from 256 to 65535.  There are no constraints regarding the order of the Template ID allocation.
	FieldCount uint16 //Number of fields in this Template Record.
}

// NewTemplateRecord returns a new *TemplateRecord that has the given templateid and 0 Field Specifiers.
// Note that templateid *must* be between a number in the range 256-65535, else an error is returned.
func NewTemplateRecord(templateid uint16) (*TemplateRecord, error) {
	if templateid < 256 {
		return nil, fmt.Errorf("Invalid template id. Must be >=256 but got %d", templateid)
	}
	return &TemplateRecord{
		Header: TemplateRecordHeader{
			TemplateID: templateid,
			FieldCount: 0,
		},
	}, nil
}

//AddSpecifier adds a Field Specifier to the record
func (tmplrec *TemplateRecord) AddSpecifier(fsp FieldSpecifier) *TemplateRecord {
	tmplrec.FieldSpecifiers = append(tmplrec.FieldSpecifiers, fsp)
	tmplrec.Header.FieldCount++
	return tmplrec
}

// Len returns the size in octets of the template record
func (tmplrec *TemplateRecord) Len() uint16 {
	reclen := uint16(4) //header is 4 bytes
	for _, rec := range tmplrec.FieldSpecifiers {
		reclen += rec.Len()
	}
	return reclen
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (tmplrec *TemplateRecord) MarshalBinary() (data []byte, err error) {
	if tmplrec.Header.FieldCount < 1 || len(tmplrec.FieldSpecifiers) < 1 {
		return nil, fmt.Errorf("Can not marshal record, must have at least one Field Specifier")
	}
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (tmplrec *TemplateRecord) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("Can not unmarshal, invalid data. %#v", data)
	}

	return fmt.Errorf("Not yet implemented!")
}
