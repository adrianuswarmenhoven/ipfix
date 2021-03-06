package ipfix

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

var (
	// bufferfiller is used when the exporting process encodes a value in less bytes than real length. See below for explanation
	bufferfiller = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	//Reduced-size encoding MAY be applied to the following integer types:
	//unsigned64, signed64, unsigned32, signed32, unsigned16, and signed16.
	//The signed versus unsigned property of the reported value MUST be preserved.
	//The reduction in size can be to any number of octets smaller than the original type if the data value still fits, i.e., so that only leading zeroes are dropped.  For example, an unsigned64 can be reduced in size to 7, 6, 5, 4, 3, 2, or 1 octet(s).

	//Reduced-size encoding MAY be used to reduce float64 to float32.  The float32 not only has a reduced number range but, due to the smaller mantissa, is also less precise.
	//In this case, the float64 would be reduced in size to 4 octets.

	//Reduced-size encoding MUST NOT be applied to any other data type defined in [RFC7012] that implies a fixed length, as these types either have internal structure (such as ipv4Address or dateTimeMicroseconds) or restricted ranges that are not suitable for reduced-size encoding (such as dateTimeMilliseconds).
)

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
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
	if len(data) < 1 {
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	} else if len(data) < 2 {
		//We prepend 0s (zeroes) if the exporter encoded it in less bytes than we need
		fv.value = binary.BigEndian.Uint16(append(bufferfiller[:1], data...))
	} else {
		fv.value = binary.BigEndian.Uint16(data)
	}
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
	if len(data) < 1 {
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	} else if len(data) < 4 {
		//We prepend 0s if the exporter encoded it in less bytes than we need
		fv.value = binary.BigEndian.Uint32(append(bufferfiller[:4-len(data)], data...))
	} else {
		fv.value = binary.BigEndian.Uint32(data)
	}
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
	}
	return nil
}

/* */
// FieldValueUnsigned64 , "unsigned64" represents a non-negative integer value in the range of 0 to 18446744073709551615
type FieldValueUnsigned64 struct {
	value uint64
} // The TemplateRecord points to the field types

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueUnsigned64) MarshalBinary() ([]byte, error) {
	return marshalBinarySingleValue(fv.value)
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueUnsigned64) UnmarshalBinary(data []byte) error {
	if len(data) < 1 {
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	} else if len(data) < 8 {
		//We prepend 0s if the exporter encoded it in less bytes than we need
		fv.value = binary.BigEndian.Uint64(append(bufferfiller[:8-len(data)], data...))
	} else {
		fv.value = binary.BigEndian.Uint64(data)
	}
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
	if len(data) < 1 {
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	} else if len(data) < 2 {
		//We prepend 0s if the exporter encoded it in less bytes than we need
		fv.value = int16(binary.BigEndian.Uint16(append(bufferfiller[:2-len(data)], data...)))
	} else {
		fv.value = int16(binary.BigEndian.Uint16(data))
	}
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
	if len(data) < 1 {
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	} else if len(data) < 4 {
		//We prepend 0s if the exporter encoded it in less bytes than we need
		fv.value = int32(binary.BigEndian.Uint32(append(bufferfiller[:4-len(data)], data...)))
	} else {
		fv.value = int32(binary.BigEndian.Uint32(data))
	}
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
	if len(data) < 1 {
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	} else if len(data) < 8 {
		//We prepend 0s if the exporter encoded it in less bytes than we need
		fv.value = int64(binary.BigEndian.Uint64(append(bufferfiller[:8-len(data)], data...)))
	} else {
		fv.value = int64(binary.BigEndian.Uint64(data))
	}
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
	if len(data) < 4 {
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	} else if len(data) == 4 {
		fv.value = float64(math.Float32frombits(binary.BigEndian.Uint32(data)))
	} else if len(data) == 8 {
		fv.value = math.Float64frombits(binary.BigEndian.Uint64(data))
	} else {
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	}
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	}
	switch data[0] {
	case 1:
		fv.value = true
	case 2:
		fv.value = false
	default:
		return NewError(fmt.Sprintf("Invalid encoded value for boolean: %d, must be either 1 or 2", data[0]), ErrCritical)
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	}
	tmpval := make([]byte, len(data))
	err := unmarshalBinaryOctets(data, tmpval)
	if err != nil {
		return err
	}
	fv.value = net.HardwareAddr(tmpval)
	return nil
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
	marshalValue, err := marshalBinarySingleValue(fv.value)
	if err != nil {
		return []byte{}, err
	}
	return marshalValue, nil
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
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
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
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
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
	}
	return nil
}

