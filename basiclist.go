package ipfixmessage

import "fmt"

type BasicList struct {
	Semantic                     uint8  //one of: SemanticsNoneOf, ExactlyOneOf, OneOrMoreOf, AllOf, Ordered or Undefined
	E                            bool   //Enterprise bit.  This is the first bit of the Field Specifier.  If this bit is zero, the Information Element Identifier identifies an IETF-specified Information Element, and the four-octet Enterprise Number field MUST NOT be present.  If this bit is one the Information Element identifier identifies an enterprise-specific Information Element, and the Enterprise Number filed MUST be present.
	InformationElementIdentifier uint16 //A numeric value that represents the type of Information Element. Refer to [RFC5102].
	FieldLength                  uint16 //The length of the corresponding encoded Information Element, in octets. The value 65535 is reserved for variable- length Information Elements (see Section 7).
	EnterpriseNumber             uint32 //IANA enterprise number [PEN] of the authority defining the Information Element identifier in this Template Record.

	FieldValues []FieldValue //The list of Field Values
}

// NewBasicList returns a BasicList. If the Enterprise ID is 0 then the Enterprise Bit will not be set.
func NewBasicList(semantic uint8, enterpriseid uint32, informationelementid, fieldlength uint16) (*BasicList, error) {
	if informationelementid >= 32768 {
		return nil, fmt.Errorf("Information Element ID can not be greater than 32767, but got %d", informationelementid)
	}
	if semantic >= 0x05 && semantic <= 0xFE {
		return nil, fmt.Errorf("Semantic undefined: %d", semantic)
	}
	return &BasicList{
		Semantic: semantic,
		E:        (enterpriseid != 0),
		InformationElementIdentifier: informationelementid,
		FieldLength:                  fieldlength,
		EnterpriseNumber:             enterpriseid,
	}, nil
}

// Len returns the length of the field specifier, in octets.
func (blst *BasicList) Len() uint16 {
	bllen := uint16(0)
	if blst.E {
		bllen = 8 //If the Enterprise Bit is set, we have to add the Enterprise Number
	} else {
		bllen = 4
	}
	for _, listitem := range blst.FieldValues {
		bllen += listitem.Len()
	}
	return bllen
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (blst *BasicList) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (blst *BasicList) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("Can not unmarshal, invalid data. %#v", data)
	}

	return fmt.Errorf("Not yet implemented!")
}
