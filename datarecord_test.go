package ipfixmessage

import (
	"fmt"
	"testing"
)

func TestDataRecordMarker(t *testing.T) {
	fmt.Printf(testMarkerString, "DataRecord")
}

func TestDataRecordBasic(t *testing.T) {

	dr := &DataRecord{}

	dr.FieldValues = append(dr.FieldValues, &FieldValueBoolean{value: true})
	fmt.Println("DataRecord")
}
