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
