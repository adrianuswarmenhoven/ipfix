package ipfixmessage

import (
	"bytes"
	"fmt"
	"net"
	"testing"
)

func TestDataRecordMarker(t *testing.T) {
	fmt.Printf(testMarkerString, "DataRecord")
}

func TestDataRecordBasic(t *testing.T) {
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

		19: {SourceVal: &FieldValueOctetArray{value: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}}, CompEncoded: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}},   //Not variable
		20: {SourceVal: &FieldValueOctetArray{value: largeOctetArray(100)}, CompEncoded: append([]byte{100}, largeOctetArray(100)...), VariableLength: true}, //VariableLength < 255
		/*21: {SourceVal: &FieldValueOctetArray{value: largeOctetArray(513)}, CompEncoded: append([]byte{255, 2, 1}, largeOctetArray(513)...), VariableLength: true}, //VariableLength >=255 (513)

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

		33: {SourceVal: &FieldValueBasicList{value: basicListA}, CompEncoded: []byte{3, 0, 12, 0, 4, 127, 0, 0, 1, 8, 8, 8, 8}, VariableLength: false},
		34: {SourceVal: &FieldValueBasicList{value: basicListB}, CompEncoded: []byte{3, 128, 11, 0, 4, 0, 0, 175, 113, 127, 0, 0, 1, 8, 8, 8, 8}, VariableLength: false},
		35: {SourceVal: &FieldValueBasicList{value: basicListC},
			CompEncoded: append([]byte{0, 128, 21, 255, 255, 0, 0, 175, 113, 12, 47, 102, 97, 118, 105, 99, 111, 110, 46, 105, 99, 111, 15, 47, 115, 116, 121, 108, 101, 115, 104, 101, 101, 116, 46, 99, 115, 115, 255, 1, 44}, largeOctetArray(300)...), VariableLength: false},
		*/
	}

	dr := &DataRecord{}
	compval := []byte{}
	fmt.Println("Must FIX Length encoding in fields")
	for _, testcase := range testset {
		dr.FieldValues = append(dr.FieldValues, testcase.SourceVal)
		compval = append(compval, testcase.CompEncoded...)
		if bindata, err := dr.MarshalBinary(); err == nil {
			if !bytes.Equal(compval, append(lendata, bindata...)) {
				t.Errorf("Error marshalling %#v: expected %#v, but got %#v", testcase.SourceVal, compval, bindata)
			}
		} else {
			t.Errorf("Error marshalling %#v: %#v", testcase.SourceVal, err)
		}

	}

	//    		dr.FieldValues = append(dr.FieldValues, &FieldValueBoolean{value: true})
	//		dr.FieldValues = append(dr.FieldValues, &FieldValueUnsigned64{value: uint64(0x0B16212C37424D58)})
}