/* */
const (
	epochDelta = 2208988800 // unix epoch = seconds since January 1, 1970 UTC,ntp prime epoch = seconds since January 1, 1900 UTC
)

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
	marshalValueSeconds := fv.value.Unix() + int64(epochDelta)
	marshalValueFractions := ((uint64(fv.value.UnixNano()) % 1e9) << 32) / 1e9

	marshalSec, err := marshalBinarySingleValue(uint32(marshalValueSeconds))
	if err != nil {
		return nil, err
	}
	marshalFrac, err := marshalBinarySingleValue(uint32(marshalValueFractions))
	if err != nil {
		return nil, err
	}

	marshalValue := make([]byte, 0, 0)
	marshalValue = append(marshalValue, marshalSec...)
	marshalValue = append(marshalValue, marshalFrac...)
	if len(marshalValue) != 8 {
		return nil, NewError(fmt.Sprintf("Incorrect length when marshalling. Wanted %d, got %d.", 8, len(marshalValue)), ErrCritical)
	}
	return marshalValue, nil
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueDateTimeMicroseconds) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	}
	baseValueSeconds := int64(binary.BigEndian.Uint32(data[:4])) - int64(epochDelta)
	baseValueFractions := (int64(1+binary.BigEndian.Uint32(data[4:])) * 1e9) >> 32 //Yeah... that offset... for some reason it is necessary
	fv.value = time.Unix(baseValueSeconds, baseValueFractions)
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueDateTimeMicroseconds) Len() uint16 {
	return 8
}

// Value returns FieldValue's value
func (fv *FieldValueDateTimeMicroseconds) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueDateTimeMicroseconds) Set(val interface{}) error {
	switch val.(type) {
	case time.Time:
		fv.value = val.(time.Time)
	default:
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
	}
	return nil
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
	marshalValueSeconds := fv.value.Unix() + int64(epochDelta)
	marshalValueFractions := ((uint64(fv.value.UnixNano()) % 1e9) << 32) / 1e9 //We only want the remainder of nanoseconds

	marshalSec, err := marshalBinarySingleValue(uint32(marshalValueSeconds))
	if err != nil {
		return nil, err
	}
	marshalFrac, err := marshalBinarySingleValue(uint32(marshalValueFractions))
	if err != nil {
		return nil, err
	}

	marshalValue := make([]byte, 0, 0)
	marshalValue = append(marshalValue, marshalSec...)
	marshalValue = append(marshalValue, marshalFrac...)
	if len(marshalValue) != 8 {
		return nil, NewError(fmt.Sprintf("Incorrect length when marshalling. Wanted %d, got %d.", 8, len(marshalValue)), ErrCritical)
	}
	return marshalValue, nil
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueDateTimeNanoseconds) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return NewError(fmt.Sprintf("Insufficient data. Need length %d, but got %d.", fv.Len(), len(data)), ErrCritical)
	}
	baseValueSeconds := int64(binary.BigEndian.Uint32(data[:4])) - int64(epochDelta)
	baseValueFractions := (int64(1+binary.BigEndian.Uint32(data[4:])) * 1e9) >> 32 //Yeah... that offset... for some reason it is necessary
	fv.value = time.Unix(baseValueSeconds, baseValueFractions)
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueDateTimeNanoseconds) Len() uint16 {
	return 8
}

