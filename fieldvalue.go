package ipfixmessage

import (
	"fmt"
	"net"
	"time"
)

// FieldValue Interface definition
type FieldValue interface {
	MarshalBinary() (data []byte, err error) // Each FieldValue *must* implement the encoding/BinaryMarshaler interface
	UnmarshalBinary(data []byte) error       // Each FieldValue *must* implement the encoding/BinaryUnmarshaler interface
	Len() uint16                             // The size in Octets of this record, when Marshalled
}

/* */
// FieldValueUnsigned8 , "unsigned8" represents a non-negative integer value in the range of 0 to 255.
type FieldValueUnsigned8 struct {
	Value uint8
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueUnsigned8) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueUnsigned8) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueUnsigned8) Len() uint16 {
	return 1
}

/* */
// FieldValueUnsigned16 , "unsigned16" represents a non-negative integer value in the range of 0 to 65535.
type FieldValueUnsigned16 struct {
	Value uint16
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueUnsigned16) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueUnsigned16) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueUnsigned16) Len() uint16 {
	return 2
}

/* */
// FieldValueUnsigned32 , "unsigned32" represents a non-negative integer value in the range of 0 to 4294967295
type FieldValueUnsigned32 struct {
	Value uint32
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueUnsigned32) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueUnsigned32) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueUnsigned32) Len() uint16 {
	return 4
}

/* */
// FieldValueUnsigned64 , "unsigned64" represents a non-negative integer value in the range of 0 to 18446744073709551615
type FieldValueUnsigned64 struct {
	Value uint64
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueUnsigned64) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueUnsigned64) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueUnsigned64) Len() uint16 {
	return 8
}

/* */
// FieldValueSigned8 , "signed8" represents an integer value in the range of -128 to 127
type FieldValueSigned8 struct {
	Value int8
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueSigned8) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueSigned8) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueSigned8) Len() uint16 {
	return 1
}

/* */
// FieldValueSigned16 , "signed16" represents an integer value in the range of -32768 to 32767.
type FieldValueSigned16 struct {
	Value int16
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueSigned16) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueSigned16) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueSigned16) Len() uint16 {
	return 2
}

/* */
// FieldValueSigned32 , "signed32" represents an integer value in the range of -2147483648 to 2147483647.
type FieldValueSigned32 struct {
	Value int32
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueSigned32) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueSigned32) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueSigned32) Len() uint16 {
	return 4
}

/* */
// FieldValueSigned64 , "signed64" represents an integer value in the range of -9223372036854775808 to 9223372036854775807.
type FieldValueSigned64 struct {
	Value int64
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueSigned64) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueSigned64) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueSigned64) Len() uint16 {
	return 8
}

/* */
// FieldValueFloat32 , "float32" corresponds to an IEEE single-precision 32-bit floating-point type as defined in [IEEE.754.2008].
type FieldValueFloat32 struct {
	Value float32
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueFloat32) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueFloat32) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueFloat32) Len() uint16 {
	return 4
}

/* */
// FieldValueFloat64 , "float64" corresponds to an IEEE double-precision 64-bit floating-point type as defined in [IEEE.754.2008].
type FieldValueFloat64 struct {
	Value float64
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueFloat64) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueFloat64) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueFloat64) Len() uint16 {
	return 8
}

/* */
// FieldValueBoolean ,  "boolean" represents a binary value.  The only allowed values are "true" and "false".
//The boolean data type is specified according to the TruthValue in [RFC2579].
//It is encoded as a single-octet integer per Section 6.1.1, with the value 1 for true and value 2 for false.
//Every other value is undefined.
type FieldValueBoolean struct {
	Value bool
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueBoolean) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueBoolean) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueBoolean) Len() uint16 {
	return 1
}

/* */
// FieldValueMacAddress , "macAddress" represents a MAC-48 address as defined in [IEEE.802-3.2012].
type FieldValueMacAddress struct {
	Value net.HardwareAddr
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueMacAddress) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueMacAddress) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueMacAddress) Len() uint16 {
	return 6 //48 bits or 12 characters in hex
}

/* */
// FieldValueOctetArray , "octetArray" represents a finite-length string of octets.
type FieldValueOctetArray struct {
	Value []byte
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueOctetArray) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueOctetArray) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueOctetArray) Len() uint16 {
	return uint16(len(fv.Value))
}

/* */
// FieldValueString , "string" represents a finite-length string of valid characters from the Unicode coded character set [ISO.10646].
//Unicode incorporates ASCII [RFC20] and the characters of many other international character sets.
type FieldValueString struct {
	Value string
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueString) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueString) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueString) Len() uint16 {
	return uint16(len([]byte(fv.Value))) //Must convert to []byte first...
}

