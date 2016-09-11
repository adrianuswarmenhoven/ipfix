package ipfixmessage

import (
	"fmt"
)

type Set struct {
	SetHeader SetHeader //Generic Set Header.
	Records   []*Record //Various types of record, depending on the SetID value (or rather, type of IPFIX set)
}

type SetHeader struct {
	SetID  uint16 //Set ID value identifies the Set.  A value of 2 is reserved for the Template Set.  A value of 3 is reserved for the Option Template Set.  All other values from 4 to 255 are reserved for future use. Values above 255 are used for Data Sets.
	Length uint16 //Total length of the Set, in octets, including the Set Header, all records, and the optional padding.
}

//Creates a new IPFIX Set with specified set id
func NewSet(setid int) (*Set, error) {

	//No code should ever produce an invalid set ID
	if setid < SetID_Template {
		return nil, fmt.Errorf("Invalid value for Set ID: %d", setid)
	}

	return &Set{
		SetHeader: SetHeader{
			SetID:  uint16(setid),
			Length: uint16(0),
		},
		Records: []*Record{},
	}, nil
}

//Finalizes the Set and calculates it's length
func (ipfixset *Set) Finalize() error {

	//Calculate the length of the message
	ipfixset.SetHeader.Length = ipfixSetHeaderLength
	for _, rec := range ipfixset.Records {
		ipfixset.SetHeader.Length += (*rec).Size()
	}

	return nil
}