// Value returns FieldValue's value
func (fv *FieldValueDateTimeNanoseconds) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueDateTimeNanoseconds) Set(val interface{}) error {
	switch val.(type) {
	case time.Time:
		fv.value = val.(time.Time)
	default:
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
	}
	return nil
}

/* */
// FieldValueIPv4Address , "ipv4Address" represents an IPv4 address.
type FieldValueIPv4Address struct {
	value net.IP
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueIPv4Address) MarshalBinary() ([]byte, error) {
	marshalValue, err := marshalBinarySingleValue(fv.value.To4())
	if err != nil {
		return []byte{}, err
	}
	return marshalValue, nil
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueIPv4Address) UnmarshalBinary(data []byte) error {
	fv.value = make([]byte, len(data))
	return unmarshalBinaryOctets(data, &fv.value)
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueIPv4Address) Len() uint16 {
	return 4
}

// Value returns FieldValue's value
func (fv *FieldValueIPv4Address) Value() interface{} {
	return fv.value.To4()
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueIPv4Address) Set(val interface{}) error {
	switch val.(type) {
	case string:
		tmpip := net.ParseIP(val.(string))
		if tmpip != nil {
			fv.value = tmpip
		} else {
			return NewError(fmt.Sprintf("Value is not an IP Address: %s", val.(string)), ErrCritical)
		}
	case net.IP:
		fv.value = val.(net.IP)
	default:
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
	}
	return nil
}

/* */
// FieldValueIPv6Address , "ipv6Address" represents an IPv6 address.
type FieldValueIPv6Address struct {
	value net.IP
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueIPv6Address) MarshalBinary() ([]byte, error) {
	marshalValue, err := marshalBinarySingleValue(fv.value.To16())
	if err != nil {
		return []byte{}, err
	}
	return marshalValue, nil
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueIPv6Address) UnmarshalBinary(data []byte) error {
	fv.value = make([]byte, len(data))
	return unmarshalBinaryOctets(data, &fv.value)
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueIPv6Address) Len() uint16 {
	return 16
}

// Value returns FieldValue's value
func (fv *FieldValueIPv6Address) Value() interface{} {
	return fv.value.To16()
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueIPv6Address) Set(val interface{}) error {
	switch val.(type) {
	case string:
		tmpip := net.ParseIP(val.(string))
		if tmpip != nil {
			fv.value = tmpip
		} else {
			return NewError(fmt.Sprintf("Value is not an IP Address: %s", val.(string)), ErrCritical)
		}
	case net.IP:
		fv.value = val.(net.IP)
	default:
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
	}
	return nil
}

