package ipfixmessage

import (
	"fmt"
	"testing"
)

const (
	templaterecord_test_print = false
)

func TestTemplateRecordMarker(t *testing.T) {
	if templaterecord_test_print {
		fmt.Printf(testMarkerString, "Template Record")
	}
}

func TestTemplateRecordBasic(t *testing.T) {
	_, err := NewTemplateRecord(10)
	if err == nil {
		t.Errorf(errorPrefixMarker + "Error creating New Template Record. Should have gotten error, but got nil.")
	}

	tr, err := NewTemplateRecord(257)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new template: %#v", err)
	}
	if templaterecord_test_print {
		fmt.Println(tr)
	}
	fsp1, err := NewFieldSpecifier(0, 12, 4)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp1)
	if templaterecord_test_print {
		fmt.Println(tr)
	}
	fsp2, err := NewFieldSpecifier(44913, 20, VariableLength)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp2)
	if templaterecord_test_print {
		fmt.Println(tr)
	}
	binarydata, err := tr.MarshalBinary()
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error marshalling: %#v", err)
	}
	if templaterecord_test_print {
		fmt.Println(binarydata)
	}
	tr2, err := NewTemplateRecord(257)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new template: %#v", err)
	}
	err = tr2.UnmarshalBinary(binarydata)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error marshalling: %#v", err)
	}
	if templaterecord_test_print {
		fmt.Println(tr2)
	}
}

func TestOptionsTemplateRecordBasic(t *testing.T) {
	_, err := NewOptionsTemplateRecord(10)
	if err == nil {
		t.Errorf(errorPrefixMarker + "Error creating New Template Record. Should have gotten error, but got nil.")
	}

	tr, err := NewOptionsTemplateRecord(257)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new template: %#v", err)
	}
	if templaterecord_test_print {
		fmt.Println(tr)
	}
	fsp1, err := NewFieldSpecifier(0, 12, 4)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp1)
	if templaterecord_test_print {
		fmt.Println(tr)
	}
	fsp2, err := NewFieldSpecifier(44913, 20, VariableLength)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddSpecifier(fsp2)
	if templaterecord_test_print {
		fmt.Println(tr)
	}
	binarydata, err := tr.MarshalBinary()
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error marshalling: %#v", err)
	}

	if templaterecord_test_print {
		fmt.Println(tr)
	}
	ofsp1, err := NewFieldSpecifier(0, 12, 4)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddScopeSpecifier(ofsp1)
	if templaterecord_test_print {
		fmt.Println(tr)
	}
	ofsp2, err := NewFieldSpecifier(44913, 20, VariableLength)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new field specifier: %#v", err)
	}
	tr.AddScopeSpecifier(ofsp2)
	if templaterecord_test_print {
		fmt.Println(tr)
	}
	binarydata, err = tr.MarshalBinary()
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error marshalling: %#v", err)
	}

	if templaterecord_test_print {
		fmt.Println(binarydata)
	}
	tr2, err := NewOptionsTemplateRecord(257)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new template: %#v", err)
	}
	err = tr2.UnmarshalBinary(binarydata)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error marshalling: %#v", err)
	}
	if templaterecord_test_print {
		fmt.Println(tr2)
	}
}
