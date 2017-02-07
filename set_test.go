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
	if ipfixset_test_print {
		fmt.Println("--- EMPTY SET ---")
	}
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
	if ipfixset_test_print {
		fmt.Println("--- TEMPLATE RECORD SET ---")
	}
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
	umtestset.Pad(8) //Can not implicitly determine padding boundary

	if fmt.Sprintf("%s", umtestset) != fmt.Sprintf("%s", testset) {
		t.Fatalf("Expected: \n%s\nbut got:\n%s", testset, umtestset)
	}
}

func TestSetWithOptionsTemplateRecord(t *testing.T) {
	if ipfixset_test_print {
		fmt.Println("--- OPTIONS TEMPLATE SET ---")
	}
	testset, err := NewSet(SetIDOptionTemplate)
	if err != nil {
		t.Fatalf("New IPFIX Set creation failed: %#v", err)
	}

	newtemplrec, err := NewOptionsTemplateRecord(2525)
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

	newtemplrec2, err := NewOptionsTemplateRecord(1010)
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

	newtemplrec3, err := NewOptionsTemplateRecord(42042)
	if err != nil {
		t.Fatalf("New Options Template Record creation failed: %#v", err)
	}
	newfsp5, err := NewFieldSpecifier(44913, 101, 8)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec3.AddSpecifier(newfsp5)

	newfsp6, err := NewFieldSpecifier(44913, 102, 8)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec3.AddScopeSpecifier(newfsp6)

	testset.AddRecord(newtemplrec3)

	testset.Pad(8)
	data, err := testset.MarshalBinary()
	if err != nil {
		t.Fatalf("Marshalling set failed: %#v", err)
	}
	expectedmarshalresult := []byte{0, 3, 0, 72, 9, 221, 0, 3, 0, 0, 129, 2, 0, 8, 0, 0, 0, 123, 128, 101, 0, 4, 0, 0, 48, 57, 128, 102, 0, 4, 0, 0, 48, 57, 3, 242, 0, 1, 0, 0, 128, 100, 0, 8, 0, 0, 1, 200, 164, 58, 0, 2, 0, 1, 128, 102, 0, 8, 0, 0, 175, 113, 128, 101, 0, 8, 0, 0, 175, 113, 0, 0}
	if fmt.Sprintf("%+v", data) != fmt.Sprintf("%+v", expectedmarshalresult) {
		t.Fatalf("Marshalling failed. Expected \n%+v but got \n%+v", expectedmarshalresult, data)
	}
	if ipfixset_test_print {
		fmt.Println(testset)
		fmt.Println(testset.MarshalBinary())
	}

	umtestset := NewBlankSet()
	err = umtestset.UnmarshalBinary(expectedmarshalresult)
	umtestset.Pad(8) //Can not implicitly determine padding boundary

	if fmt.Sprintf("%s", umtestset) != fmt.Sprintf("%s", testset) {
		t.Fatalf("Expected: \n%s\nbut got:\n%s", testset, umtestset)
	}

}

func TestSetWithDataRecord(t *testing.T) {
	if ipfixset_test_print {
		fmt.Println("--- DATA RECORD SET ---")
	}
	testset, err := NewSet(257)
	if err != nil {
		t.Fatalf("New IPFIX Set creation failed: %#v", err)
	}

	tr, err := NewTemplateRecord(257)
	if err != nil {
		t.Errorf(errorPrefixMarker+"Error creating new template: %#v", err)
	}

	newfsp, err := NewFieldSpecifier(123, 258, 8)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	tr.AddSpecifier(newfsp)

	tmpat := NewActiveTemplateList()
	tmpat.Set(257, tr)

	testset.AssociateTemplates(tmpat)

}

func init() {
	if false {
		fmt.Println("ok")
	}
}
