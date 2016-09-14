package ipfixmessage

import (
	"fmt"
	"time"
)

// Message represents an IPFIX Message.
// An IPFIX Message consists of a Message Header, followed by zero or
// more Sets.  The Sets can be any of these three possible types:
// Data Set, Template Set, or Options Template Set.
type Message struct {
	MessageHeader *MessageHeader

	IsFinalized bool // Needed so we know whether we actually *can* marshal to bytes

	Sets []*Set
}

// NewMessage creates a new IPFIX message.
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

// ODID sets the observation domain id on the message.
func (ipfixmsg *Message) ODID(odid uint32) *Message {
	ipfixmsg.MessageHeader.ObservationDomainID = odid

	return ipfixmsg
}

// ExportTime sets the export time on the message.
func (ipfixmsg *Message) ExportTime(exporttime time.Time) *Message {
	ipfixmsg.MessageHeader.ExportTime = exporttime

	return ipfixmsg
}

// SequenceNumber sets the sequence number on the message.
func (ipfixmsg *Message) SequenceNumber(sequencenumber uint32) *Message {
	ipfixmsg.MessageHeader.SequenceNumber = sequencenumber

	return ipfixmsg
}

// Set adds an existing set to the message.
func (ipfixmsg *Message) Set(newset *Set) *Message {
	ipfixmsg.Sets = append(ipfixmsg.Sets, newset)

	return ipfixmsg
}

// Finalize finalizes the IPFIX message and calculates the length.
// It does so by getting all the lengths of messages and sets.
func (ipfixmsg *Message) Finalize() (*Message, error) {
	//Calculate the length of the message
	ipfixmsg.MessageHeader.Length = ipfixMessageHeaderLength
	for _, rec := range ipfixmsg.Sets {
		ipfixmsg.MessageHeader.Length += rec.SetHeader.Length
	}
	ipfixmsg.IsFinalized = true

	return ipfixmsg, nil
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
func (ipfixmsg *Message) MarshalBinary() (data []byte, err error) {
	if !ipfixmsg.IsFinalized {
		return nil, fmt.Errorf("Can not marshal message; not yet finalized")
	}

	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
func (ipfixmsg *Message) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("Can not unmarshal, invalid data. %#v", data)
	}

	return fmt.Errorf("Not yet implemented!")
}
