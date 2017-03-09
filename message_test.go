package ipfixmessage

import (
	"fmt"
	"testing"
	"time"
)

func TestMessageMarker(t *testing.T) {
	if ipfixset_test_print {
		fmt.Printf(testMarkerString, "Message")
	}
}

func TestEmptyMessage(t *testing.T) {
	if ipfixset_test_print {
		fmt.Println("--- EMPTY MESSAGE ---")
	}
	testmessage, err := NewMessage()

	if err != nil {
		t.Fatalf("Error creating new message %#v", err)
	}

	if (err == nil) && (testmessage == nil) {
		t.Fatalf("Should have gotten error or non-nil pointer but got %#v", testmessage)
	}

	testmessage.SequenceNumber = 1
	testmessage.ObservationDomainID = 2
	testmessage.ExportTime, err = time.Parse(time.RFC822, "20 Jan 14 23:45 CET")
	if err != nil {
		t.Fatal(err)
	}

	if ipfixset_test_print {
		fmt.Println(testmessage.MarshalBinary())
	}

	data, err := testmessage.MarshalBinary()

	if err != nil {
		t.Fatalf("Error marshalling message %#v", err)
	}

	if fmt.Sprintf("%+v", data) != "[0 10 0 16 82 221 166 236 0 0 0 1 0 0 0 2]" {
		t.Fatalf("Error marshalling. Expected [0 10 0 16 82 221 166 236 0 0 0 1 0 0 0 2] but got %+v", data)
	}

	receivermessage, err := NewMessage()
	receivermessage.AssociatedTemplates = NewActiveTemplateList() //Active templates belong to a session

	if err != nil {
		t.Fatalf("Error creating new message %#v", err)
	}

	err = receivermessage.UnmarshalBinary(data)

	if err != nil {
		t.Fatalf("Error unmarshalling message %#v", err)
	}

	if fmt.Sprintf("%+v", receivermessage) != fmt.Sprintf("%+v", testmessage) {
		t.Fatalf("Error unmarshalling message. Expected '%+v' but got '%+v'", fmt.Sprintf("%+v", testmessage), fmt.Sprintf("%+v", receivermessage))
	}
}

