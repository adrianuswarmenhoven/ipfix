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
	SetHeader SetHeader //Generic Set Header.
	Records   []*Record //Various types of record, depending on the SetID value (or rather, type of IPFIX set)
	Padding   []byte    //Optional padding bytes
}

/* N.B. per Padding:

The Exporting Process MAY insert some padding octets, so that the subsequent Set starts at an aligned boundary.
For security reasons, the padding octet(s) MUST be composed of octets with value zero (0).
The padding length MUST be shorter than any allowable record in this Set.
If padding of the IPFIX Message is desired in combination with very short records, then the padding Information Element 'paddingOctets' can be used for padding records such that their length is increased to a multiple of 4 or 8 octets.
Because Template Sets are always 4-octet aligned by definition, padding is only needed in the case of other alignments, e.g., on 8-octet boundaries.
*/

// SetHeader defines the Generic Set Header. Basically an ID and the size of the Set.
type SetHeader struct {
	SetID  uint16 //Set ID value identifies the Set.  A value of 2 is reserved for the Template Set.  A value of 3 is reserved for the Option Template Set.  All other values from 4 to 255 are reserved for future use. Values above 255 are used for Data Sets.
	Length uint16 //Total length of the Set, in octets, including the Set Header, all records, and the optional padding.
}

// NewSet creates a new IPFIX Set with specified set ID
func NewSet(setid uint16) (*Set, error) {
	if setid < SetIDTemplate || (setid > SetIDOptionTemplate && setid < 256) {
		return nil, fmt.Errorf("Invalid value for Set ID: %d", setid)
	}

	return &Set{
		SetHeader: SetHeader{
			SetID:  setid,
			Length: uint16(0),
		},
		Records: []*Record{},
	}, nil
}

// Finalize inalizes the Set and calculates it's length.
// The paddingboundary is the number of octets to align to, for example 8 for 8-octet boundaries.
// If paddingboundary is greater than 0 then padding will be added to fill the set to that boundary.
func (ipfixset *Set) Finalize(paddingboundary uint8) error {
	//Calculate the length of the message
	ipfixset.SetHeader.Length = ipfixSetHeaderLength
	for _, rec := range ipfixset.Records {
		ipfixset.SetHeader.Length += (*rec).Len()
	}
	//FIXME
	if paddingboundary == 0 {
		ipfixset.SetHeader.Length += 0
	}

	return fmt.Errorf("Not yet implemented properly!")
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (ipfixset *Set) MarshalBinary() (data []byte, err error) {
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
