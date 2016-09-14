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

func init() {
	if false {
		fmt.Println("ok")
	}
}
