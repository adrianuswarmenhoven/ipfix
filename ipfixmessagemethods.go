package ipfixmessage

import (
	"fmt"
	"time"
)

//Creates a new IPFIX message with specified observation domain id
//The other fields can not be filled in yet.
func NewIPFIXMessage(odid int) (*IPFIXMessage, error) {
	if odid < 0 {
		return nil, fmt.Errorf("Can not have a negative observation domain: %d", odid)
	}
	return &IPFIXMessage{
		MessageHeader: IPFIXMessageHeader{
			VersionNumber:       IPFIXVersion,
			Length:              0,                       //Needs to be set when finalized
			ExportTime:          time.Unix(233431200, 0), //A long time ago. Needs to be set when message is sent (finalized)
			SequenceNumber:      0,                       //Needs to be set by sending entity (finalized)
			ObservationDomainID: uint32(odid),
		},
		Sets: []*IPFIXSet{},
	}, nil
}

//Finalizes the IPFIX message and calculates the length
func (ipfixmsg *IPFIXMessage) Finalize(timestamp time.Time, sequencenumber int) error {
	ipfixmsg.MessageHeader.ExportTime = timestamp
	ipfixmsg.MessageHeader.SequenceNumber = uint32(sequencenumber)

	//Calculate the length of the message
	ipfixmsg.MessageHeader.Length = ipfixMessageHeaderLength
	for _, rec := range ipfixmsg.Sets {
		ipfixmsg.MessageHeader.Length += rec.SetHeader.Length
	}

	return nil
}
