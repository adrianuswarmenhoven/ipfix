package ipfixmessage

import (
  "fmt"
  // "io"
  "time"
)

type MessageHeader struct {
  VersionNumber       uint16    //Must be Version
  Length              uint16    //Total length of the IPFIX Message, measured in octets, including Message Header and Set(s).
  ExportTime          time.Time //Time, in seconds (int32), since 0000 UTC Jan 1, 1970, at which the IPFIX Message Header leaves the Exporter.
  SequenceNumber      uint32    //Incremental sequence counter modulo 2^32 of all IPFIX Data Records sent on this PR-SCTP stream from the current Observation Domain by the Exporting Process.
  ObservationDomainID uint32    //A 32-bit identifier of the Observation Domain that is locally unique to the Exporting Process.
}

// Len returns the length of the Message Header in bytes.
func (h *MessageHeader) Len() int {
  return ipfixMessageHeaderLength
}

func (h *MessageHeader) String() string {
  return fmt.Sprintf("version=%d, length=%d, time=%s, seq=%d, odid=%d",
    h.VersionNumber, h.Length, h.ExportTime.Unix(), h.SequenceNumber, h.ObservationDomainID)
}

/*
// Unmarshal a message header from a reader.
func (h *MessageHeader) Unmarshal(r io.Reader) error {
  if err := read.Uint16(&h.VersionNumber, r); err != nil {
    return err
  }
  if err := read.Uint16(&h.Length, r); err != nil {
    return err
  }
  if err := read.Uint32(&h.ExportTime, r); err != nil {
    return err
  }
  if err := read.Uint32(&h.SequenceNumber, r); err != nil {
    return err
  }
  if err := read.Uint32(&h.ObservationDomainID, r); err != nil {
    return err
  }

  return nil
}
*/
