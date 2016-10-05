package ipfixmessage

import "fmt"

//--- Data Record

// DataRecord - A Data Record is a record that contains values of the parameters corresponding to a Template Record.
// The Data Records are sent in Data Sets which consist only of one or more Field Values.
// The Template ID to which the Field Values belong is encoded in the Set Header field "Set ID", i.e., "Set ID" == "Template ID".
// Note that Field Values do not necessarily have a length of 16 bits.
// Field Values are encoded according to their data type as specified in [RFC7012].
// Interpretation of the Data Record format can be done only if the Template Record corresponding to the Template ID is available at the Collecting Process.
type DataRecord struct {
	//AssociatedTemplate points to the template record. Without a template record a data record can not be encoded or decoded
	AssociatedTemplate *TemplateRecord

	FieldValues []FieldValue //Note that Field Values do not necessarily have a length of 16 bits. Field Values are encoded according to their data type specified in [RFC5102].
}

// Len returns the size in octets of the DataRecord
func (datrec *DataRecord) Len() uint16 {
	reclen := uint16(0)
	for _, listitem := range datrec.FieldValues {
		reclen += listitem.Len()
	}
	return reclen
}

// String returns the string representation of the Template Record
func (datrec *DataRecord) String() string {
	stringresult := ""
	for _, listitem := range datrec.FieldValues {
		stringresult += fmt.Sprintf("%#v", listitem)
	}
	return stringresult
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
// FieldValues have a type when added so there is implicit information on each field value to marshal it
func (datrec *DataRecord) MarshalBinary() (data []byte, err error) {
	if datrec.AssociatedTemplate == nil {
		return nil, fmt.Errorf("Can not marshal without a template")
	}
	if len(datrec.FieldValues) < 1 {
		return nil, fmt.Errorf("Can not marshal record, must have at least one Field Value")
	}
	marshalValue := []byte{}
	for _, listitem := range datrec.FieldValues {
		item, err := listitem.MarshalBinary()
		if err != nil {
			return nil, err
		}
		marshalValue = append(marshalValue, item...)
	}
	return marshalValue, nil
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
func (datrec *DataRecord) UnmarshalBinary(template *TemplateRecord, data []byte) error {
	if datrec.AssociatedTemplate == nil {
		return fmt.Errorf("Can not unmarshal without a template")
	}
	if data == nil || len(data) == 0 {
		return fmt.Errorf("Can not unmarshal, invalid data. %#v", data)
	}
	cursor := 0
	for _, recitem := range template.FieldSpecifiers {
		newval, err := NewFieldValueByID(int(recitem.EnterpriseNumber), int(recitem.InformationElementIdentifier))
		if err != nil {
			return err
		}
		if recitem.FieldLength != VariableLength {
			if cursor+int(recitem.FieldLength) > len(data) {
				return fmt.Errorf("Insufficient data to decode. Needed %d, but have %d", recitem.FieldLength, len(data[cursor:]))
			}
			err := newval.UnmarshalBinary(data[cursor : cursor+int(recitem.FieldLength)])
			if err != nil {
				return err
			}
			datrec.FieldValues = append(datrec.FieldValues, newval)
			cursor += int(recitem.FieldLength)
		} else {
			if cursor+3 > len(data) {
				return fmt.Errorf("Insufficient data to decode. Needed %d, but have %d", 3, len(data[cursor:]))
			}
			fieldlen, cursorshift, err := DecodeVariableLength(data[cursor : cursor+3])
			if err != nil {
				return err
			}
			if cursor+int(fieldlen)+int(cursorshift) > len(data) {
				return fmt.Errorf("Insufficient data to decode. Needed %d, but have %d", int(fieldlen)+int(cursorshift), len(data[cursor:]))
			}
			cursor += int(cursorshift)
			err = newval.UnmarshalBinary(data[cursor : cursor+int(fieldlen)])
			if err != nil {
				return err
			}
			datrec.FieldValues = append(datrec.FieldValues, newval)
			cursor += int(fieldlen)
		}
	}
	return nil
}
