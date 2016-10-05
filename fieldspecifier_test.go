package ipfixmessage

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFieldSpecifierMarker(t *testing.T) {
	fmt.Printf(testMarkerString, "Field Specifier")
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
	}

	for _, testcase := range testset {
		fl, err := FieldLengthByID(int(testcase.EnterpriseID), int(testcase.InformationFieldElement))
		if (err != nil) == testcase.MustNotFailCreate {
			t.Errorf("Error determining field length of testcase %#v: %#v", testcase, err)
		}
		fsp, err := NewFieldSpecifier(testcase.EnterpriseID, testcase.InformationFieldElement, fl)
		if (err != nil) == testcase.MustNotFailCreate {
			t.Errorf("Error creating new field specifier of testcase %#v: %#v", testcase, err)
		}
		binarydata, err := fsp.MarshalBinary()
		if err != nil {
			t.Errorf("Error marshalling new field specifier of testcase %#v: %#v", testcase, err)
		}
		if bytes.Equal(binarydata, testcase.CompEncoded) != testcase.MustNotFailEncode {
			t.Errorf("Error marshalling %#v, became %v but should have been %v (%v)", testcase, binarydata, testcase.CompEncoded, testcase.MustNotFailEncode)
		}
		if len(binarydata) > 0 && testcase.MustNotFailEncode {
			fsp2, _ := NewFieldSpecifier(0, 0, 0)
			fsp2.UnmarshalBinary(binarydata)
			if fmt.Sprintf("%#v", fsp) != fmt.Sprintf("%#v", fsp2) {
				t.Errorf("Error unmarshalling %#v, became %#v", fsp, fsp2)
			}
		}

	}
}
