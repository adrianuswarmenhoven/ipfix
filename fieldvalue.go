package ipfixmessage

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"reflect"
	"time"
)

// FieldValue Interface definition
// Note that all implementations must be pointer receivers so that unmarshal can change the value
type FieldValue interface {
	MarshalBinary() ([]byte, error)    // Each FieldValue *must* implement the encoding/BinaryMarshaler interface
	UnmarshalBinary(data []byte) error // Each FieldValue *must* implement the encoding/BinaryUnmarshaler interface
	Len() uint16                       // The size in Octets of this record, when Marshalled
	Value() interface{}                // Returns the value of this fieldvalue
	Set(val interface{}) error         // Sets the value of this FieldValue
}

func marshalBinarySingleValue(val interface{}) ([]byte, error) {
	buf := new(bytes.Buffer) //should get from pool?
	err := binary.Write(buf, binary.BigEndian, val)
	if err == nil {
		return buf.Bytes(), nil
	}
	return nil, err
}

func unmarshalBinaryOctets(data []byte, val interface{}) error {
	buf := bytes.NewReader(data)
	return binary.Read(buf, binary.BigEndian, val)
}

/* */
// FieldValueUnsigned8 , "unsigned8" represents a non-negative integer value in the range of 0 to 255.
type FieldValueUnsigned8 struct {
	value uint8
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueUnsigned8) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueUnsigned8) UnmarshalBinary(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	fv.value = data[0]
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueUnsigned8) Len() uint16 {
	return 1
}

// Value returns FieldValue's value
func (fv *FieldValueUnsigned8) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueUnsigned8) Set(val interface{}) error {
	switch val.(type) {
	case uint8:
		fv.value = val.(uint8)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueUnsigned16 , "unsigned16" represents a non-negative integer value in the range of 0 to 65535.
type FieldValueUnsigned16 struct {
	value uint16
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueUnsigned16) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueUnsigned16) UnmarshalBinary(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	fv.value = binary.BigEndian.Uint16(data)
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueUnsigned16) Len() uint16 {
	return 2
}

// Value returns FieldValue's value
func (fv *FieldValueUnsigned16) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueUnsigned16) Set(val interface{}) error {
	switch val.(type) {
	case uint16:
		fv.value = val.(uint16)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueUnsigned32 , "unsigned32" represents a non-negative integer value in the range of 0 to 4294967295
type FieldValueUnsigned32 struct {
	value uint32
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueUnsigned32) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueUnsigned32) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	fv.value = binary.BigEndian.Uint32(data)
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueUnsigned32) Len() uint16 {
	return 4
}

// Value returns FieldValue's value
func (fv *FieldValueUnsigned32) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueUnsigned32) Set(val interface{}) error {
	switch val.(type) {
	case uint32:
		fv.value = val.(uint32)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueUnsigned64 , "unsigned64" represents a non-negative integer value in the range of 0 to 18446744073709551615
type FieldValueUnsigned64 struct {
	value uint64
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueUnsigned64) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueUnsigned64) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	fv.value = binary.BigEndian.Uint64(data)
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueUnsigned64) Len() uint16 {
	return 8
}

// Value returns FieldValue's value
func (fv *FieldValueUnsigned64) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueUnsigned64) Set(val interface{}) error {
	switch val.(type) {
	case uint64:
		fv.value = val.(uint64)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueSigned8 , "signed8" represents an integer value in the range of -128 to 127
type FieldValueSigned8 struct {
	value int8
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueSigned8) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueSigned8) UnmarshalBinary(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	fv.value = int8(data[0])
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueSigned8) Len() uint16 {
	return 1
}

// Value returns FieldValue's value
func (fv *FieldValueSigned8) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueSigned8) Set(val interface{}) error {
	switch val.(type) {
	case int8:
		fv.value = val.(int8)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueSigned16 , "signed16" represents an integer value in the range of -32768 to 32767.
type FieldValueSigned16 struct {
	value int16
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueSigned16) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueSigned16) UnmarshalBinary(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	fv.value = int16(binary.BigEndian.Uint16(data))
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueSigned16) Len() uint16 {
	return 2
}

// Value returns FieldValue's value
func (fv *FieldValueSigned16) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueSigned16) Set(val interface{}) error {
	switch val.(type) {
	case int16:
		fv.value = val.(int16)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueSigned32 , "signed32" represents an integer value in the range of -2147483648 to 2147483647.
type FieldValueSigned32 struct {
	value int32
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueSigned32) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueSigned32) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	fv.value = int32(binary.BigEndian.Uint32(data))
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueSigned32) Len() uint16 {
	return 4
}

