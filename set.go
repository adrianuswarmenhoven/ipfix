package ipfixmessage

import (
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
		return nil, fmt.Errorf("Invalid value for Set ID: %d", setid)
	}

	return &Set{
		SetID:   setid,
		Padding: uint16(0),
		Records: make([]*Record, 0, 0),
	}, nil
}

// Len returns the size in octets of the Set
func (ipfixset *Set) Len() uint16 {
	setlen := uint16(4) //We start out with 2 bytes for ID and 2 bytes for length
	for _, rec := range ipfixset.Records {
		setlen += (*rec).Len()
	}
	setlen += ipfixset.Padding
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
		return fmt.Errorf("Can not add record. Record size %d + Set Size %d > 65535", rec.Len(), ipfixset.Len())
	}
	ipfixset.Records = append(ipfixset.Records, &rec)
	return nil
}

// Pad calculates the padding bytes.
// The paddingboundary is the number of octets to align to, for example 8 for 8-octet boundaries.
// If the result is greater than 0 then padding will be added to fill the set to that boundary.
func (ipfixset *Set) Pad(paddingboundary uint16) {
	//Calculate the length of the message
	ipfixset.Padding = paddingboundary - (ipfixset.Len() % paddingboundary)
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (ipfixset *Set) MarshalBinary() (data []byte, err error) {
	//If template or optionstemplate do not use associate
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (ipfixset *Set) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("Can not unmarshal, invalid data. %#v", data)
	}

	return fmt.Errorf("Not yet implemented!")
}
