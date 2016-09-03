package ipfixmessage

import (
	"fmt"
)

//Creates a new IPFIX Set with specified set id
func NewIPFIXSet(setid int) (*IPFIXSet, error) {

	//No code should ever produce an invalid set ID
	if setid < SetID_Template {
		return nil, fmt.Errorf("Invalid value for Set ID: %d", setid)
	}

	return &IPFIXSet{
		SetHeader: IPFIXSetHeader{
			SetID:  uint16(setid),
			Length: uint16(0),
		},
		Records: []*IPFIXRecord{},
	}, nil
}

//Finalizes the Set and calculates it's length
func (ipfixset *IPFIXSet) Finalize() error {

	//Calculate the length of the message
	ipfixset.SetHeader.Length = ipfixSetHeaderLength
	for _, rec := range ipfixset.Records {
		ipfixset.SetHeader.Length += (*rec).Size()
	}

	return nil
}
