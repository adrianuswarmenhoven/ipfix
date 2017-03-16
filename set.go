package ipfixmessage

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
A Set is a generic term for a collection of records that have a similar structure.
There are three different types of Sets: Template Sets, Options Template Sets, and Data Sets.
Each of these Sets consists of a Set Header and one or more records.
*/

// Set defines the Generic Set type
type Set struct {
	SetID   uint16    //Set ID value identifies the Set.  A value of 2 is reserved for the Template Set.  A value of 3 is reserved for the Option Template Set.  All other values from 4 to 255 are reserved for future use. Values above 255 are used for Data Sets.
	Records []*Record //Various types of record, depending on the SetID value (or rather, type of IPFIX set)
	Padding uint16    //Optional padding bytes (only the number, the actual bytes will be added in encoding)

	//AssociatedTemplates Templates points to the list of active templates (whether in a session or not). Without a template record a data record can not be encoded or decoded
	AssociatedTemplates *ActiveTemplates
}

/* N.B. per Padding:

The Exporting Process MAY insert some padding octets, so that the subsequent Set starts at an aligned boundary.
For security reasons, the padding octet(s) MUST be composed of octets with value zero (0).
The padding length MUST be shorter than any allowable record in this Set.
If padding of the IPFIX Message is desired in combination with very short records, then the padding Information Element 'paddingOctets' can be used for padding records such that their length is increased to a multiple of 4 or 8 octets.
Because Template Sets are always 4-octet aligned by definition, padding is only needed in the case of other alignments, e.g., on 8-octet boundaries.
*/

// NewSet creates a new IPFIX Set with specified set ID
func NewSet(setid uint16) (*Set, error) {
	if setid < SetIDTemplate || (setid > SetIDOptionTemplate && setid < 256) {
		return nil, NewError(fmt.Sprintf("Invalid value for Set ID: %d", setid), ErrCritical)
	}

	return &Set{
		SetID:   setid,
		Padding: uint16(0),
		Records: make([]*Record, 0, 0),
	}, nil
}

// NewBlankSet creates a new IPFIX Set for unmarshalling
func NewBlankSet() *Set {
	return &Set{
		SetID:   uint16(0),
		Padding: uint16(0),
		Records: make([]*Record, 0, 0),
	}
}

// AssociateTemplates sets the template to be used marshalling/unmarshalling this DataRecord
func (ipfixset *Set) AssociateTemplates(at *ActiveTemplates) error {
	if at == nil {
		return NewError("Can not use nil as Template List", ErrCritical)
	}
	ipfixset.AssociatedTemplates = at
	return nil
}

// Len returns the size in octets of the Set
func (ipfixset *Set) Len() uint16 {
	setlen := uint16(4) //We start out with 2 bytes for ID and 2 bytes for length
	for _, rec := range ipfixset.Records {
		setlen += (*rec).Len()
	}
	if ipfixset.Padding > 0 {
		setlen += (ipfixset.Padding - (setlen % ipfixset.Padding)) % ipfixset.Padding
	}
	return setlen
}

// String returns the string representation of the Set
func (ipfixset *Set) String() string {
	retstring := fmt.Sprintf("set id=%d, ", ipfixset.SetID)
	if ipfixset.Padding == 0 {
		retstring += fmt.Sprintf("set length (without padding)=%d, \n", ipfixset.Len())
	} else {
		retstring += fmt.Sprintf("set length (padding=%d)=%d, \n", ipfixset.Padding, ipfixset.Len())
	}
	for _, rec := range ipfixset.Records {
		retstring += (*rec).String() + "\n"
	}
	return retstring
}

// AddRecord adds a new record to this set
func (ipfixset *Set) AddRecord(rec Record) error {
	if int(rec.Len())+int(ipfixset.Len()) > 65535 {
		return NewError(fmt.Sprintf("Can not add record. Record size %d + Set Size %d > 65535", rec.Len(), ipfixset.Len()), ErrCritical)
	}
	switch rec.(type) {
	case *TemplateRecord:
		switch ipfixset.SetID {
		case SetIDTemplate:
			if rec.(*TemplateRecord).ScopeFieldSpecifiers != nil {
				return NewError("Can not add Option Template Record to Template Set", ErrCritical)
			}

		case SetIDOptionTemplate:
			if rec.(*TemplateRecord).ScopeFieldSpecifiers == nil {
				return NewError("Can not add Template Record to Scope Field Set", ErrCritical)
			}
		default:
			return NewError("Can not add Template Record to Data Set", ErrCritical)
		}
	case *DataRecord:
		switch ipfixset.SetID {
		case SetIDTemplate, SetIDOptionTemplate:
			if rec.(*TemplateRecord).ScopeFieldSpecifiers == nil {
				return NewError("Can not add Data Record to (Scope) Field Set", ErrCritical)
			}
		}
	}
	ipfixset.Records = append(ipfixset.Records, &rec)
	return nil
}