/* */
// FieldValueBasicList , "basicList" supports structured data export as described in [RFC6313];
type FieldValueBasicList struct {
	value BasicList
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueBasicList) MarshalBinary() ([]byte, error) {
	marshalValue := make([]byte, 0, 0)
	marshalValue = append(marshalValue, fv.value.Semantic)
	marshalFieldID, err := marshalBinarySingleValue(fv.value.InformationElementIdentifier)
	if err != nil {
		return nil, err
	}
	if fv.value.E {
		marshalFieldID[0] = marshalFieldID[0] | 128 //Setting the EnterpriseID bit
	}
	marshalValue = append(marshalValue, marshalFieldID...)
	marshalFieldLength, err := marshalBinarySingleValue(fv.value.FieldLength)
	if err != nil {
		return nil, err
	}
	marshalValue = append(marshalValue, marshalFieldLength...)
	if fv.value.E {
		marshalEnterpriseID, err := marshalBinarySingleValue(fv.value.EnterpriseNumber)
		if err != nil {
			return nil, err
		}
		marshalValue = append(marshalValue, marshalEnterpriseID...)
	}

	for _, listitem := range fv.value.FieldValues {
		itemdata, err := listitem.MarshalBinary()
		if err != nil {
			return nil, err
		}
		if fv.value.FieldLength == VariableLength {
			var marshalLength []byte
			switch listitem.(type) {
			case *FieldValueBasicList, *FieldValueSubTemplateList, *FieldValueSubTemplateMultiList:
				marshalLength, err = EncodeVariableLength(itemdata, true)
				if err != nil {
					return nil, err
				}

			default:
				marshalLength, err = EncodeVariableLength(itemdata, false)
				if err != nil {
					return nil, err
				}
			}
			marshalValue = append(marshalValue, marshalLength...)
		}
		marshalValue = append(marshalValue, itemdata...)
	}

	return marshalValue, nil
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueBasicList) UnmarshalBinary(data []byte) error {
	fv.value = BasicList{}
	fv.value.FieldValues = make([]FieldValue, 0, 0)

	fv.value.Semantic = data[0]
	if (data[1] & 128) != 0 {
		fv.value.E = true
		data[1] = data[1] & 127 //Remove the bit
	}
	fv.value.InformationElementIdentifier = binary.BigEndian.Uint16(data[1:3])
	if fv.value.E {
		data[1] = data[1] | 128 //Restore the bit (so we keep the original datablob)
	}
	fv.value.FieldLength = binary.BigEndian.Uint16(data[3:5])

	cursor := int(5)
	if fv.value.E {
		fv.value.EnterpriseNumber = binary.BigEndian.Uint32(data[cursor : cursor+4])
		cursor += 4
	}
	for fv.value.FieldLength != VariableLength && cursor+int(fv.value.FieldLength) <= len(data) ||
		(fv.value.FieldLength == VariableLength && cursor < len(data)) {
		if fv.value.FieldLength == VariableLength {
			fieldlength, cursorshift, err := DecodeVariableLength(data[cursor : cursor+3]) //Max. 3 positions cursor shift
			if err != nil {
				return err
			}
			cursor += int(cursorshift)
			newval, err := NewFieldValueByID(fv.value.EnterpriseNumber, fv.value.InformationElementIdentifier)
			if err != nil {
				return err
			}
			err = newval.UnmarshalBinary(data[cursor : cursor+int(fieldlength)])
			if err != nil {
				return err
			}
			cursor += int(fieldlength)
			fv.value.FieldValues = append(fv.value.FieldValues, newval)
		} else {
			newval, err := NewFieldValueByID(fv.value.EnterpriseNumber, fv.value.InformationElementIdentifier)
			if err != nil {
				return err
			}
			err = newval.UnmarshalBinary(data[cursor : cursor+int(fv.value.FieldLength)])
			if err != nil {
				return err
			}
			cursor += int(fv.value.FieldLength)
			fv.value.FieldValues = append(fv.value.FieldValues, newval)
		}
	}

	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueBasicList) Len() uint16 {
	return fv.value.Len()
}

// Value returns FieldValue's value
func (fv *FieldValueBasicList) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueBasicList) Set(val interface{}) error {
	switch val.(type) {
	case BasicList:
		fv.value = val.(BasicList)
	default:
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
	}
	return nil
}

/* */
// FieldValueSubTemplateList , "subTemplateList" supports structured data export as described in [RFC6313];
type FieldValueSubTemplateList struct {
	value SubTemplateList
}