func TestMessageSingleSet(t *testing.T) {
	if ipfixset_test_print {
		fmt.Println("--- MESSAGE SINGLE SET ---")
	}
	testmessage, err := NewMessage()

	if err != nil {
		t.Fatalf("Error creating new message %#v", err)
	}

	if (err == nil) && (testmessage == nil) {
		t.Fatalf("Should have gotten error or non-nil pointer but got %#v", testmessage)
	}

	testmessage.SequenceNumber = 1
	testmessage.ObservationDomainID = 2
	testmessage.ExportTime, err = time.Parse(time.RFC822, "20 Jan 14 23:45 CET")
	if err != nil {
		t.Fatal(err)
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
	//testset.Pad(8)

	testmessage.Sets = append(testmessage.Sets, testset)

	if ipfixset_test_print {
		fmt.Println(testmessage.MarshalBinary())
	}

	data, err := testmessage.MarshalBinary()

	if err != nil {
		t.Fatalf("Error marshalling message %#v", err)
	}

	receivermessage, err := NewMessage()
	receivermessage.AssociatedTemplates = NewActiveTemplateList()

	if err != nil {
		t.Fatalf("Error creating new message %#v", err)
	}

	err = receivermessage.UnmarshalBinary(data)

	if err != nil {
		t.Fatalf("Error unmarshalling message %#v", err)
	}

	if fmt.Sprintf("%+v", receivermessage) != fmt.Sprintf("%+v", testmessage) {
		t.Fatalf("Error unmarshalling message. Expected '%+v' but got '%+v'", fmt.Sprintf("%+v", testmessage), fmt.Sprintf("%+v", receivermessage))
	}

	if ipfixset_test_print {
		fmt.Println(receivermessage.AssociatedTemplates.Get(1010))
	}
}

func TestMessageDataSet(t *testing.T) {
	if ipfixset_test_print {
		fmt.Println("--- MESSAGE DATA SET ---")
	}
	//First test that we should get an error marshalling when we do not have the correct template
	testmessagenotmpl, err := NewMessage()

	if err != nil {
		t.Fatalf("Error creating new message %#v", err)
	}

	if (err == nil) && (testmessagenotmpl == nil) {
		t.Fatalf("Should have gotten error or non-nil pointer but got %#v", testmessagenotmpl)
	}

	testmessagenotmpl.SequenceNumber = 1
	testmessagenotmpl.ObservationDomainID = 2
	testmessagenotmpl.ExportTime, err = time.Parse(time.RFC822, "20 Jan 14 23:45 CET")
	if err != nil {
		t.Fatal(err)
	}

	drnotmpl, err := NewDataRecord(257, nil)
	if err != nil {
		t.Fatalf("New Data record creation failed: %#v", err)
	}

	err = drnotmpl.AddFieldValue(&FieldValueString{value: "test string samp"})
	if err != nil {
		t.Fatalf("Addition of string failed: %#v", err)
	}
	err = drnotmpl.AddFieldValue(&FieldValueUnsigned64{value: 0x0807060504030201})
	if err != nil {
		t.Fatalf("Addition of 64-bit unsigned integer failed: %#v", err)
	}
	testsetnotmpl, err := NewSet(257)
	if err != nil {
		t.Fatalf("New IPFIX Set creation failed: %#v", err)
	}

	testsetnotmpl.AddRecord(drnotmpl)

	testmessagenotmpl.Sets = append(testmessagenotmpl.Sets, testsetnotmpl)

	_, err = testmessagenotmpl.MarshalBinary()
	if err == nil {
		t.Fatalf("Should have gotten error marshalling datarecords without proper template")
	}
	if ipfixset_test_print {
		fmt.Println(testmessagenotmpl.MarshalBinary())
	}

	testmessage, err := NewMessage()

	if err != nil {
		t.Fatalf("Error creating new message %#v", err)
	}

	if (err == nil) && (testmessage == nil) {
		t.Fatalf("Should have gotten error or non-nil pointer but got %#v", testmessage)
	}

	testmessage.SequenceNumber = 2
	testmessage.ObservationDomainID = 2
	testmessage.ExportTime, err = time.Parse(time.RFC822, "20 Jan 14 23:45 CET")
	if err != nil {
		t.Fatal(err)
	}
	testmessage.AssociatedTemplates = NewActiveTemplateList()

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

	testmessage.AssociatedTemplates.Set(257, tr)

	dr, err := NewDataRecord(257, testmessage.AssociatedTemplates)
	if err != nil {
		t.Fatalf("New Data record creation failed: %#v", err)
	}

	err = dr.AddFieldValue(&FieldValueString{value: "test string samp"})
	if err != nil {
		t.Fatalf("Addition of string failed: %#v", err)
	}
	err = dr.AddFieldValue(&FieldValueUnsigned64{value: 0x0807060504030201})
	if err != nil {
		t.Fatalf("Addition of 64-bit unsigned integer failed: %#v", err)
	}
	testset, err := NewSet(257)
	if err != nil {
		t.Fatalf("New IPFIX Set creation failed: %#v", err)
	}

	testset.AddRecord(dr)

	testmessage.Sets = append(testmessage.Sets, testset)

	//	_, err = testmessagenotmpl.MarshalBinary()

	if ipfixset_test_print {
		fmt.Println(testmessage.MarshalBinary())
	}

	data, err := testmessage.MarshalBinary()
	if err != nil {
		t.Fatalf("Error marshalling record: %+v", err)
	}

	receivermessage, err := NewMessage()
	receivermessage.AssociatedTemplates = testmessage.AssociatedTemplates

	if err != nil {
		t.Fatalf("Error creating new message %#v", err)
	}

	err = receivermessage.UnmarshalBinary(data)

	if err != nil {
		t.Fatalf("Error unmarshalling message %#v", err)
	}

	if fmt.Sprintf("%+v", receivermessage) != fmt.Sprintf("%+v", testmessage) {
		t.Fatalf("Error unmarshalling message. Expected '%+v' but got '%+v'", fmt.Sprintf("%+v", testmessage), fmt.Sprintf("%+v", receivermessage))
	}

	if ipfixset_test_print {
		fmt.Println(receivermessage)
	}

}

func TestMessageInterleavedSet(t *testing.T) {
	if ipfixset_test_print {
		fmt.Println("--- MESSAGE INTERLEAVED SET ---")
	}
	testmessage, err := NewMessage()

	if err != nil {
		t.Fatalf("Error creating new message %#v", err)
	}

	if (err == nil) && (testmessage == nil) {
		t.Fatalf("Should have gotten error or non-nil pointer but got %#v", testmessage)
	}

	testmessage.SequenceNumber = 3
	testmessage.ObservationDomainID = 2
	testmessage.ExportTime, err = time.Parse(time.RFC822, "20 Jan 14 23:45 CET")
	if err != nil {
		t.Fatal(err)
	}
	testmessage.AssociatedTemplates = NewActiveTemplateList()

	//Template set
	//Create template for 257
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

	testtplset, err := NewSet(SetIDTemplate)
	if err != nil {
		t.Fatalf("New IPFIX Set creation failed: %#v", err)
	}
	err = testtplset.AddRecord(tr)
	if err != nil {
		t.Fatalf("Template Record Addition to set failed: %#v", err)
	}
	testmessage.AssociatedTemplates.Set(257, tr)

	newtemplrec, err := NewTemplateRecord(4242)
	if err != nil {
		t.Fatalf("New Template Record creation failed: %#v", err)
	}
	newfsp, err = NewFieldSpecifier(123, 258, 8)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec.AddSpecifier(newfsp)

	newfsp2, err = NewFieldSpecifier(12345, 101, 4)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec.AddSpecifier(newfsp2)

	newfsp3, err := NewFieldSpecifier(12345, 102, 4)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec.AddSpecifier(newfsp3)

	err = testtplset.AddRecord(newtemplrec)
	if err != nil {
		t.Fatalf("Template Record Addition to set failed: %#v", err)
	}
	testmessage.AssociatedTemplates.Set(4242, newtemplrec)

	newtemplrec2, err := NewTemplateRecord(1010)
	if err != nil {
		t.Fatalf("New Template Record creation failed: %#v", err)
	}
	newfsp4, err := NewFieldSpecifier(456, 100, 8)
	if err != nil {
		t.Fatalf("New Field Specifier creation failed: %#v", err)
	}
	newtemplrec2.AddSpecifier(newfsp4)

	err = testtplset.AddRecord(newtemplrec2)
	if err != nil {
		t.Fatalf("Template Record Addition to set failed: %#v", err)
	}
	testmessage.AssociatedTemplates.Set(1010, newtemplrec2)
	//End of templates
	testmessage.Sets = append(testmessage.Sets, testtplset)

	dr, err := NewDataRecord(257, testmessage.AssociatedTemplates)
	if err != nil {
		t.Fatalf("New Data record creation failed: %#v", err)
	}

	err = dr.AddFieldValue(&FieldValueString{value: "test string samp"})
	if err != nil {
		t.Fatalf("Addition of string failed: %#v", err)
	}
	err = dr.AddFieldValue(&FieldValueUnsigned64{value: 0x0807060504030201})
	if err != nil {
		t.Fatalf("Addition of 64-bit unsigned integer failed: %#v", err)
	}

	testset, err := NewSet(257)
	if err != nil {
		t.Fatalf("New IPFIX Set creation failed: %#v", err)
	}

	err = testset.AddRecord(dr)
	if err != nil {
		t.Fatalf("Addition data record failed: %#v", err)
	}

	testmessage.Sets = append(testmessage.Sets, testset)

	//	_, err = testmessagenotmpl.MarshalBinary()

	if ipfixset_test_print {
		fmt.Println(testmessage.MarshalBinary())
	}

	data, err := testmessage.MarshalBinary()
	if err != nil {
		t.Fatalf("Error marshalling record: %+v", err)
	}

	receivermessage, err := NewMessage()
	receivermessage.AssociatedTemplates = testmessage.AssociatedTemplates

	if err != nil {
		t.Fatalf("Error creating new message %#v", err)
	}

	err = receivermessage.UnmarshalBinary(data)

	if err != nil {
		t.Fatalf("Error unmarshalling message %#v", err)
	}

	if fmt.Sprintf("%+v", receivermessage) != fmt.Sprintf("%+v", testmessage) {
		t.Fatalf("Error unmarshalling message. Expected '%+v' but got '%+v'", fmt.Sprintf("%+v", testmessage), fmt.Sprintf("%+v", receivermessage))
	}

	if ipfixset_test_print {
		fmt.Println(receivermessage)
	}

}