// Value returns FieldValue's value
func (fv *FieldValueSigned32) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueSigned32) Set(val interface{}) error {
	switch val.(type) {
	case int32:
		fv.value = val.(int32)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueSigned64 , "signed64" represents an integer value in the range of -9223372036854775808 to 9223372036854775807.
type FieldValueSigned64 struct {
	value int64
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueSigned64) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueSigned64) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	fv.value = int64(binary.BigEndian.Uint64(data))
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueSigned64) Len() uint16 {
	return 8
}

// Value returns FieldValue's value
func (fv *FieldValueSigned64) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueSigned64) Set(val interface{}) error {
	switch val.(type) {
	case int64:
		fv.value = val.(int64)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueFloat32 , "float32" corresponds to an IEEE single-precision 32-bit floating-point type as defined in [IEEE.754.2008].
type FieldValueFloat32 struct {
	value float32
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueFloat32) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueFloat32) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	fv.value = math.Float32frombits(binary.BigEndian.Uint32(data))
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueFloat32) Len() uint16 {
	return 4
}

// Value returns FieldValue's value
func (fv *FieldValueFloat32) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueFloat32) Set(val interface{}) error {
	switch val.(type) {
	case float32:
		fv.value = val.(float32)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueFloat64 , "float64" corresponds to an IEEE double-precision 64-bit floating-point type as defined in [IEEE.754.2008].
type FieldValueFloat64 struct {
	value float64
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueFloat64) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueFloat64) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	fv.value = math.Float64frombits(binary.BigEndian.Uint64(data))
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueFloat64) Len() uint16 {
	return 8
}

