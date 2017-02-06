package ipfixmessage

import (
	"fmt"
	"testing"
)

const (
	ipfixset_test_print = true
)

func TestSetMarker(t *testing.T) {
	if ipfixset_test_print {
		fmt.Printf(testMarkerString, "Set")
	}
}

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

	if ipfixset_test_print {
		fmt.Println(testset.MarshalBinary())
	}

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
	newfsp, err := NewFieldSpecifier(123, 258, 8)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec.AddSpecifier(newfsp)

	newfsp2, err := NewFieldSpecifier(12345, 101, 4)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec.AddSpecifier(newfsp2)

	newfsp3, err := NewFieldSpecifier(12345, 102, 4)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec.AddSpecifier(newfsp3)

	err = testset.AddRecord(newtemplrec)
	if err != nil {
		t.Fatalf("Template Record Addition to set failed: %#v", err)
	}

	newtemplrec2, err := NewTemplateRecord(1010)
	if err != nil {
		t.Fatalf("New Template Record creation failed: %#v", err)
	}
	newfsp4, err := NewFieldSpecifier(456, 100, 8)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec2.AddSpecifier(newfsp4)

	err = testset.AddRecord(newtemplrec2)
	if err != nil {
		t.Fatalf("Template Record Addition to set failed: %#v", err)
	}
	testset.Pad(8)
	data, err := testset.MarshalBinary()
	if err != nil {
		t.Fatalf("Marshalling set failed: %#v", err)
	}
	expectedmarshalresult := []byte{0, 2, 0, 48, 16, 146, 0, 3, 129, 2, 0, 8, 0, 0, 0, 123, 128, 101, 0, 4, 0, 0, 48, 57, 128, 102, 0, 4, 0, 0, 48, 57, 3, 242, 0, 1, 128, 100, 0, 8, 0, 0, 1, 200, 0, 0, 0, 0}
	if fmt.Sprintf("%+v", data) != fmt.Sprintf("%+v", expectedmarshalresult) {
		t.Fatalf("Marshalling failed. Expected \n%+v but got \n%+v", expectedmarshalresult, data)
	}
	if ipfixset_test_print {
		fmt.Println(testset)
		fmt.Println(testset.MarshalBinary())
	}

	umtestset := NewBlankSet()
	err = umtestset.UnmarshalBinary(expectedmarshalresult)
	fmt.Println(err, "aaaa")
	umtestset.Pad(8) //Can not implicitly determine padding boundary

	if fmt.Sprintf("%s", umtestset) != fmt.Sprintf("%s", testset) {
		t.Fatalf("Expected: \n%s\nbut got:\n%s", testset, umtestset)
	}
}

func init() {
	if false {
		fmt.Println("ok")
	}
}
