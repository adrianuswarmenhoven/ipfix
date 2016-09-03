package ipfixmessage

import (
	"fmt"
	"testing"
	"time"
)

func TestEmptyMsg(t *testing.T) {
	testmsg, err := NewIPFIXMessage(-1)

	if (err == nil) || (testmsg != nil) {
		t.Fatalf("Should have gotten error and nil pointer but got %#v", testmsg)
	}

	testmsg, err = NewIPFIXMessage(0)

	if err != nil {
		t.Fatalf("New IPFIX Message creation failed: %#v", err)
	}

	if testmsg.MessageHeader.Length != 0 {
		t.Errorf("Expected message length of 0 but got %d", testmsg.MessageHeader.Length)
	}

	timestamp := time.Now()
	err = testmsg.Finalize(timestamp, 42)

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