/* */
// FieldValueDateTimeSeconds , "dateTimeSeconds" represents a time value expressed with second-level precision.
//The dateTimeSeconds data type is an unsigned 32-bit integer in network byte order containing the number of seconds since the UNIX epoch, 1 January 1970 at 00:00 UTC, as defined in [POSIX.1].
//dateTimeSeconds is encoded identically to the IPFIX Message Header Export Time field.
type FieldValueDateTimeSeconds struct {
	Value time.Time
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueDateTimeSeconds) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueDateTimeSeconds) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueDateTimeSeconds) Len() uint16 {
	return 4
}

/* */
// FieldValueDateTimeMilliseconds , "dateTimeMilliseconds" represents a time value expressed with millisecond-level precision.
// The dateTimeMilliseconds data type is an unsigned 64-bit integer in network byte order containing the number of milliseconds since the UNIX epoch, 1 January 1970 at 00:00 UTC, as defined in [POSIX.1].
type FieldValueDateTimeMilliseconds struct {
	Value time.Time
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueDateTimeMilliseconds) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueDateTimeMilliseconds) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueDateTimeMilliseconds) Len() uint16 {
	return 8
}

/* */
// FieldValueDateTimeMicroseconds , "dateTimeMicroseconds" represents a time value expressed with microsecond-level precision.
//The dateTimeMicroseconds data type is a 64-bit field encoded according to the NTP Timestamp format as defined in Section 6 of [RFC5905].
//This field is made up of two unsigned 32-bit integers in network byte order: Seconds and Fraction.
//The Seconds field is the number of seconds since the NTP epoch, 1 January 1900 at 00:00 UTC.
//The Fraction field is the fractional number of seconds in units of 1/(2^32) seconds (approximately 233 picoseconds).
//It can represent dates between 1 January 1900 and 8 February 2036 in the current NTP era.
type FieldValueDateTimeMicroseconds struct {
	Value time.Time
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueDateTimeMicroseconds) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueDateTimeMicroseconds) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueDateTimeMicroseconds) Len() uint16 {
	return 8
}

/* */
// FieldValueDateTimeNanoseconds , "dateTimeNanoseconds" represents a time value expressed with nanosecond-level precision.
//The dateTimeNanoseconds data type is a 64-bit field encoded according to the NTP Timestamp format as defined in Section 6 of [RFC5905].
//This field is made up of two unsigned 32-bit integers in network byte order: Seconds and Fraction.
//The Seconds field is the number of seconds since the NTP epoch, 1 January 1900 at 00:00 UTC.
//The Fraction field is the fractional number of seconds in units of 1/(2^32) seconds (approximately 233 picoseconds).
//It can represent dates between 1 January 1900 and 8 February 2036 in the current NTP era.
//
//Note that dateTimeMicroseconds and dateTimeNanoseconds share an identical encoding.  There is no restriction on the interpretation of the Fraction field for the dateTimeNanoseconds data type.
type FieldValueDateTimeNanoseconds struct {
	Value time.Time
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueDateTimeNanoseconds) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueDateTimeNanoseconds) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueDateTimeNanoseconds) Len() uint16 {
	return 8
}

/* */
// FieldValueIPv4Address , "ipv4Address" represents an IPv4 address.
type FieldValueIPv4Address struct {
	Value net.IP
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueIPv4Address) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueIPv4Address) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueIPv4Address) Len() uint16 {
	return 4
}

/* */
// FieldValueIPv6Address , "ipv6Address" represents an IPv6 address.
type FieldValueIPv6Address struct {
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueIPv6Address) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueIPv6Address) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueIPv6Address) Len() uint16 {
	return 16
}

/* */
// FieldValueBasicList , "basicList" supports structured data export as described in [RFC6313];
//see Section 4.5.1 of that document for encoding details.
type FieldValueBasicList struct {
	Value *BasicList
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv FieldValueBasicList) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv FieldValueBasicList) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv FieldValueBasicList) Len() uint16 {
	return fv.Value.Len()
}

/* */
// FieldValueSubTemplateList , "subTemplateList" supports structured data export as described in [RFC6313];
//see Section 4.5.2 of that document for encoding details.
type FieldValueSubTemplateList struct {
}

// FieldValueSubTemplateMultiList , "subTemplateMultiList" supports structured data export as described in [RFC6313];
//see Section 4.5.3 of that document for encoding details.
type FieldValueSubTemplateMultiList struct {
}
