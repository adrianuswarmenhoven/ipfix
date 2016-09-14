package ipfixmessage

import (
	"fmt"
)

/*

Template Record - A Template Record defines the structure and interpretation of fields in a Data Record.
Options Template Record - An Options Template Record is a Template Record that defines the structure and interpretation of fields in a Data Record, including defining how to scope the applicability of the Data Record.
Data Record - A Data Record is a record that contains values of the parameters corresponding to a Template Record.

One of the essential elements in the IPFIX record format is the Template Record.
Templates greatly enhance the flexibility of the record format because they allow the Collecting Process to process IPFIX Messages without necessarily knowing the interpretation of all Data Records.
A Template Record contains any combination of IANA-assigned and/or enterprise-specific Information Element identifiers.

The format of the Template Record consists of a Template Record Header and one or more Field Specifiers.

The scope, which is only available in the Options Template Set, gives the context of the reported Information Elements in the Data Records.

The scope is one or more Information Elements, specified in the Options Template Record.  At a minimum, Collecting Processes SHOULD support as scope the observationDomainId, exportingProcessId, meteringProcessId, templateId, lineCardId, exporterIPv4Address, exporterIPv6Address, and ingressInterface Information Elements.
The IPFIX protocol doesn't prevent the use of any Information Elements for scope.  However, some Information Element types don't make sense if specified as scope (for example, the counter Information Elements).

Multiple Scope Fields MAY be present in the Options Template Record, in which case the composite scope is the combination of the scopes.
For example, if the two scopes are meteringProcessId and templateId, the combined scope is this Template for this Metering Process.
If a different order of Scope Fields would result in a Record having a different semantic meaning, then the order of Scope Fields MUST be preserved by the Exporting Process.  For example, in the context of PSAMP [RFC5476], if the first scope defines the filtering function, while the second scope defines the sampling function, the order of the scope is important.
Applying the sampling function first, followed by the filtering function, would lead to potentially different Data Records than applying the filtering function first, followed by the sampling function.

*/

// TemplateRecord defines the structure of a (Options) Template Record
// If no Scope Fields are present then it is a regular Template Record
type TemplateRecord struct {
	Header               TemplateRecordHeader
	ScopeFieldSpecifiers []FieldSpecifier
	FieldSpecifiers      []FieldSpecifier
}

// TemplateRecordHeader defines the structure of a (Options) Template Record Header
type TemplateRecordHeader struct {
	TemplateID      uint16 //Each of the newly generated Template Records is given a unique Template ID.  This uniqueness is local to the Transport Session and Observation Domain that generated the Template ID. Template IDs 0-255 are reserved for Template Sets, Options Template Sets, and other reserved Sets yet to be created.  Template IDs of Data Sets are numbered from 256 to 65535.  There are no constraints regarding the order of the Template ID allocation.
	ScopeFieldCount uint16 //Number of scope fields in ths template record. If 0 then this is a regular template
	FieldCount      uint16 //Number of fields in this Template Record.
}

// NewTemplateRecord returns a new *TemplateRecord that has the given templateid and 0 Field Specifiers.
// Note that templateid *must* be between a number in the range 256-65535, else an error is returned.
func NewTemplateRecord(templateid uint16) (*TemplateRecord, error) {
	if templateid < 256 {
		return nil, fmt.Errorf("Invalid template id. Must be >=256 but got %d", templateid)
	}
	return &TemplateRecord{
		Header: TemplateRecordHeader{
			TemplateID:      templateid,
			ScopeFieldCount: 0,
			FieldCount:      0,
		},
		ScopeFieldSpecifiers: make([]FieldSpecifier, 0, 0),
		FieldSpecifiers:      make([]FieldSpecifier, 0, 0),
	}, nil
}

//AddScopeSpecifier adds a Scope Field Specifier to the record
func (tmplrec *TemplateRecord) AddScopeSpecifier(fsp FieldSpecifier) *TemplateRecord {
	tmplrec.FieldSpecifiers = append(tmplrec.FieldSpecifiers, fsp)
	tmplrec.Header.ScopeFieldCount++
	return tmplrec
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
	for _, rec := range tmplrec.ScopeFieldSpecifiers {
		reclen += rec.Len()
	}
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