// SetAssiocatedTemplates sets the list of templates belonging to this session
func (fv *FieldValueSubTemplateList) SetAssiocatedTemplates(at *ActiveTemplates) error {
	if at == nil {
		return NewError("Can not set associated templates to nil", ErrCritical)
	}
	fv.value.AssociatedTemplates = at
	return nil
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueSubTemplateList) MarshalBinary() ([]byte, error) {
	if fv.value.AssociatedTemplates == nil {
		return nil, NewError("Can not marshal without associated templates", ErrCritical)
	}
	if fv.value.TemplateID < 256 {
		return nil, NewError("Can not marshal without a template id", ErrCritical)
	}
	marshalValue := make([]byte, 0, 0)
	marshalValue = append(marshalValue, fv.value.Semantic)

	marshalTemplateID, err := marshalBinarySingleValue(fv.value.TemplateID)
	if err != nil {
		return nil, err
	}
	marshalValue = append(marshalValue, marshalTemplateID...)
	for _, listitem := range fv.value.Records {
		listitem.(*DataRecord).AssociateTemplates(fv.value.AssociatedTemplates)
		listitem.(*DataRecord).SetTemplateID(fv.value.TemplateID)
		recordBinary, err := listitem.(*DataRecord).MarshalBinary()
		if err != nil {
			return nil, err
		}
		marshalValue = append(marshalValue, recordBinary...)
	}

	return marshalValue, nil
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueSubTemplateList) UnmarshalBinary(data []byte) error {
	if fv.value.AssociatedTemplates == nil {
		return NewError("Can not marshal without associated templates", ErrFailure) //This is a failure and not critical because we can re-do later
	}
	if data == nil || len(data) == 0 {
		return NewError(fmt.Sprintf("Can not unmarshal, invalid data. %#v", data), ErrCritical)
	}

	fv.value = SubTemplateList{AssociatedTemplates: fv.value.AssociatedTemplates, TemplateID: fv.value.TemplateID} //Create a clean copy with correct data, may not be necessary
	fv.value.Records = make([]Record, 0, 0)

	fv.value.Semantic = data[0]
	fv.value.TemplateID = binary.BigEndian.Uint16(data[1:3])
	if fv.value.TemplateID < 256 {
		return NewError(fmt.Sprintf("Can not marshal without a template id"), ErrCritical)
	}
	cursor := uint16(3)
	for cursor < uint16(len(data)) {
		newdatrec := &DataRecord{
			AssociatedTemplates: fv.value.AssociatedTemplates,
			TemplateID:          fv.value.TemplateID,
			FieldValues:         make([]FieldValue, 0, 0),
		}
		err := newdatrec.UnmarshalBinary(data[cursor:])
		if err != nil {
			return err
		}
		fv.value.Records = append(fv.value.Records, newdatrec)
		cursor += newdatrec.Len()
	}
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueSubTemplateList) Len() uint16 {
	return fv.value.Len()
}

// Value returns FieldValue's value
func (fv *FieldValueSubTemplateList) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueSubTemplateList) Set(val interface{}) error {
	switch val.(type) {
	case FieldValueSubTemplateList:
		fv.value = val.(SubTemplateList)
	default:
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
	}
	return nil
}

/* */
// FieldValueSubTemplateMultiList , "subTemplateMultiList" supports structured data export as described in [RFC6313];
type FieldValueSubTemplateMultiList struct {
	value SubTemplateMultiList
}

// SetAssiocatedTemplates sets the list of templates belonging to this session
func (fv *FieldValueSubTemplateMultiList) SetAssiocatedTemplates(at *ActiveTemplates) error {
	if at == nil {
		return NewError(fmt.Sprintf("Can not set associated templates to nil"), ErrCritical)
	}
	fv.value.AssociatedTemplates = at
	return nil
}

