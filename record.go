package ipfixmessage

/*
   IPFIX defines three record formats:
   the Template Record format, the Options Template Record format, and  the Data Record format.

   This is just the Record Interface definition
*/

// Record defines the interface for IPFIX Set Records
type Record interface {
	MarshalBinary() (data []byte, err error) // Each record *must* implement the encoding/BinaryMarshaler interface
	UnmarshalBinary(data []byte) error       // Each record *must* implement the encoding/BinaryUnmarshaler interface
	Len() uint16                             // The size in Octets of this record, when Marshalled
	String() string                          // Return a string representation
}