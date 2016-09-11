package ipfixmessage

import (
	"fmt"
	"testing"
	"time"
)

func TestMessageHeader(t *testing.T) {
	testmsg, err := NewMessage()

	if err != nil {
		t.Fatalf("New IPFIX Message creation failed: %#v", err)
	}

	// We have a defeault of ODID 0 anyway
	//	testmsg = testmsg.ODID(0)

	if testmsg.MessageHeader.Length != 0 {
		t.Errorf("Expected message length of 0 but got %d", testmsg.MessageHeader.Length)
	}

	timestamp := time.Now()
	testmsg, err = testmsg.ExportTime(timestamp).SequenceNumber(42).Finalize()

	if err != nil {
		t.Fatalf("Finalize failed: %#v", err)
	}

	if testmsg.MessageHeader.Length != ipfixMessageHeaderLength {
		t.Errorf("Expected message length of %d but got %d", ipfixMessageHeaderLength, testmsg.MessageHeader.Length)
	}

	if testmsg.MessageHeader.ExportTime != timestamp {
		t.Errorf("Expected timestamp of %v but got %v", timestamp, testmsg.MessageHeader.ExportTime)
	}

	if testmsg.MessageHeader.SequenceNumber != 42 {
		t.Errorf("Expected sequence number of 42 but got %d", testmsg.MessageHeader.SequenceNumber)
	}
}

func init() {
	if false {
		fmt.Println("ok")
	}
}
