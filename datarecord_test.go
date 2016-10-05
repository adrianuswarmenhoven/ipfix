package ipfixmessage

import (
	"bytes"
	"fmt"
	"net"
	"testing"
)

func TestDataRecordMarker(t *testing.T) {
	fmt.Printf(testMarkerString, "Data Record")
}

type dataRecordTestcase struct {
	SourceVal   FieldValue
	CompEncoded []byte

	TemplateEnterpriseID         uint32
	TemplateInformationElementID uint16
	TemplateFieldLength          uint16
}

func TestDataRecordBasic(t *testing.T) {
	var testset = []dataRecordTestcase{
		0: {SourceVal: &FieldValueUnsigned8{value: 42}, CompEncoded: []byte{42}, TemplateEnterpriseID: 0, TemplateInformationElementID: 9, TemplateFieldLength: 1},
		1: {SourceVal: &FieldValueUnsigned16{value: 0x0102}, CompEncoded: []byte{1, 2}, TemplateEnterpriseID: 0, TemplateInformationElementID: 11, TemplateFieldLength: 2},
		2: {SourceVal: &FieldValueUnsigned32{value: 0x01020304}, CompEncoded: []byte{1, 2, 3, 4}, TemplateEnterpriseID: 16982, TemplateInformationElementID: 200, TemplateFieldLength: 4},
		3: {SourceVal: &FieldValueUnsigned64{value: 0x0807060504030201}, CompEncoded: []byte{8, 7, 6, 5, 4, 3, 2, 1}, TemplateEnterpriseID: 39499, TemplateInformationElementID: 48, TemplateFieldLength: 8},
		4: {SourceVal: &FieldValueBoolean{value: true}, CompEncoded: []byte{1}, TemplateEnterpriseID: 0, TemplateInformationElementID: 276, TemplateFieldLength: 1},
		5: {SourceVal: &FieldValueMacAddress{value: net.HardwareAddr{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}}, CompEncoded: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab}, TemplateEnterpriseID: 0, TemplateInformationElementID: 365, TemplateFieldLength: 6},
		6: {SourceVal: &FieldValueString{value: "abcdef"}, CompEncoded: []byte{97, 98, 99, 100, 101, 102}, TemplateEnterpriseID: 0, TemplateInformationElementID: 371, TemplateFieldLength: 6}, //Not variable
		7: {SourceVal: &FieldValueIPv4Address{value: net.ParseIP("127.0.0.1")}, CompEncoded: []byte{127, 0, 0, 1}, TemplateEnterpriseID: 0, TemplateInformationElementID: 8, TemplateFieldLength: 4},
		8: {SourceVal: &FieldValueString{value: "abcdef"}, CompEncoded: []byte{6, 97, 98, 99, 100, 101, 102}, TemplateEnterpriseID: 0, TemplateInformationElementID: 371, TemplateFieldLength: VariableLength},                                   //Variable
		9: {SourceVal: &FieldValueString{value: string(largeOctetArray(513))}, CompEncoded: append([]byte{255, 2, 1}, largeOctetArray(513)...), TemplateEnterpriseID: 0, TemplateInformationElementID: 371, TemplateFieldLength: VariableLength}, //Variable
	}
	tr, err := NewTemplateRecord(257)
	if err != nil {
		t.Errorf("Error creating new template: %#v", err)
	}
	dr := &DataRecord{}
	_, err = dr.MarshalBinary()
	if err == nil {
		t.Errorf("Should have gotten error for trying to marshal without template")
	}
	err = dr.AssociateTemplate(tr)
	if err != nil {
		t.Errorf("Got error associating template: %#v", tr)
	}
	compval := []byte{}
	for _, testcase := range testset {
		newfsp, err := NewFieldSpecifier(testcase.TemplateEnterpriseID, testcase.TemplateInformationElementID, testcase.TemplateFieldLength)
		if err != nil {
			t.Errorf("Error creating new field specifier: %#v", err)
		}
		tr.AddSpecifier(newfsp)
		dr.FieldValues = append(dr.FieldValues, testcase.SourceVal)
		bindata, err := dr.MarshalBinary()
		if err != nil {
			t.Errorf("Error marshalling datarecord: %#v", err)
		}
		compval = append(compval, testcase.CompEncoded...)
		if !bytes.Equal(compval, bindata) {
			t.Errorf("Error marshalling %#v: expected %#v, but got %#v", testcase.SourceVal, compval, bindata)
		}
	}
	fmt.Println(dr)
	dr2 := &DataRecord{}
	dr2.AssociateTemplate(tr)
	dr2.UnmarshalBinary(compval)
	fmt.Println(dr2)
	for idx, item := range dr.FieldValues {
		if fmt.Sprintf("%#v", item.Value()) != fmt.Sprintf("%#v", dr2.FieldValues[idx].Value()) {
			t.Errorf("Error unmarshalling %#v: expected %#v but got %#v", item, item.Value(), dr2.FieldValues[idx].Value())
		}
		fmt.Println("------\n", fmt.Sprintf("%#v", item.Value()), "\n", fmt.Sprintf("%#v", dr2.FieldValues[idx].Value()), "\n------\n")
	}
}
