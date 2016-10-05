package ipfixmessage

import (
	"fmt"
	"testing"
)

func TestTemplateRecordMarker(t *testing.T) {
	fmt.Printf(testMarkerString, "Template Record")
}

func TestTemplateRecordBasic(t *testing.T) {
	_, err := NewTemplateRecord(10)
	if err == nil {
		t.Errorf("Error creating New Template Record. Should have gotten error, but got nil.")
	}

	tr, err := NewTemplateRecord(257)
	if err != nil {
		t.Errorf("Error creating new template: %#v", err)
	}
	fmt.Println(tr)
	fsp1, err := NewFieldSpecifier(0, 12, 4)
	if err != nil {
		t.Errorf("Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp1)
	fmt.Println(tr)
	fsp2, err := NewFieldSpecifier(44913, 20, VariableLength)
	if err != nil {
		t.Errorf("Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp2)
	fmt.Println(tr)
	binarydata, err := tr.MarshalBinary()
	if err != nil {
		t.Errorf("Error marshalling: %#v", err)
	}

	fmt.Println(binarydata)
	tr2, err := NewTemplateRecord(257)
	if err != nil {
		t.Errorf("Error creating new template: %#v", err)
	}
	err = tr2.UnmarshalBinary(binarydata)
	if err != nil {
		t.Errorf("Error marshalling: %#v", err)
	}
	fmt.Println(tr2)
}

func TestOptionsTemplateRecordBasic(t *testing.T) {
	_, err := NewOptionsTemplateRecord(10)
	if err == nil {
		t.Errorf("Error creating New Template Record. Should have gotten error, but got nil.")
	}

	tr, err := NewOptionsTemplateRecord(257)
	if err != nil {
		t.Errorf("Error creating new template: %#v", err)
	}
	fmt.Println(tr)
	fsp1, err := NewFieldSpecifier(0, 12, 4)
	if err != nil {
		t.Errorf("Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp1)
	fmt.Println(tr)
	fsp2, err := NewFieldSpecifier(44913, 20, VariableLength)
	if err != nil {
		t.Errorf("Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp2)
	fmt.Println(tr)
	binarydata, err := tr.MarshalBinary()
	if err != nil {
		t.Errorf("Error marshalling: %#v", err)
	}

	fmt.Println(tr)
	ofsp1, err := NewFieldSpecifier(0, 12, 4)
	if err != nil {
		t.Errorf("Error creating new field specifier: %#v", err)
	}
	tr.AddScopeSpecifier(ofsp1)
	fmt.Println(tr)
	ofsp2, err := NewFieldSpecifier(44913, 20, VariableLength)
	if err != nil {
		t.Errorf("Error creating new field specifier: %#v", err)
	}
	tr.AddScopeSpecifier(ofsp2)
	fmt.Println(tr)
	binarydata, err = tr.MarshalBinary()
	if err != nil {
		t.Errorf("Error marshalling: %#v", err)
	}

	fmt.Println(binarydata)
	tr2, err := NewOptionsTemplateRecord(257)
	if err != nil {
		t.Errorf("Error creating new template: %#v", err)
	}
	err = tr2.UnmarshalBinary(binarydata)
	if err != nil {
		t.Errorf("Error marshalling: %#v", err)
	}
	fmt.Println(tr2)
}