// MarshalBinary returns the Network Byte Order byte representation of this Field Value
func (fv *FieldValueSubTemplateMultiList) MarshalBinary() ([]byte, error) {
	if fv.value.AssociatedTemplates == nil {
		return nil, NewError("Can not marshal without associated templates", ErrFailure) //Failure because we may be able to do this later
	}
	marshalValue := make([]byte, 0, 0)
	marshalValue = append(marshalValue, fv.value.Semantic)
	for idx, subtpldat := range fv.value.SubTemplates {
		if subtpldat.AssociatedTemplates == nil {
			subtpldat.AssociateTemplates(fv.value.AssociatedTemplates)
		}
		if subtpldat.TemplateID < 256 {
			return nil, NewError(fmt.Sprintf("Can not marshal without a template id. Error in sub template %d (%#v)", idx, *subtpldat), ErrCritical)
		}

		marshalTemplateID, err := marshalBinarySingleValue(subtpldat.TemplateID)
		if err != nil {
			return nil, err
		}
		marshalValue = append(marshalValue, marshalTemplateID...)

		// Data Records Length
		// This is the total length of the Data Records encoding for the Template ID previously specified, including the two bytes for the Template ID and the two bytes for the Data Records Length field itself.
		// In the exceptional case of zero instances in the subTemplateMultiList, no data is encoded, only the Semantic field and Template ID field(s), and the Data Record Length field is set to zero.
		enclen := uint16(subtpldat.Len() + uint16(4))
		if len(subtpldat.Records) == 0 {
			enclen = uint16(0)
		}
		marshalLen, err := marshalBinarySingleValue(enclen)
		if err != nil {
			return nil, err
		}
		marshalValue = append(marshalValue, marshalLen...)

		for _, listitem := range subtpldat.Records {
			listitem.(*DataRecord).AssociateTemplates(subtpldat.AssociatedTemplates)
			listitem.(*DataRecord).SetTemplateID(subtpldat.TemplateID)
			recordBinary, err := listitem.(*DataRecord).MarshalBinary()
			if err != nil {
				return nil, err
			}
			marshalValue = append(marshalValue, recordBinary...)
		}
	}
	return marshalValue, nil
}

// UnmarshalBinary fills the value from Network Byte Order byte representation
func (fv *FieldValueSubTemplateMultiList) UnmarshalBinary(data []byte) error {
	if fv.value.AssociatedTemplates == nil {
		return NewError("Can not marshal without associated templates", ErrFailure) //Failure because we may be able to do this later
	}
	if data == nil || len(data) == 0 {
		return NewError(fmt.Sprintf("Can not unmarshal, invalid data. %#v", data), ErrCritical)
	}

	fv.value = SubTemplateMultiList{AssociatedTemplates: fv.value.AssociatedTemplates, SubTemplates: make([]*SubTemplateData, 0, 0)} //Create a clean copy with correct data, may not be necessary

	cursor := uint16(1)
	for cursor < uint16(len(data)) {
		newtplid := binary.BigEndian.Uint16(data[cursor : cursor+2])
		if newtplid < 256 {
			return NewError("Can not unmarshal without a proper template id", ErrCritical)
		}
		cursor += 2
		newtpllen := binary.BigEndian.Uint16(data[cursor : cursor+2])
		newsubtemplate, err := NewSubTemplateData(newtplid)
		if err != nil {
			return err
		}
		cursor += 2
		if newtpllen > 4 { //Length is including template and length itself
			newsubtemplate.AssociateTemplates(fv.value.AssociatedTemplates)
			for cursor < newtpllen-4 {
				newdatrec := &DataRecord{
					AssociatedTemplates: fv.value.AssociatedTemplates,
					TemplateID:          newtplid,
					FieldValues:         make([]FieldValue, 0, 0),
				}
				err := newdatrec.UnmarshalBinary(data[cursor:])
				if err != nil {
					return err
				}
				newsubtemplate.Records = append(newsubtemplate.Records, newdatrec)
				cursor += newdatrec.Len()
			}
			cursor += newtpllen
		}
		fv.value.SubTemplates = append(fv.value.SubTemplates, newsubtemplate)
	}
	return nil
}

// Len returns the number of octets this FieldValue is wide
func (fv *FieldValueSubTemplateMultiList) Len() uint16 {
	return fv.value.Len()
}

// Value returns FieldValue's value
func (fv *FieldValueSubTemplateMultiList) Value() interface{} {
	return fv.value
}

// Set sets the FieldValue's value. Will return an error if the type is incorrect.
func (fv *FieldValueSubTemplateMultiList) Set(val interface{}) error {
	switch val.(type) {
	case FieldValueSubTemplateMultiList:
		fv.value = val.(SubTemplateMultiList)
	default:
		return NewError(fmt.Sprintf("Invalid type for %s", reflect.TypeOf(fv)), ErrCritical)
	}
	return nil
}
