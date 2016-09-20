package ipfixmessage

import (
	"fmt"
	"time"
)

/* N.B. per SequenceNumber:

Incremental sequence counter modulo 2^32 of all IPFIX Data Records sent in the current stream from the current Observation Domain by the Exporting Process.
Each SCTP Stream counts sequence numbers separately, while all messages in a TCP connection or UDP session are considered to be part of the same stream.
This value can be used by the Collecting Process to identify whether any IPFIX Data Records have been missed.
Template and Options Template Records do not increase the Sequence Number.

The sequence number comes from an external source and is context based.

*/

// Message represents an IPFIX Message.
// An IPFIX Message consists of a Message Header, followed by zero or
// more Sets.  The Sets can be any of these three possible types:
// Data Set, Template Set, or Options Template Set.
type Message struct {
	VersionNumber       uint16    // Must be Version IPFIXVersion
	ExportTime          time.Time // Time, in seconds (int32), since the UNIX epoch, 1 January 1970 at 00:00 UTC, modulo 2^32, at which the IPFIX Message Header leaves the Exporter.
	SequenceNumber      uint32    // Incremental sequence counter modulo 2^32 of all IPFIX Data Records sent on this PR-SCTP stream from the current Observation Domain by the Exporting Process.
	ObservationDomainID uint32    // A 32-bit identifier of the Observation Domain that is locally unique to the Exporting Process.
	Sets                []*Set
}

// NewMessage creates a new IPFIX message.
func NewMessage() (*Message, error) {
	return &Message{
		VersionNumber:       IPFIXVersion,
		ExportTime:          time.Unix(233431200, 0), //A long time ago. Needs to be set when message is sent
		SequenceNumber:      0,                       //Needs to be set by sending entity
		ObservationDomainID: 0,                       //Default ODID is 0
		Sets:                make([]*Set, 0, 0),      //Create an empty slice of sets
	}, nil
}

// String returns the string representation of the Message
func (ipfixmsg *Message) String() string {
	retstr := fmt.Sprintf("version=%d, length=%d, export time=%s, sequence=%d, odid=%d", ipfixmsg.VersionNumber, ipfixmsg.Len(), ipfixmsg.ExportTime, ipfixmsg.SequenceNumber, ipfixmsg.ObservationDomainID)
	if len(ipfixmsg.Sets) > 0 {
		retstr += "\n"
		for _, st := range ipfixmsg.Sets {
			retstr += st.String() + "\n"
		}
	}
	return retstr
}

// Len returns the size in octets of the Set
func (ipfixmsg *Message) Len() uint16 {
	msglen := uint16(ipfixMessageHeaderLength) //We start out with 16 bytes for the length of the header
	for _, st := range ipfixmsg.Sets {
		msglen += (*st).Len()
	}
	return msglen
}

// SetODID sets the observation domain id on the message.
func (ipfixmsg *Message) SetODID(odid uint32) *Message {
	ipfixmsg.ObservationDomainID = odid

	return ipfixmsg
}

// SetExportTime sets the export time on the message.
func (ipfixmsg *Message) SetExportTime(exporttime time.Time) *Message {
	ipfixmsg.ExportTime = exporttime

	return ipfixmsg
}

// SetSequenceNumber sets the sequence number on the message.
func (ipfixmsg *Message) SetSequenceNumber(sequencenumber uint32) *Message {
	ipfixmsg.SequenceNumber = sequencenumber

	return ipfixmsg
}

// AddSet adds an existing set to the message.
func (ipfixmsg *Message) AddSet(newset *Set) *Message {
	ipfixmsg.Sets = append(ipfixmsg.Sets, newset)

	return ipfixmsg
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
func (ipfixmsg *Message) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
func (ipfixmsg *Message) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("Can not unmarshal, invalid data. %#v", data)
	}

	return fmt.Errorf("Not yet implemented!")
}
