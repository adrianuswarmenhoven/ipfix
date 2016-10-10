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
	//AssociatedTemplates Templates points to the list of active templates (whether in a session or not). Without a template record a data record can not be encoded or decoded
	AssociatedTemplates *ActiveTemplates

	// TemplateID points to the template in the active templates this datarecord is adhering to
	TemplateID uint16

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

// AssociateTemplates sets the template to be used marshalling/unmarshalling this DataRecord
func (datrec *DataRecord) AssociateTemplates(at *ActiveTemplates) error {
	if at == nil {
		return fmt.Errorf("Can not use nil as Template List")
	}
	datrec.AssociatedTemplates = at
	return nil
}

// String returns the string representation of the Template Record
func (datrec *DataRecord) String() string {
	stringresult := ""
	for _, listitem := range datrec.FieldValues {
		stringresult += fmt.Sprintf("%#v", listitem)
	}
	return stringresult
}

// SetTemplateID sets the template ID the current DataRecord adheres to.
func (datrec *DataRecord) SetTemplateID(id uint16) error {
	if id < 256 {
		return fmt.Errorf("Can not use a template id < 256. Was %d", id)
	}
	datrec.TemplateID = id
	return nil
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
// FieldValues have a type when added so there is implicit information on each field value to marshal it
func (datrec *DataRecord) MarshalBinary() (data []byte, err error) {
	if datrec.AssociatedTemplates == nil {
		return nil, fmt.Errorf("Can not marshal without associated templates")
	}
	if datrec.TemplateID < 256 {
		return nil, fmt.Errorf("Can not marshal without a template id")
	}
	if len(datrec.FieldValues) < 1 {
		return nil, fmt.Errorf("Can not marshal record, must have at least one Field Value")
	}
	marshalValue := []byte{}
	curtemplate, err := datrec.AssociatedTemplates.Get(datrec.TemplateID)
	if err != nil {
		return nil, fmt.Errorf("Can not marshal record, error in retrieving template %#v", err)
	}
	for fieldidx, listitem := range datrec.FieldValues {
		item := []byte{}
		switch listitem.(type) {
		case *FieldValueSubTemplateList:
			listitem.(*FieldValueSubTemplateList).SetAssiocatedTemplates(datrec.AssociatedTemplates)
		case *FieldValueSubTemplateMultiList:
			listitem.(*FieldValueSubTemplateMultiList).SetAssiocatedTemplates(datrec.AssociatedTemplates)
		}
		item, err = listitem.MarshalBinary()
		if err != nil {
			return nil, err
		}
		if curtemplate.FieldSpecifiers[fieldidx].FieldLength != VariableLength {
			if len(item) != int(curtemplate.FieldSpecifiers[fieldidx].FieldLength) {
				return nil, fmt.Errorf("Wrong marshalled size for item %#v, expected %d, but got %d", listitem, len(item), curtemplate.FieldSpecifiers[fieldidx].FieldLength)
			}
		} else {
			var marshalLength []byte
			switch listitem.(type) {
			case *FieldValueBasicList, *FieldValueSubTemplateList, *FieldValueSubTemplateMultiList:
				marshalLength, err = EncodeVariableLength(item, true)
				if err != nil {
					return nil, err
				}

			default:
				marshalLength, err = EncodeVariableLength(item, false)
				if err != nil {
					return nil, err
				}
			}
			marshalValue = append(marshalValue, marshalLength...)
		}
		marshalValue = append(marshalValue, item...)
	}
	return marshalValue, nil
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
func (datrec *DataRecord) UnmarshalBinary(data []byte) error {
	if datrec.AssociatedTemplates == nil {
		return fmt.Errorf("Can not marshal without associated templates")
	}
	if datrec.TemplateID < 256 {
		return fmt.Errorf("Can not marshal without a template id")
	}
	if data == nil || len(data) == 0 {
		return fmt.Errorf("Can not unmarshal, invalid data. %#v", data)
	}
	curtemplate, err := datrec.AssociatedTemplates.Get(datrec.TemplateID)
	if err != nil {
		return fmt.Errorf("Can not marshal record, error in retrieving template %#v", err)
	}
	cursor := 0
	for _, recitem := range curtemplate.FieldSpecifiers {
		newval, err := NewFieldValueByID(int(recitem.EnterpriseNumber), int(recitem.InformationElementIdentifier))
		if err != nil {
			return err
		}
		switch newval.(type) {
		case *FieldValueSubTemplateList:
			newval.(*FieldValueSubTemplateList).SetAssiocatedTemplates(datrec.AssociatedTemplates)
		case *FieldValueSubTemplateMultiList:
			newval.(*FieldValueSubTemplateMultiList).SetAssiocatedTemplates(datrec.AssociatedTemplates)
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
