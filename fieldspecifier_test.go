package ipfix

import (
	"bytes"
	"fmt"
	"testing"
)

const (
	fieldspecifierTestPrint = false
)

func TestFieldSpecifierMarker(t *testing.T) {
	if fieldspecifierTestPrint {
		fmt.Printf(testMarkerString, "Field Specifier")
	}
}

type fieldSpecifierTestCase struct {
	EnterpriseID            uint32
	InformationFieldElement uint16
	CompEncoded             []byte
	MustNotFailCreate       bool
	MustNotFailEncode       bool
}

func TestFieldSpecifier(t *testing.T) {
	testset := []fieldSpecifierTestCase{
		0: {EnterpriseID: 0, InformationFieldElement: 12, CompEncoded: []byte{0, 12, 0, 4}, MustNotFailCreate: true, MustNotFailEncode: true},
		1: {EnterpriseID: 0, InformationFieldElement: 12, CompEncoded: []byte{0, 12, 0, 1}, MustNotFailCreate: true, MustNotFailEncode: false},
		2: {EnterpriseID: 65535, InformationFieldElement: 65535, CompEncoded: []byte{}, MustNotFailCreate: false, MustNotFailEncode: false},
		3: {EnterpriseID: 44913, InformationFieldElement: 20, CompEncoded: []byte{128, 20, 255, 255, 0, 0, 175, 113}, MustNotFailCreate: true, MustNotFailEncode: true},
		4: {EnterpriseID: 65534, InformationFieldElement: 1, CompEncoded: []byte{128, 1, 0, 100, 0, 0, 255, 254}, MustNotFailCreate: true, MustNotFailEncode: true},
		5: {EnterpriseID: 65534, InformationFieldElement: 10, CompEncoded: []byte{128, 10, 0, 4, 0, 0, 255, 254}, MustNotFailCreate: true, MustNotFailEncode: true},
	}

	RegisterCustomField(65534, 1, 100, "test", &FieldValueOctetArray{})
	RegisterCustomField(65534, 10, 4, "ipAddress", &FieldValueIPv4Address{})

	for _, testcase := range testset {
		fl, err := FieldLengthByID(testcase.EnterpriseID, testcase.InformationFieldElement)
		if (err != nil) == testcase.MustNotFailCreate {
			t.Errorf(errorPrefixMarker+"Error determining field length of testcase %#v: %#v", testcase, err)
		}
		fsp, err := NewFieldSpecifier(testcase.EnterpriseID, testcase.InformationFieldElement, fl)
		if (err != nil) == testcase.MustNotFailCreate {
			t.Errorf(errorPrefixMarker+"Error creating new field specifier of testcase %#v: %#v", testcase, err)
		}
		binarydata, err := fsp.MarshalBinary()
		if err != nil {
			t.Errorf(errorPrefixMarker+"Error marshalling new field specifier of testcase %#v: %#v", testcase, err)
		}
		if bytes.Equal(binarydata, testcase.CompEncoded) != testcase.MustNotFailEncode {
			t.Errorf(errorPrefixMarker+"Error marshalling %#v, became %v but should have been %v (%v)", testcase, binarydata, testcase.CompEncoded, testcase.MustNotFailEncode)
		}
		if len(binarydata) > 0 && testcase.MustNotFailEncode {
			fsp2, _ := NewFieldSpecifier(0, 0, 0)
			fsp2.UnmarshalBinary(binarydata)
			if fmt.Sprintf("%#v", fsp) != fmt.Sprintf("%#v", fsp2) {
				t.Errorf(errorPrefixMarker+"Error unmarshalling %#v, became %#v", fsp, fsp2)
			}
		}

	}
}
