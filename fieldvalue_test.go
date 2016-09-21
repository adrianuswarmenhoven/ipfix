package ipfixmessage

import (
	"fmt"
	"math"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestFieldValueValueGoTypes(t *testing.T) {
	testsetmatch := map[interface{}]int{
		FieldValueUnsigned8{}.value:  tg_uint8,
		FieldValueUnsigned16{}.value: tg_uint16,
		FieldValueUnsigned32{}.value: tg_uint32,
		FieldValueUnsigned64{}.value: tg_uint64,

		FieldValueSigned8{}.value:  tg_int8,
		FieldValueSigned16{}.value: tg_int16,
		FieldValueSigned32{}.value: tg_int32,
		FieldValueSigned64{}.value: tg_int64,

		FieldValueFloat32{}.value: tg_float32,
		FieldValueFloat64{}.value: tg_float64,

		FieldValueBoolean{}.value: tg_bool,

		FieldValueString{}.value: tg_string,

		FieldValueDateTimeSeconds{}.value:      tg_time,
		FieldValueDateTimeMilliseconds{}.value: tg_time,
		FieldValueDateTimeMicroseconds{}.value: tg_time,
		FieldValueDateTimeNanoseconds{}.value:  tg_time,
	}
	for testval, checkval := range testsetmatch {
		if goTypeName(testval) != checkval {
			t.Errorf("Should have gotten %d but got %d for %#v", checkval, goTypeName(testval), testval)
		}
	}

	//Now checking the unhashables
	if reflect.TypeOf(FieldValueMacAddress{}.value) != reflect.TypeOf(net.HardwareAddr{}) {
		t.Errorf("Should have gotten %s but got %s for %#v", reflect.TypeOf(net.HardwareAddr{}), reflect.TypeOf(FieldValueMacAddress{}.value), FieldValueMacAddress{}.value)
	}
	if reflect.TypeOf(FieldValueIPv4Address{}.value) != reflect.TypeOf(net.IP{}) {
		t.Errorf("Should have gotten %s but got %s for %#v", reflect.TypeOf(net.IP{}), reflect.TypeOf(FieldValueIPv4Address{}.value), FieldValueIPv4Address{}.value)
	}
	if reflect.TypeOf(FieldValueIPv6Address{}.value) != reflect.TypeOf(net.IP{}) {
		t.Errorf("Should have gotten %s but got %s for %#v", reflect.TypeOf(net.IP{}), reflect.TypeOf(FieldValueIPv6Address{}.value), FieldValueIPv6Address{}.value)
	}
	if reflect.TypeOf(FieldValueOctetArray{}.value) != reflect.TypeOf([]byte{}) {
		t.Errorf("Should have gotten %s but got %s for %#v", reflect.TypeOf([]byte{}), reflect.TypeOf(FieldValueOctetArray{}.value), FieldValueOctetArray{}.value)
	}

	//We don't check the RFC 6313 values here since they are complex types and all end up using the above Field Values anyway
}

type fieldvalueSetGetTestcase struct {
	TestVal     FieldValue
	CompVal     interface{}
	MustFail    bool
	ByteCompare bool
}

func TestFieldValueSetGet(t *testing.T) {
	var testset = []fieldvalueSetGetTestcase{
		0: {TestVal: &FieldValueUnsigned8{value: 0}, CompVal: uint16(42), MustFail: true, ByteCompare: false},
		1: {TestVal: &FieldValueUnsigned8{value: 0}, CompVal: uint8(42), MustFail: false, ByteCompare: false},
		2: {TestVal: &FieldValueUnsigned16{value: 0}, CompVal: uint8(42), MustFail: true, ByteCompare: false},
		3: {TestVal: &FieldValueUnsigned16{value: 0}, CompVal: uint16(42), MustFail: false, ByteCompare: false},
		4: {TestVal: &FieldValueUnsigned32{value: 0}, CompVal: uint16(42), MustFail: true, ByteCompare: false},
		5: {TestVal: &FieldValueUnsigned32{value: 0}, CompVal: uint32(42), MustFail: false, ByteCompare: false},
		6: {TestVal: &FieldValueUnsigned64{value: 0}, CompVal: uint16(42), MustFail: true, ByteCompare: false},
		7: {TestVal: &FieldValueUnsigned64{value: 0}, CompVal: uint64(42), MustFail: false, ByteCompare: false},

		8:  {TestVal: &FieldValueSigned8{value: 0}, CompVal: int16(42), MustFail: true, ByteCompare: false},
		9:  {TestVal: &FieldValueSigned8{value: 0}, CompVal: int8(42), MustFail: false, ByteCompare: false},
		10: {TestVal: &FieldValueSigned16{value: 0}, CompVal: int8(42), MustFail: true, ByteCompare: false},
		11: {TestVal: &FieldValueSigned16{value: 0}, CompVal: int16(42), MustFail: false, ByteCompare: false},
		12: {TestVal: &FieldValueSigned32{value: 0}, CompVal: int16(42), MustFail: true, ByteCompare: false},
		13: {TestVal: &FieldValueSigned32{value: 0}, CompVal: int32(42), MustFail: false, ByteCompare: false},
		14: {TestVal: &FieldValueSigned64{value: 0}, CompVal: int16(42), MustFail: true, ByteCompare: false},
		15: {TestVal: &FieldValueSigned64{value: 0}, CompVal: int64(42), MustFail: false, ByteCompare: false},

		16: {TestVal: &FieldValueFloat32{value: 0}, CompVal: float64(math.Phi), MustFail: true, ByteCompare: false},
		17: {TestVal: &FieldValueFloat32{value: 0}, CompVal: float32(math.Phi), MustFail: false, ByteCompare: false},
		18: {TestVal: &FieldValueFloat64{value: 0}, CompVal: float32(math.Pi), MustFail: true, ByteCompare: false},
		19: {TestVal: &FieldValueFloat64{value: 0}, CompVal: float64(math.Pi), MustFail: false, ByteCompare: false},

		20: {TestVal: &FieldValueBoolean{value: false}, CompVal: int32(4), MustFail: true, ByteCompare: false},
		21: {TestVal: &FieldValueBoolean{value: false}, CompVal: bool(true), MustFail: false, ByteCompare: false},

		22: {TestVal: &FieldValueMacAddress{value: net.HardwareAddr{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}}, CompVal: int32(4), MustFail: true, ByteCompare: true},
		23: {TestVal: &FieldValueMacAddress{value: net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}, CompVal: net.HardwareAddr{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}, MustFail: false, ByteCompare: true},
	}

	for _, testcase := range testset {
		err := testcase.TestVal.Set(testcase.CompVal)
		if (err != nil) != testcase.MustFail {
			t.Errorf("Testcase did not react correctly. Wanted fail(%v) but got fail(%v) for testcase %#v", testcase.MustFail, (err != nil), testcase)
		}
		if !testcase.MustFail {
			if !testcase.ByteCompare {
				if testcase.TestVal.Value() != testcase.CompVal {
					t.Errorf("Wrong value returned. Wanted %d but got %d for testcase %#v", testcase.CompVal, testcase.TestVal.Value(), testcase)
				}
			} else {
				if fmt.Sprintf("%v", testcase.TestVal.Value()) != fmt.Sprintf("%v", testcase.CompVal) {
					t.Errorf("Wrong value returned. Wanted %v but got %v for testcase %#v", testcase.CompVal, testcase.TestVal.Value(), testcase)
				}
			}
		}
	}
}

type fieldvalueMarshalUnmarshalTestcase struct {
	SourceVal FieldValue
	DestVal   FieldValue
	CompVal   interface{}
}

func TestFieldValueMarshalUnmarshal(t *testing.T) {
	var testset = []fieldvalueMarshalUnmarshalTestcase{
		0: {SourceVal: &FieldValueUnsigned8{value: 0}, DestVal: &FieldValueUnsigned8{value: 0}, CompVal: uint8(0)},
		1: {SourceVal: &FieldValueUnsigned8{value: math.MaxUint8}, DestVal: &FieldValueUnsigned8{value: 0}, CompVal: uint8(math.MaxUint8)},
		2: {SourceVal: &FieldValueUnsigned16{value: 0}, DestVal: &FieldValueUnsigned16{value: 0}, CompVal: uint16(0)},
		3: {SourceVal: &FieldValueUnsigned16{value: math.MaxUint16}, DestVal: &FieldValueUnsigned16{value: 0}, CompVal: uint16(math.MaxUint16)},
		4: {SourceVal: &FieldValueUnsigned32{value: 0}, DestVal: &FieldValueUnsigned32{value: 0}, CompVal: uint32(0)},
		5: {SourceVal: &FieldValueUnsigned32{value: math.MaxUint32}, DestVal: &FieldValueUnsigned32{value: 0}, CompVal: uint32(math.MaxUint32)},
		6: {SourceVal: &FieldValueUnsigned64{value: 0}, DestVal: &FieldValueUnsigned64{value: 0}, CompVal: uint64(0)},
		7: {SourceVal: &FieldValueUnsigned64{value: math.MaxUint64}, DestVal: &FieldValueUnsigned64{value: 0}, CompVal: uint64(math.MaxUint64)},

		8:  {SourceVal: &FieldValueSigned8{value: math.MinInt8}, DestVal: &FieldValueSigned8{value: 0}, CompVal: int8(math.MinInt8)},
		9:  {SourceVal: &FieldValueSigned8{value: math.MaxInt8}, DestVal: &FieldValueSigned8{value: 0}, CompVal: int8(math.MaxInt8)},
		10: {SourceVal: &FieldValueSigned16{value: math.MinInt16}, DestVal: &FieldValueSigned16{value: 0}, CompVal: int16(math.MinInt16)},
		11: {SourceVal: &FieldValueSigned16{value: math.MaxInt16}, DestVal: &FieldValueSigned16{value: 0}, CompVal: int16(math.MaxInt16)},
		12: {SourceVal: &FieldValueSigned32{value: math.MinInt32}, DestVal: &FieldValueSigned32{value: 0}, CompVal: int32(math.MinInt32)},
		13: {SourceVal: &FieldValueSigned32{value: math.MaxInt32}, DestVal: &FieldValueSigned32{value: 0}, CompVal: int32(math.MaxInt32)},
		14: {SourceVal: &FieldValueSigned64{value: math.MinInt64}, DestVal: &FieldValueSigned64{value: 0}, CompVal: int64(math.MinInt64)},
		15: {SourceVal: &FieldValueSigned64{value: math.MaxInt64}, DestVal: &FieldValueSigned64{value: 0}, CompVal: int64(math.MaxInt64)},

		16: {SourceVal: &FieldValueFloat32{value: math.MaxFloat32}, DestVal: &FieldValueFloat32{value: 0}, CompVal: float32(math.MaxFloat32)},
		17: {SourceVal: &FieldValueFloat32{value: math.SmallestNonzeroFloat32}, DestVal: &FieldValueFloat32{value: 0}, CompVal: float32(math.SmallestNonzeroFloat32)},
		18: {SourceVal: &FieldValueFloat64{value: math.MaxFloat64}, DestVal: &FieldValueFloat64{value: 0}, CompVal: float64(math.MaxFloat64)},
		19: {SourceVal: &FieldValueFloat64{value: math.SmallestNonzeroFloat64}, DestVal: &FieldValueFloat64{value: 0}, CompVal: float64(math.SmallestNonzeroFloat64)},

		20: {SourceVal: &FieldValueBoolean{value: false}, DestVal: &FieldValueBoolean{value: true}, CompVal: bool(false)},
		21: {SourceVal: &FieldValueBoolean{value: true}, DestVal: &FieldValueBoolean{value: false}, CompVal: bool(true)},
	}

	for _, testcase := range testset {
		binarydata, err := testcase.SourceVal.MarshalBinary()
		if err != nil {
			t.Errorf("Error marshalling %#v: %#v", testcase.SourceVal, err)
		}
		if len(binarydata) != int(testcase.SourceVal.Len()) {
			t.Errorf("Error marshalling %#v: length of binary data should be %d, but was %d", testcase.SourceVal, testcase.SourceVal.Len(), len(binarydata))
		}
		err = testcase.DestVal.UnmarshalBinary(binarydata)
		if err != nil {
			t.Errorf("Error unmarshalling %#v: %#v", testcase.SourceVal, err)
		}
		if !reflect.DeepEqual(testcase.SourceVal, testcase.DestVal) {
			t.Errorf("Error in value after conversions, wanted %#v, but got %#v", testcase.SourceVal, testcase.DestVal)
		}
		if testcase.DestVal.Value() != testcase.CompVal {
			t.Errorf("Error in value after conversions, wanted %#v, but got %#v", testcase.CompVal, testcase.DestVal)
		}
	}
}

//goTypeName is a helper function
func goTypeName(fv interface{}) int {
	switch fv.(type) {
	case uint8:
		return tg_uint8
	case uint16:
		return tg_uint16
	case uint32:
		return tg_uint32
	case uint64:
		return tg_uint64
	case int8:
		return tg_int8
	case int16:
		return tg_int16
	case int32:
		return tg_int32
	case int64:
		return tg_int64
	case float32:
		return tg_float32
	case float64:
		return tg_float64
	case bool:
		return tg_bool
	case net.HardwareAddr:
		return tg_mac
	case []byte:
		return tg_byteslice
	case string:
		return tg_string
	case time.Time:
		return tg_time
	case net.IP:
		return tg_ipaddress
	default:
		return t_unknown
	}
}
