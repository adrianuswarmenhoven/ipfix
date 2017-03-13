package ipfixmessage

import (
	"bytes"
	"encoding/binary"
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

	//AssociatedTemplates Templates points to the list of active templates (whether in a session or not). Without a template record a data record can not be encoded or decoded
	AssociatedTemplates *ActiveTemplates
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
// Observation Domain ID: A 32-bit identifier of the Observation Domain that is locally unique to the Exporting Process.
// The Exporting Process uses the Observation Domain ID to uniquely identify to the Collecting Process the Observation Domain that metered the Flows.
// It is RECOMMENDED that this identifier also be unique per IPFIX Device.
// Collecting Processes SHOULD use the Transport Session and the Observation Domain ID field to separate different export streams originating from the same Exporter.
// The Observation Domain ID SHOULD be 0 when no specific Observation Domain ID is relevant for the entire IPFIX Message, for example, when exporting the Exporting Process Statistics, or in the case of a hierarchy of Collectors when aggregated Data Records are exported.
func (ipfixmsg *Message) SetObservationDomainID(odid uint32) error {
	//Since the only rule is that the odid is 32 bits we always return nil for error.
	//Although the text mentions 32-bit identifier, we enforce it as unsigned 32-bit integer
	ipfixmsg.ObservationDomainID = odid
	return nil
}

// SetExportTime sets the export time on the message.
// Export Time: Time at which the IPFIX Message Header leaves the Exporter, expressed in seconds since the UNIX epoch of 1 January 1970 at 00:00 UTC, encoded as an unsigned 32-bit integer.
func (ipfixmsg *Message) SetExportTime(exporttime time.Time) error {
	//Since an unsigned 32-bit integer can not be negative, we know for sure that we can not export any timestamps before UNIX epoch of 1 January 1970
	if exporttime.Before(time.Unix(0, 0)) {
		return NewError(fmt.Sprintf("Can not set export time to before UNIX epoch of 1 January 1970, but got %+v", exporttime), ErrCritical)
	}

	ipfixmsg.ExportTime = exporttime
	return nil
}

// SetSequenceNumber sets the sequence number on the message
// Sequence Number: Incremental sequence counter modulo 2^32 of all IPFIX Data Records sent in the current stream from the current Observation Domain by the Exporting Process.
// Each SCTP Stream counts sequence numbers separately, while all messages in a TCP connection or UDP session are considered to be part of the same stream.
// This value can be used by the Collecting Process to identify whether any IPFIX Data Records have been missed.
// Template and Options Template Records do not increase the Sequence Number..
func (ipfixmsg *Message) SetSequenceNumber(sequencenumber uint32) error {
	//Since it is an incremental counter we will never go negative hence we assume uint32
	//NB: Sequence Numbers do *not* have to be unique as a message may contain only Template and/or Template Options Records
	ipfixmsg.SequenceNumber = sequencenumber
	return nil
}

// AddSet adds an existing set to the message.
func (ipfixmsg *Message) AddSet(newset *Set) (err error) {
	if int(ipfixmsg.Len())+int(newset.Len()) > 65535 {
		return NewError(fmt.Sprintf("Can not add set to message; resulting length would be %d but must be <= 65535", int(ipfixmsg.Len())+int(newset.Len())), ErrCritical)
	}
	ipfixmsg.Sets = append(ipfixmsg.Sets, newset)
	return nil
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
func (ipfixmsg *Message) MarshalBinary() (data []byte, err error) {
	//Set ID value identifies the Set.  A value of 2 is reserved for the Template Set.  A value of 3 is reserved for the Option Template Set.
	//All other values from 4 to 255 are reserved for future use. Values above 255 are used for Data Sets.
	if ipfixmsg.VersionNumber < IPFIXVersion {
		return nil, NewError(fmt.Sprintf("Invalid IPFIX Version Number: %d", ipfixmsg.VersionNumber), ErrCritical)
	}
	buf := new(bytes.Buffer) //should get from pool?

	err = binary.Write(buf, binary.BigEndian, uint16(IPFIXVersion))
	if err != nil {
		return nil, err
	}

	totalsetlength := uint16(16)
	for _, set := range ipfixmsg.Sets {
		if int(totalsetlength)+int(set.Len()) > 65535 {
			return nil, NewError(fmt.Sprintf("Invalid total length of message. Got %d but should be <= 65535", int(totalsetlength)+int(set.Len())), ErrCritical)
		}
		totalsetlength += set.Len()
	}
	err = binary.Write(buf, binary.BigEndian, totalsetlength)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, uint32(ipfixmsg.ExportTime.Unix()))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, uint32(ipfixmsg.SequenceNumber))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, uint32(ipfixmsg.ObservationDomainID))
	if err != nil {
		return nil, err
	}

	data = buf.Bytes()
	for _, set := range ipfixmsg.Sets {
		setdat, suberr := (*set).MarshalBinary()
		if suberr != nil {
			if err == nil {
				err = NewError("Sub errors marshalling message.", ErrFailure)
			}
			err.(*ProtocolError).Stack(*suberr.(*ProtocolError))
		}
		data = append(data, setdat...)
	}

	return data, err
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
func (ipfixmsg *Message) UnmarshalBinary(data []byte) (err error) {
	if data == nil || len(data) < 16 {
		return NewError(fmt.Sprintf("Can not unmarshal, invalid data. %#v", data), ErrCritical)
	}
	if ipfixmsg.AssociatedTemplates == nil {
		return NewError(fmt.Sprintf("Can not have nil pointer to associated templates"), ErrCritical)
	}

	ipfixmsg.VersionNumber = binary.BigEndian.Uint16(data[0:2])
	if ipfixmsg.VersionNumber < IPFIXVersion {
		return NewError(fmt.Sprintf("Unusable IPFIX version. Want at least 10, but got %d", ipfixmsg.VersionNumber), ErrCritical)
	}

	totalmessagelength := binary.BigEndian.Uint16(data[2:4])
	if int(totalmessagelength) > len(data) {
		return NewError(fmt.Sprintf("Can not unmarshal, invalid length. Message states %d but only have %d bytes of data", totalmessagelength, len(data)), ErrCritical)
	}
	err = ipfixmsg.SetExportTime(time.Unix(int64(binary.BigEndian.Uint32(data[4:8])), 0))
	if err != nil {
		return err
	}

	err = ipfixmsg.SetSequenceNumber(binary.BigEndian.Uint32(data[8:12]))
	if err != nil {
		return err
	}

	err = ipfixmsg.SetObservationDomainID(binary.BigEndian.Uint32(data[12:16]))
	if err != nil {
		return err
	}

	cursor := uint16(16)
	for cursor < (totalmessagelength - 4) { //Must have at least 4 bytes (actually more) for a set to be unmarshalled
		tmpset := NewBlankSet()
		tmpset.AssociateTemplates(ipfixmsg.AssociatedTemplates)
		suberr := tmpset.UnmarshalBinary(data[cursor:])
		if suberr != nil {
			if err == nil {
				err = NewError("Sub errors unmarshalling message.", ErrFailure)
			}
			err.(*ProtocolError).Stack(suberr)
		}

		if tmpset.SetID == 2 { //Need to add/update all the messages
			for _, rec := range tmpset.Records {
				switch (*rec).(type) {
				case *TemplateRecord:
					suberr := ipfixmsg.AssociatedTemplates.Set((*rec).(*TemplateRecord).TemplateID, (*rec).(*TemplateRecord))
					if suberr != nil {
						if err == nil {
							err = NewError("Sub errors unmarshalling message.", ErrFailure)
						}
						err.(*ProtocolError).Stack(suberr)
					}
				case *DataRecord:
					return NewError(fmt.Sprintf("Datarecord in template set"), ErrCritical)
				}
			}
		}
		ipfixmsg.Sets = append(ipfixmsg.Sets, tmpset)
		cursor += tmpset.Len()
	}
	return err
}
