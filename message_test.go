package ipfixmessage

import (
	"fmt"
	"testing"
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

	if ipfixset_test_print {
		fmt.Println(testmessage.MarshalBinary())
	}

}
