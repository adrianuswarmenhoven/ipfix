package ipfixmessage

import (
	"time"
)

type Message struct {
	MessageHeader *MessageHeader

	IsFinalized bool

	Sets []*Set
}

//Creates a new IPFIX message with specified observation domain id
//The other fields can not be filled in yet.
func NewMessage() (*Message, error) {
	return &Message{
		MessageHeader: &MessageHeader{
			VersionNumber:       IPFIXVersion,
			Length:              0,                       //Needs to be set when finalized
			ExportTime:          time.Unix(233431200, 0), //A long time ago. Needs to be set when message is sent
			SequenceNumber:      0,                       //Needs to be set by sending entity
			ObservationDomainID: 0,                       //Default ODID is 0
		},
		Sets:        make([]*Set, 0, 0), //Create an empty slice of sets
		IsFinalized: false,              //Safeguard before sending
	}, nil
}

//Sets the observation domain id on the message
func (ipfixmsg *Message) ODID(odid uint32) *Message {
	ipfixmsg.MessageHeader.ObservationDomainID = odid

	return ipfixmsg
}

//Sets the export time on the message
func (ipfixmsg *Message) ExportTime(exporttime time.Time) *Message {
	ipfixmsg.MessageHeader.ExportTime = exporttime

	return ipfixmsg
}

//Sets the sequence number on the message
func (ipfixmsg *Message) SequenceNumber(sequencenumber uint32) *Message {
	ipfixmsg.MessageHeader.SequenceNumber = sequencenumber

	return ipfixmsg
}

//Adds a set to the message
func (ipfixmsg *Message) Set(newset *Set) *Message {
	ipfixmsg.Sets = append(ipfixmsg.Sets, newset)

	return ipfixmsg
}

//Finalizes the IPFIX message and calculates the length
func (ipfixmsg *Message) Finalize() (*Message, error) {
	//Calculate the length of the message
	ipfixmsg.MessageHeader.Length = ipfixMessageHeaderLength
	for _, rec := range ipfixmsg.Sets {
		ipfixmsg.MessageHeader.Length += rec.SetHeader.Length
	}
	ipfixmsg.IsFinalized = true

	return ipfixmsg, nil
}
