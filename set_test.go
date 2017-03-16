package ipfixmessage

import (
	"fmt"
	"testing"
)

const (
	ipfixsetTestPrint = false
)

func TestSetMarker(t *testing.T) {
	if ipfixsetTestPrint {
		fmt.Printf(testMarkerString, "Set")
	}
}

func TestEmptySet(t *testing.T) {
	if ipfixsetTestPrint {
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

	if ipfixsetTestPrint {
		fmt.Println(testset.MarshalBinary())
	}

}

func TestSetWithTemplateRecord(t *testing.T) {
	if ipfixsetTestPrint {
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
	if ipfixsetTestPrint {
		fmt.Println(testset)
	}

	umtestset := NewBlankSet()
	err = umtestset.UnmarshalBinary(expectedmarshalresult)
	if err != nil {
		t.Fatalf("Unmarshalling failed:%+v %+v", umtestset, err)
	}
	umtestset.Pad(8) //Can not implicitly determine padding boundary

	if fmt.Sprintf("%s", umtestset) != fmt.Sprintf("%s", testset) {
		t.Fatalf("Expected: \n%s\nbut got:\n%s", testset, umtestset)
	}
}

func TestSetWithOptionsTemplateRecord(t *testing.T) {
	if ipfixsetTestPrint {
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
	if ipfixsetTestPrint {
		fmt.Println(testset)
	}

	umtestset := NewBlankSet()
	err = umtestset.UnmarshalBinary(expectedmarshalresult)
	if err != nil {
		t.Fatalf("Unmarshalling failed:%+v", err)
	}
	umtestset.Pad(8) //Can not implicitly determine padding boundary

	if fmt.Sprintf("%s", umtestset) != fmt.Sprintf("%s", testset) {
		t.Fatalf("Expected: \n%s\nbut got:\n%s", testset, umtestset)
	}

}

func TestSetWithDataRecord(t *testing.T) {
	if ipfixsetTestPrint {
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

	newfsp, err := NewFieldSpecifier(0, 84, VariableLength) //Add a string
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	tr.AddSpecifier(newfsp)

	newfsp2, err := NewFieldSpecifier(0, 85, 8) //Add a single 64-bit field
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	tr.AddSpecifier(newfsp2)

	tmpat := NewActiveTemplateList()
	tmpat.Set(257, tr)

	testset.AssociateTemplates(tmpat)

	dr, err := NewDataRecord(257, testset.AssociatedTemplates)
	if err != nil {
		t.Fatalf("New Data record creation failed: %#v", err)
	}
	err = dr.AddFieldValue(&FieldValueString{value: "test string samp"})
	if err != nil {
		t.Fatalf("Addition of string failed: %#v", err)
	}
	err = dr.AddFieldValue(&FieldValueString{value: "test string"})
	if err == nil {
		t.Fatalf("Addition of string should have failed but error is %#v", err)
	}
	err = dr.AddFieldValue(&FieldValueUnsigned64{value: 0x0807060504030201})
	if err != nil {
		t.Fatalf("Addition of 64-bit unsigned integer failed: %#v", err)
	}
	err = dr.AddFieldValue(&FieldValueUnsigned64{value: 0x0807060504030201})
	if err == nil {
		t.Fatalf("Addition of 64-bit unsigned integer should have failed but error is %#v", err)
	}

	err = testset.AddRecord(dr)
	if err != nil {
		t.Fatalf("Additon of data record to test set failed: %#v", err)
	}

	testset.Pad(7)
	if ipfixsetTestPrint {
		fmt.Println(testset, testset.Len())
	}

	data, err := testset.MarshalBinary()
	if err != nil {
		t.Fatalf("Marshalling set failed: %#v", err)
	}
	if ipfixsetTestPrint {
		fmt.Println(data)
	}

	umtestset := NewBlankSet()
	umtestset.AssociateTemplates(tmpat) //Must associate templates
	err = umtestset.UnmarshalBinary(data)
	if err != nil {
		t.Fatalf("Unmarshalling set failed: %#v", err)
	}
	umtestset.Pad(7) //Can not implicitly determine padding boundary

	if ipfixsetTestPrint {
		fmt.Println(umtestset, err)
	}

	if fmt.Sprintf("%s", umtestset) != fmt.Sprintf("%s", testset) {
		t.Fatalf("Expected: \n%s\nbut got:\n%s", testset, umtestset)
	}
}

func init() {
	if false {
		fmt.Println("ok")
	}
}