// Value returns FieldValue's value
func (fv *FieldValueFloat64) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueFloat64) Set(val interface{}) error {
	switch val.(type) {
	case float64:
		fv.value = val.(float64)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueBoolean ,  "boolean" represents a binary value.  The only allowed values are "true" and "false".
//The boolean data type is specified according to the TruthValue in [RFC2579].
//It is encoded as a single-octet integer per Section 6.1.1, with the value 1 for true and value 2 for false.
//Every other value is undefined.
type FieldValueBoolean struct {
	value bool
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueBoolean) MarshalBinary() ([]byte, error) {
	if fv.value {
		return marshalBinarySingleValue(uint8(1))
	}
	return marshalBinarySingleValue(uint8(2))
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueBoolean) UnmarshalBinary(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	switch data[0] {
	case 1:
		fv.value = true
	case 2:
		fv.value = false
	default:
		return fmt.Errorf("Invalid encoded value for boolean: %d, must be either 1 or 2", data[0])
	}
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueBoolean) Len() uint16 {
	return 1
}

// Value returns FieldValue's value
func (fv *FieldValueBoolean) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueBoolean) Set(val interface{}) error {
	switch val.(type) {
	case bool:
		fv.value = val.(bool)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueMacAddress , "macAddress" represents a MAC-48 address as defined in [IEEE.802-3.2012].
type FieldValueMacAddress struct {
	value net.HardwareAddr
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
// Address types -- macAddress, ipv4Address, and ipv6Address -- MUST be encoded the same way as the integral data types, as six, four, and sixteen octets in network byte order, respectively.
func (fv *FieldValueMacAddress) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueMacAddress) UnmarshalBinary(data []byte) error {
	if len(data) < 6 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	return unmarshalBinaryOctets(data, &fv.value)
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueMacAddress) Len() uint16 {
	return 6 //48 bits or 12 characters in hex
}

// Value returns FieldValue's value
func (fv *FieldValueMacAddress) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueMacAddress) Set(val interface{}) error {
	switch val.(type) {
	case net.HardwareAddr:
		fv.value = val.(net.HardwareAddr)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueOctetArray , "octetArray" represents a finite-length string of octets.
// The octetArray data type has no encoding rules; it represents a raw array of zero or more octets, with the interpretation of the octets defined in the Information Element definition.
type FieldValueOctetArray struct {
	value []byte
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueOctetArray) MarshalBinary() ([]byte, error) {
	content, err := marshalBinarySingleValue(fv.value)
	if err != nil {
		return []byte{}, err
	}
	return content, nil
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueOctetArray) UnmarshalBinary(data []byte) error {
	fv.value = make([]byte, len(data))
	return unmarshalBinaryOctets(data, &fv.value)
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueOctetArray) Len() uint16 {
	return uint16(len(fv.value))
}

// Value returns FieldValue's value
func (fv *FieldValueOctetArray) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueOctetArray) Set(val interface{}) error {
	switch val.(type) {
	case string:
		fv.value = []byte(val.(string))
	case []byte:
		fv.value = val.([]byte)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueString , "string" represents a finite-length string of valid characters from the Unicode coded character set [ISO.10646].
// Unicode incorporates ASCII [RFC20] and the characters of many other international character sets.
// The string data type MUST be encoded in UTF-8 [RFC3629] format.  The string is sent as an array of zero or more octets using an Information Element of fixed or variable length.
type FieldValueString struct {
	value string
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueString) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue([]byte(fv.value))
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueString) UnmarshalBinary(data []byte) error {
	tmpval := make([]byte, len(data))
	err := unmarshalBinaryOctets(data, tmpval)
	if err != nil {
		return err
	}
	fv.value = string(tmpval)
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueString) Len() uint16 {
	return uint16(len([]byte(fv.value))) //Must convert to []byte first...
}

// Value returns FieldValue's value
func (fv *FieldValueString) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueString) Set(val interface{}) error {
	switch val.(type) {
	case string:
		fv.value = val.(string)
	case []byte:
		fv.value = string(val.([]byte))
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueDateTimeSeconds , "dateTimeSeconds" represents a time value expressed with second-level precision.
//The dateTimeSeconds data type is an unsigned 32-bit integer in network byte order containing the number of seconds since the UNIX epoch, 1 January 1970 at 00:00 UTC, as defined in [POSIX.1].
//dateTimeSeconds is encoded identically to the IPFIX Message Header Export Time field.
type FieldValueDateTimeSeconds struct {
	value time.Time
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueDateTimeSeconds) MarshalBinary() ([]byte, error) {
	marshalValue := uint32(fv.value.Unix())
	return marshalBinarySingleValue(marshalValue)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueDateTimeSeconds) UnmarshalBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	secondsSinceEpoch := binary.BigEndian.Uint32(data)
	fv.value = time.Unix(int64(secondsSinceEpoch), 0)
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueDateTimeSeconds) Len() uint16 {
	return 4
}

// Value returns FieldValue's value
func (fv *FieldValueDateTimeSeconds) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueDateTimeSeconds) Set(val interface{}) error {
	switch val.(type) {
	case time.Time:
		fv.value = val.(time.Time)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueDateTimeMilliseconds , "dateTimeMilliseconds" represents a time value expressed with millisecond-level precision.
// The dateTimeMilliseconds data type is an unsigned 64-bit integer in network byte order containing the number of milliseconds since the UNIX epoch, 1 January 1970 at 00:00 UTC, as defined in [POSIX.1].
type FieldValueDateTimeMilliseconds struct {
	value time.Time
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueDateTimeMilliseconds) MarshalBinary() ([]byte, error) {
	marshalValue := uint64(fv.value.UnixNano() / 1000000)
	return marshalBinarySingleValue(marshalValue)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueDateTimeMilliseconds) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data))
	}
	milliSecondsSinceEpoch := binary.BigEndian.Uint64(data)
	secondsSinceEpoch := uint64(milliSecondsSinceEpoch) / uint64(1000)
	nanosecondsSinceEpoch := uint64(milliSecondsSinceEpoch) % uint64(1000)
	fv.value = time.Unix(int64(secondsSinceEpoch), int64(nanosecondsSinceEpoch))
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueDateTimeMilliseconds) Len() uint16 {
	return 8
}

// Value returns FieldValue's value
func (fv *FieldValueDateTimeMilliseconds) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueDateTimeMilliseconds) Set(val interface{}) error {
	switch val.(type) {
	case time.Time:
		fv.value = val.(time.Time)
	default:
		return fmt.Errorf("Invalid type for %s", reflect.TypeOf(fv))
	}
	return nil
}

/* */
// FieldValueDateTimeMicroseconds , "dateTimeMicroseconds" represents a time value expressed with microsecond-level precision.
//The dateTimeMicroseconds data type is a 64-bit field encoded according to the NTP Timestamp format as defined in Section 6 of [RFC5905].
//This field is made up of two unsigned 32-bit integers in network byte order: Seconds and Fraction.
//The Seconds field is the number of seconds since the NTP epoch, 1 January 1900 at 00:00 UTC.
//The Fraction field is the fractional number of seconds in units of 1/(2^32) seconds (approximately 233 picoseconds).
//It can represent dates between 1 January 1900 and 8 February 2036 in the current NTP era.
type FieldValueDateTimeMicroseconds struct {
	value time.Time
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueDateTimeMicroseconds) MarshalBinary() ([]byte, error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueDateTimeMicroseconds) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueDateTimeMicroseconds) Len() uint16 {
	return 8
}