// Pad calculates the padding bytes.
// The paddingboundary is the number of octets to align to, for example 8 for 8-octet boundaries.
// If the result is greater than 0 then padding will be added to fill the set to that boundary.
func (ipfixset *Set) Pad(paddingboundary uint16) {
	ipfixset.Padding = paddingboundary
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
func (ipfixset *Set) MarshalBinary() (data []byte, err error) {
	//Set ID value identifies the Set.  A value of 2 is reserved for the Template Set.  A value of 3 is reserved for the Option Template Set.
	//All other values from 4 to 255 are reserved for future use. Values above 255 are used for Data Sets.
	if ipfixset.SetID > 3 && ipfixset.SetID < 256 && ipfixset.AssociatedTemplates == nil {
		return nil, NewError(fmt.Sprintf("Need associated templates for Set ID %d", ipfixset.SetID), ErrCritical)
	}
	buf := new(bytes.Buffer) //should get from pool?

	//   Each Set Header field is exported in network format.  The fields are defined as follows:
	//   Set ID
	//   Length
	//      Total length of the Set, in octets, including the Set Header, all records, and the optional padding.
	//      Because an individual Set MAY contain multiple records, the Length value MUST be used to determine the position of the next Set.
	err = binary.Write(buf, binary.BigEndian, ipfixset.SetID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, ipfixset.Len())
	if err != nil {
		return nil, err
	}
	data = buf.Bytes()
	for _, rec := range ipfixset.Records {
		recdat, err := (*rec).MarshalBinary()
		if err != nil {
			return nil, err
		}
		data = append(data, recdat...)
	}
	if ipfixset.Padding > 0 {
		padlen := int(ipfixset.Len()) - len(data)
		data = append(data, make([]byte, padlen, padlen)...)
	}
	return data, nil
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
func (ipfixset *Set) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) < 4 {
		return NewError(fmt.Sprintf("Can not unmarshal, invalid data. %#v", data), ErrCritical)
	}

	ipfixset.SetID = binary.BigEndian.Uint16(data[0:2])
	recordlength := uint16(0)
	hasvar := false
	if ipfixset.SetID > 255 {
		if ipfixset.AssociatedTemplates == nil {
			return NewError(fmt.Sprintf("Must have associated templates to unmarshal set with ID %d", ipfixset.SetID), ErrCritical)
		}
		ipfixsetTemplate, err := ipfixset.AssociatedTemplates.Get(ipfixset.SetID)
		if err != nil {
			return err
		}
		for _, fsp := range ipfixsetTemplate.FieldSpecifiers {
			if fsp.FieldLength != VariableLength {
				recordlength += fsp.FieldLength
			} else {
				hasvar = true
				recordlength += 2 //one byte for length, one for value
			}
		}
	} else {
		recordlength = 4 //template header
	}
	datalength := binary.BigEndian.Uint16(data[2:4])
	cursor := uint16(4)

	for cursor < datalength { //We always need at least 4 bytes to determine Template ID and Field Count
		if ((cursor + recordlength) > datalength) ||
			((cursor+recordlength == datalength) && (bytes.Count(data[cursor:], []byte{0}) == int(recordlength))) { //Must be padding
			return nil
		}
		//Set ID value identifies the Set.  A value of 2 is reserved for the Template Set.  A value of 3 is reserved for the Option Template Set.
		//All other values from 4 to 255 are reserved for future use. Values above 255 are used for Data Sets.
		switch {
		case ipfixset.SetID == SetIDTemplate, ipfixset.SetID == SetIDOptionTemplate: //We do the template or option template set
			tmprec := &TemplateRecord{}
			if ipfixset.SetID == SetIDOptionTemplate {
				tmprec.ScopeFieldSpecifiers = make([]*FieldSpecifier, 0, 0)
			}
			err := tmprec.UnmarshalBinary(data[cursor:])
			if err != nil {
				fmt.Println("A", err)
				return err
			}
			cursor += tmprec.Len()
			ipfixset.AddRecord(tmprec)
		case ipfixset.SetID > 255: //We do a dataset
			tmprec, err := NewDataRecord(ipfixset.SetID, ipfixset.AssociatedTemplates)
			if err != nil {
				fmt.Println("B", err)
				return err
			}
			err = tmprec.UnmarshalBinary(data[cursor:])
			if err != nil {
				fmt.Println("C", err)
				return err
			}
			if !hasvar {
				cursor += recordlength
			} else {
				cursor += tmprec.Len()
			}
			ipfixset.AddRecord(tmprec)
		default: //Invalid Template ID
			return NewError(fmt.Sprintf("Invalid template ID: %d", ipfixset.SetID), ErrCritical)
		}
	}
	return nil
}
