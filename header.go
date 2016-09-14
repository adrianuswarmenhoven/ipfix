package ipfixmessage

import (
	"fmt"
	"time"
)

// MessageHeader contains the mandatory fields for every IPFIX Message.
// Each Message Header field is exported in network byte order.
type MessageHeader struct {
	VersionNumber       uint16    // Must be Version IPFIXVersion
	Length              uint16    // Total length of the IPFIX Message, measured in octets, including Message Header and Set(s).
	ExportTime          time.Time // Time, in seconds (int32), since the UNIX epoch, 1 January 1970 at 00:00 UTC, modulo 2^32, at which the IPFIX Message Header leaves the Exporter.
	SequenceNumber      uint32    // Incremental sequence counter modulo 2^32 of all IPFIX Data Records sent on this PR-SCTP stream from the current Observation Domain by the Exporting Process.
	ObservationDomainID uint32    // A 32-bit identifier of the Observation Domain that is locally unique to the Exporting Process.
}

/* N.B. per SequenceNumber:

Incremental sequence counter modulo 2^32 of all IPFIX Data Records sent in the current stream from the current Observation Domain by the Exporting Process.
Each SCTP Stream counts sequence numbers separately, while all messages in a TCP connection or UDP session are considered to be part of the same stream.
This value can be used by the Collecting Process to identify whether any IPFIX Data Records have been missed.
Template and Options Template Records do not increase the Sequence Number.

*/

// Len returns the length of the message header in bytes.
// We return a constant because the package type has time.Time instead of uint32
func (h *MessageHeader) Len() uint16 {
	return ipfixMessageHeaderLength
}

// String returns the string representation of the message header
func (h *MessageHeader) String() string {
	return fmt.Sprintf("version=%d, length=%d, export time=%s, sequence=%d, odid=%d", h.VersionNumber, h.Length, h.ExportTime, h.SequenceNumber, h.ObservationDomainID)
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (h *MessageHeader) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (h *MessageHeader) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("Can not unmarshal, invalid data. %#v", data)
	}

	return fmt.Errorf("Not yet implemented!")
}
