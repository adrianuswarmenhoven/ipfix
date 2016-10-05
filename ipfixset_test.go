package ipfixmessage

import (
	"fmt"
	"testing"
)

func TestEmptySet(t *testing.T) {
	testset, err := NewSet(0)

	if (err == nil) || (testset != nil) {
		t.Fatalf("Should have gotten error and nil pointer but got %#v", testset)
	}

	testset, err = NewSet(42)

	if (err == nil) || (testset != nil) {
		t.Fatalf("Should have gotten error and nil pointer but got %#v", testset)
	}

	testset, err = NewSet(SetIDTemplate)

	if err != nil {
		t.Fatalf("New IPFIX Set creation failed: %#v", err)
	}

	/*err = testset.Finalize(0)
	if err != nil {
		t.Fatalf("Finalize failed: %#v", err)
	}

	if testset.SetHeader.Length != ipfixSetHeaderLength {
		t.Errorf("Expected message length of %d but got %d", ipfixSetHeaderLength, testset.SetHeader.Length)
	}*/
}

func TestSetWithTemplateRecord(t *testing.T) {
	testset, err := NewSet(SetIDTemplate)
	if err != nil {
		t.Fatalf("New IPFIX Set creation failed: %#v", err)
	}
	newtemplrec, err := NewTemplateRecord(4242)
	if err != nil {
		t.Fatalf("New Template Record creation failed: %#v", err)
	}
	newfsp, err := NewFieldSpecifier(12345, 100, 8)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec.AddSpecifier(newfsp)

	newfsp2, err := NewFieldSpecifier(12345, 101, 4)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec.AddSpecifier(newfsp2)

	err = testset.AddRecord(newtemplrec)
	if err != nil {
		t.Fatalf("Template Record Addition to set failed: %#v", err)
	}

	testset.Pad(4)

	fmt.Println(testset)
}

func init() {
	if false {
		fmt.Println("ok")
	}
}