// Value returns FieldValue's value
func (fv *FieldValueDateTimeMicroseconds) Value() interface{} {
	return fv.value
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
	value time.Time
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueDateTimeNanoseconds) MarshalBinary() ([]byte, error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueDateTimeNanoseconds) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueDateTimeNanoseconds) Len() uint16 {
	return 8
}

// Value returns FieldValue's value
func (fv *FieldValueDateTimeNanoseconds) Value() interface{} {
	return fv.value
}

/* */
// FieldValueIPv4Address , "ipv4Address" represents an IPv4 address.
type FieldValueIPv4Address struct {
	value net.IP
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueIPv4Address) MarshalBinary() ([]byte, error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueIPv4Address) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueIPv4Address) Len() uint16 {
	return 4
}

// Value returns FieldValue's value
func (fv *FieldValueIPv4Address) Value() interface{} {
	return fv.value
}

/* */
// FieldValueIPv6Address , "ipv6Address" represents an IPv6 address.
type FieldValueIPv6Address struct {
	value net.IP
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueIPv6Address) MarshalBinary() ([]byte, error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueIPv6Address) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueIPv6Address) Len() uint16 {
	return 16
}

// Value returns FieldValue's value
func (fv *FieldValueIPv6Address) Value() interface{} {
	return fv.value
}

/* */
// FieldValueBasicList , "basicList" supports structured data export as described in [RFC6313];
type FieldValueBasicList struct {
	value BasicList
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueBasicList) MarshalBinary() ([]byte, error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueBasicList) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueBasicList) Len() uint16 {
	return fv.value.Len()
}

// Value returns FieldValue's value
func (fv *FieldValueBasicList) Value() interface{} {
	return fv.value
}

/* */
// FieldValueSubTemplateList , "subTemplateList" supports structured data export as described in [RFC6313];
type FieldValueSubTemplateList struct {
	value SubTemplateList
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueSubTemplateList) MarshalBinary() ([]byte, error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueSubTemplateList) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueSubTemplateList) Len() uint16 {
	return fv.value.Len()
}

// Value returns FieldValue's value
func (fv *FieldValueSubTemplateList) Value() interface{} {
	return fv.value
}

/* */
// FieldValueSubTemplateMultiList , "subTemplateMultiList" supports structured data export as described in [RFC6313];
type FieldValueSubTemplateMultiList struct {
	value SubTemplateMultiList
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueSubTemplateMultiList) MarshalBinary() ([]byte, error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueSubTemplateMultiList) UnmarshalBinary(data []byte) error {
	return fmt.Errorf("Not yet implemented!")
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueSubTemplateMultiList) Len() uint16 {
	return fv.value.Len()
}

// Value returns FieldValue's value
func (fv *FieldValueSubTemplateMultiList) Value() interface{} {
	return fv.value
}

/*

Util functions

*/

// EncodeVariableLength returns the bytes for encoding a variable length as specified in RFC 7011, section 7
// In the Template Set, the Information Element Field Length is recorded as 65535.
// This reserved length value notifies the Collecting Process that the length value of the Information Element will be carried in the Information Element content itself.
// In most cases, the length of the Information Element will be less than 255 octets. In this case 1 byte is sufficient to encode the length.
// The length may also be encoded into 3 octets before the Information Element, allowing the length of the Information Element to be greater than or equal to 255 octets.
// In this case, the first octet of the Length field MUST be 255, and the length is carried in the second and third octets.
// The octets carrying the length (either the first or the first three octets) MUST NOT be included in the length of the Information Element.
func EncodeVariableLength(content []byte) ([]byte, error) {
	retval := []byte{}
	if len(content) < 255 {
		retval = []byte{uint8(len(content))}
	} else {
		if len(content) > 65535 {
			return []byte{}, fmt.Errorf("Content too large, maximum of 65535 octets, but it is %d", len(content))
		}
		lengthBytes := []byte{255}
		lengthContentBytes, err := marshalBinarySingleValue(uint16(len(content)))
		if err != nil {
			return []byte{}, err
		}
		retval = append(lengthBytes, lengthContentBytes...)
	}
	return retval, nil
}

// DecodeVariableLength returns the length for decoding a variable length as specified in RFC 7011, section 7
func DecodeVariableLength(content []byte) (uint16, error) {
	retval := uint16(0)
	if content[0] == 0 {
		return 0, fmt.Errorf("Content can not be 0 in length.")
	}
	if content[0] < 255 {
		retval = uint16(content[0])
	} else {
		retval = uint16(256*uint16(content[1])) + uint16(content[2])
	}
	return retval, nil
}
