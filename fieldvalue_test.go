package ipfixmessage

import (
	"bytes"
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

var (
	dateA, _ = time.Parse(time.RFC822, "20 Jan 14 22:45 CET")
	dateB, _ = time.Parse(time.RFC822, "09 Oct 13 10:00 CET")
	//dateC, _ = time.Parse(time.RFC822, "07 Dec 70 10:00 CET")
	dateC = time.Unix(int64(29408400), int64(502219461))
)

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

		24: {TestVal: &FieldValueOctetArray{value: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}}, CompVal: int32(13), MustFail: true, ByteCompare: true},
		25: {TestVal: &FieldValueOctetArray{value: []byte("teststring")}, CompVal: []byte("teststring"), MustFail: false, ByteCompare: true},

		26: {TestVal: &FieldValueString{value: "To be or not to be, is that even questionable?"}, CompVal: net.HardwareAddr{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}, MustFail: true, ByteCompare: true},
		27: {TestVal: &FieldValueString{value: "一帆风顺"}, CompVal: "一帆风顺", MustFail: false, ByteCompare: true},

		28: {TestVal: &FieldValueDateTimeSeconds{value: dateB}, CompVal: dateA, MustFail: false, ByteCompare: true},
		29: {TestVal: &FieldValueDateTimeSeconds{value: dateB}, CompVal: uint64(0), MustFail: true, ByteCompare: true},

		30: {TestVal: &FieldValueDateTimeMilliseconds{value: dateB}, CompVal: dateA, MustFail: false, ByteCompare: true},
		31: {TestVal: &FieldValueDateTimeMilliseconds{value: dateA}, CompVal: uint64(0), MustFail: true, ByteCompare: true},

		32: {TestVal: &FieldValueDateTimeMicroseconds{value: dateB}, CompVal: dateA, MustFail: false, ByteCompare: true},
		33: {TestVal: &FieldValueDateTimeMicroseconds{value: dateA}, CompVal: uint64(0), MustFail: true, ByteCompare: true},

		34: {TestVal: &FieldValueDateTimeNanoseconds{value: dateB}, CompVal: dateA, MustFail: false, ByteCompare: true},
		35: {TestVal: &FieldValueDateTimeNanoseconds{value: dateA}, CompVal: uint64(0), MustFail: true, ByteCompare: true},

		36: {TestVal: &FieldValueIPv4Address{value: net.IP{}}, CompVal: "127.0.0.1", MustFail: false, ByteCompare: true},
		37: {TestVal: &FieldValueIPv4Address{value: net.IP{}}, CompVal: "An ip address", MustFail: true, ByteCompare: true},

		38: {TestVal: &FieldValueIPv6Address{value: net.IP{}}, CompVal: "2001:db8::68", MustFail: false, ByteCompare: true},
		39: {TestVal: &FieldValueIPv6Address{value: net.IP{}}, CompVal: "An ip address", MustFail: true, ByteCompare: true},
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

type fieldvalueMarshalEncodingTestcase struct {
	SourceVal      FieldValue
	CompEncoded    []byte
	VariableLength bool
}

