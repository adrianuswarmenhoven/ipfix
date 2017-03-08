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

	//if fmt.Sprintf("%+v", data) != "[0 10 0 16 82 221 166 236 0 0 0 1 0 0 0 2]" {
	//	t.Fatalf("Error marshalling. Expected [0 10 0 16 82 221 166 236 0 0 0 1 0 0 0 2] but got %+v", data)
	//}

	receivermessage, err := NewMessage()

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
