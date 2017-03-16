package ipfix

import "fmt"

//BasicList represents a list of zero or more instances of any Information Element, primarily used for single-valued data types.
//Examples include a list of port numbers, a list of interface indexes, a list of AS in a BGP AS-PATH, etc.
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
		return nil, NewError(fmt.Sprintf("Information Element ID can not be greater than 32767, but got %d", informationelementid), ErrCritical)
	}
	if semantic >= 0x05 && semantic <= 0xFE {
		return nil, NewError(fmt.Sprintf("Semantic undefined: %d", semantic), ErrCritical)
	}
	return &BasicList{
		Semantic: semantic,
		E:        (enterpriseid != 0),
		InformationElementIdentifier: informationelementid,
		FieldLength:                  fieldlength,
		EnterpriseNumber:             enterpriseid,
		FieldValues:                  make([]FieldValue, 0, 0),
	}, nil
}

// Len returns the length of the field specifier, in octets.
func (blst *BasicList) Len() uint16 {
	bllen := uint16(0)
	if blst.E {
		bllen = 9 //If the Enterprise Bit is set, we have to add the Enterprise Number
	} else {
		bllen = 5
	}
	for _, listitem := range blst.FieldValues {
		bllen += listitem.Len()
		if blst.FieldLength == VariableLength {
			switch listitem.(type) {
			//RFC6313 Section 4.5.1 RECOMMENDED
			case *FieldValueBasicList, *FieldValueSubTemplateList, *FieldValueSubTemplateMultiList:
				bllen += 3
			default:
				if listitem.Len() < 255 {
					bllen++
				} else {
					bllen += 3
				}
			}
		}
	}
	return bllen
}