func TestMarshalEncoding(t *testing.T) {
	var testset = []fieldvalueMarshalEncodingTestcase{
		0: {SourceVal: &FieldValueUnsigned8{value: 42}, CompEncoded: []byte{42}},
		1: {SourceVal: &FieldValueUnsigned16{value: 0x0102}, CompEncoded: []byte{1, 2}},
		2: {SourceVal: &FieldValueUnsigned16{value: 0x0201}, CompEncoded: []byte{2, 1}},
		3: {SourceVal: &FieldValueUnsigned32{value: 0x01020304}, CompEncoded: []byte{1, 2, 3, 4}},
		4: {SourceVal: &FieldValueUnsigned32{value: 0x04030201}, CompEncoded: []byte{4, 3, 2, 1}},
		5: {SourceVal: &FieldValueUnsigned64{value: 0x0102030405060708}, CompEncoded: []byte{1, 2, 3, 4, 5, 6, 7, 8}},
		6: {SourceVal: &FieldValueUnsigned64{value: 0x0807060504030201}, CompEncoded: []byte{8, 7, 6, 5, 4, 3, 2, 1}},

		7:  {SourceVal: &FieldValueSigned8{value: -42}, CompEncoded: []byte{214}},
		8:  {SourceVal: &FieldValueSigned16{value: -0x0102}, CompEncoded: []byte{254, 254}},
		9:  {SourceVal: &FieldValueSigned16{value: -0x0201}, CompEncoded: []byte{253, 255}},
		10: {SourceVal: &FieldValueSigned32{value: -0x01020304}, CompEncoded: []byte{254, 253, 252, 252}},
		11: {SourceVal: &FieldValueSigned32{value: -0x04030201}, CompEncoded: []byte{251, 252, 253, 255}},
		12: {SourceVal: &FieldValueSigned64{value: -0x0102030405060708}, CompEncoded: []byte{254, 253, 252, 251, 250, 249, 248, 248}},
		13: {SourceVal: &FieldValueSigned64{value: -0x0807060504030201}, CompEncoded: []byte{247, 248, 249, 250, 251, 252, 253, 255}},

		14: {SourceVal: &FieldValueFloat32{value: 0x01020304}, CompEncoded: []byte{75, 129, 1, 130}},
		15: {SourceVal: &FieldValueFloat64{value: 0x0102030405060708}, CompEncoded: []byte{67, 112, 32, 48, 64, 80, 96, 112}},

		16: {SourceVal: &FieldValueBoolean{value: true}, CompEncoded: []byte{1}},
		17: {SourceVal: &FieldValueBoolean{value: false}, CompEncoded: []byte{2}},

		18: {SourceVal: &FieldValueMacAddress{value: net.HardwareAddr{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}}, CompEncoded: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}},

		19: {SourceVal: &FieldValueOctetArray{value: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}}, CompEncoded: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}},         //Not variable
		20: {SourceVal: &FieldValueOctetArray{value: largeOctetArray(100)}, CompEncoded: append([]byte{100}, largeOctetArray(100)...), VariableLength: true},       //VariableLength < 255
		21: {SourceVal: &FieldValueOctetArray{value: largeOctetArray(513)}, CompEncoded: append([]byte{255, 2, 1}, largeOctetArray(513)...), VariableLength: true}, //VariableLength >=255 (513)

		22: {SourceVal: &FieldValueString{value: "abcdef"}, CompEncoded: []byte{97, 98, 99, 100, 101, 102}},                                                            //Not variable
		23: {SourceVal: &FieldValueString{value: "一帆风顺"}, CompEncoded: []byte{12, 228, 184, 128, 229, 184, 134, 233, 163, 142, 233, 161, 186}, VariableLength: true},   //VariableLength < 255
		24: {SourceVal: &FieldValueString{value: string(largeOctetArray(513))}, CompEncoded: append([]byte{255, 2, 1}, largeOctetArray(513)...), VariableLength: true}, //VariableLength >= 255 (513)

		25: {SourceVal: &FieldValueDateTimeSeconds{value: dateB}, CompEncoded: []byte{82, 85, 27, 16}, VariableLength: false},

		26: {SourceVal: &FieldValueDateTimeMilliseconds{value: dateB}, CompEncoded: []byte{0, 0, 1, 65, 156, 113, 182, 128}, VariableLength: false},

		27: {SourceVal: &FieldValueDateTimeMicroseconds{value: dateB}, CompEncoded: []byte{213, 255, 153, 144, 0, 0, 0, 0}, VariableLength: false},

		28: {SourceVal: &FieldValueDateTimeNanoseconds{value: dateB}, CompEncoded: []byte{213, 255, 153, 144, 0, 0, 0, 0}, VariableLength: false},

		29: {SourceVal: &FieldValueIPv4Address{value: net.ParseIP("127.0.0.1")}, CompEncoded: []byte{127, 0, 0, 1}, VariableLength: false},
		30: {SourceVal: &FieldValueIPv4Address{value: net.ParseIP("1.2.3.4")}, CompEncoded: []byte{1, 2, 3, 4}, VariableLength: false},

		31: {SourceVal: &FieldValueIPv6Address{value: net.ParseIP("2001:db8::68")}, CompEncoded: []byte{32, 1, 13, 184, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 104}, VariableLength: false},
		32: {SourceVal: &FieldValueIPv6Address{value: net.ParseIP("1.2.3.4")}, CompEncoded: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 1, 2, 3, 4}, VariableLength: false},
	}

	for _, testcase := range testset {
		binarydata, err := testcase.SourceVal.MarshalBinary()
		if err != nil {
			t.Errorf("Error marshalling %#v: %#v", testcase.SourceVal, err)
		}
		lendata := []byte{}
		if testcase.VariableLength {
			lendata, err = EncodeVariableLength(binarydata)
			if err != nil {
				t.Errorf("Error encoding variable size %#v: %#v", testcase.SourceVal, err)
			}
		}
		casedata := append(lendata, binarydata...)
		if !bytes.Equal(casedata, testcase.CompEncoded) {
			t.Errorf("Error marshalling %#v, became %v but should have been %v", testcase.SourceVal, casedata, testcase.CompEncoded)
		}
		if testcase.VariableLength {
			caselen, err := DecodeVariableLength(casedata[0:3])
			if err != nil {
				t.Errorf("Error decoding length of variable element %#v", err)
			}
			if caselen != testcase.SourceVal.Len() {
				t.Errorf("Error in decoding length of variable element, wanted %d, but got %d", testcase.SourceVal.Len(), caselen)
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
	dateC.Add(12345 * time.Nanosecond)

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

		22: {SourceVal: &FieldValueMacAddress{value: net.HardwareAddr{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}}, DestVal: &FieldValueMacAddress{value: net.HardwareAddr{0, 0, 0, 0, 0, 0}}, CompVal: net.HardwareAddr{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}},

		23: {SourceVal: &FieldValueOctetArray{value: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}}, DestVal: &FieldValueOctetArray{value: []byte{}}, CompVal: []uint8{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}},
		24: {SourceVal: &FieldValueOctetArray{value: largeOctetArray(1024)}, DestVal: &FieldValueOctetArray{value: []byte{}}, CompVal: []uint8(largeOctetArray(1024))},

		25: {SourceVal: &FieldValueString{value: string([]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab})}, DestVal: &FieldValueString{value: ""}, CompVal: string([]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab})},
		26: {SourceVal: &FieldValueString{value: string(largeOctetArray(1024))}, DestVal: &FieldValueString{value: ""}, CompVal: string([]byte(largeOctetArray(1024)))},

		27: {SourceVal: &FieldValueDateTimeSeconds{value: dateA}, DestVal: &FieldValueDateTimeSeconds{value: time.Now()}, CompVal: dateA},
		28: {SourceVal: &FieldValueDateTimeSeconds{value: dateB}, DestVal: &FieldValueDateTimeSeconds{value: time.Now()}, CompVal: dateB},

		29: {SourceVal: &FieldValueDateTimeMilliseconds{value: dateA}, DestVal: &FieldValueDateTimeMilliseconds{value: time.Now()}, CompVal: dateA},
		30: {SourceVal: &FieldValueDateTimeMilliseconds{value: dateB}, DestVal: &FieldValueDateTimeMilliseconds{value: time.Now()}, CompVal: dateB},

		31: {SourceVal: &FieldValueDateTimeMicroseconds{value: dateA}, DestVal: &FieldValueDateTimeMicroseconds{value: time.Now()}, CompVal: dateA},
		32: {SourceVal: &FieldValueDateTimeMicroseconds{value: dateB}, DestVal: &FieldValueDateTimeMicroseconds{value: time.Now()}, CompVal: dateB},
		33: {SourceVal: &FieldValueDateTimeMicroseconds{value: dateC}, DestVal: &FieldValueDateTimeMicroseconds{value: time.Now()}, CompVal: dateC},

		34: {SourceVal: &FieldValueDateTimeNanoseconds{value: dateA}, DestVal: &FieldValueDateTimeNanoseconds{value: time.Now()}, CompVal: dateA},
		35: {SourceVal: &FieldValueDateTimeNanoseconds{value: dateB}, DestVal: &FieldValueDateTimeNanoseconds{value: time.Now()}, CompVal: dateB},
		36: {SourceVal: &FieldValueDateTimeNanoseconds{value: dateC}, DestVal: &FieldValueDateTimeNanoseconds{value: time.Now()}, CompVal: dateC},

		37: {SourceVal: &FieldValueIPv4Address{value: net.ParseIP("127.0.0.1").To4()}, DestVal: &FieldValueIPv4Address{value: net.ParseIP("10.0.0.1").To4()}, CompVal: net.ParseIP("127.0.0.1").To4()},
		38: {SourceVal: &FieldValueIPv6Address{value: net.ParseIP("2001:db8::68").To16()}, DestVal: &FieldValueIPv6Address{value: net.ParseIP("10.0.0.1").To16()}, CompVal: net.ParseIP("2001:db8::68").To16()},
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
		if !reflect.DeepEqual(testcase.SourceVal, testcase.DestVal) || !reflect.DeepEqual(testcase.DestVal.Value(), testcase.CompVal) {
			t.Errorf("Error in value after conversions, wanted %#v (%#v), but got %#v", testcase.SourceVal, testcase.CompVal, testcase.DestVal)
		}
		if testcase.SourceVal.Len() < 12 {
			fmt.Println(testcase.SourceVal.Value(), testcase.DestVal.Value(), testcase.CompVal)
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

func largeOctetArray(size int) []byte {
	retval := make([]byte, size)
	for idx := range retval {
		retval[idx] = byte(idx)
	}
	return retval
}
