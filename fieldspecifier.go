package ipfixmessage

import "fmt"

/*
The Information Elements are identified by the Information Element identifier.
When the Enterprise bit is set to 0, the corresponding Information Element appears in [IANA-IPFIX], and the Enterprise Number MUST NOT be present.
When the Enterprise bit is set to 1, the corresponding Information Element identifier identified an enterprise-specific Information Element; the Enterprise Number MUST be present.
*/

// FieldSpecifier defines the elements in the Field Specifier
type FieldSpecifier struct {
	E                            bool   //Enterprise bit.  This is the first bit of the Field Specifier.  If this bit is zero, the Information Element Identifier identifies an IETF-specified Information Element, and the four-octet Enterprise Number field MUST NOT be present.  If this bit is one, the Information Element identifier identifies an enterprise-specific Information Element, and the Enterprise Number filed MUST be present.
	InformationElementIdentifier uint16 //A numeric value that represents the type of Information Element. Refer to [RFC5102].
	FieldLength                  uint16 //The length of the corresponding encoded Information Element, in octets. The value 65535 is reserved for variable- length Information Elements (see Section 7).
	EnterpriseNumber             uint32 //IANA enterprise number [PEN] of the authority defining the Information Element identifier in this Template Record.
}

// NewFieldSpecifier returns a Field Specifier. If the Enterprise ID is 0 then the Enterprise Bit will not be set.
func NewFieldSpecifier(enterpriseid uint32, informationelementid, fieldlength uint16) (FieldSpecifier, error) {
	if informationelementid >= 32768 {
		return FieldSpecifier{}, fmt.Errorf("Information Element ID can not be greater than 32767, but got %d", informationelementid)
	}
	return FieldSpecifier{
		E: (enterpriseid != 0),
		InformationElementIdentifier: informationelementid,
		FieldLength:                  fieldlength,
		EnterpriseNumber:             enterpriseid,
	}, nil
}

// Len returns the length of the field specifier, in octets.
func (fsp FieldSpecifier) Len() uint16 {
	if fsp.E {
		return 8 //If the Enterprise Bit is set, we have to add the Enterprise Number
	}
	return 4
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (fsp FieldSpecifier) MarshalBinary() (data []byte, err error) {
	return nil, fmt.Errorf("Not yet implemented!")
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
// BUG(aw): NOT IMPLEMENTED
func (fsp FieldSpecifier) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) == 0 {
		return fmt.Errorf("Can not unmarshal, invalid data. %#v", data)
	}

	return fmt.Errorf("Not yet implemented!")
}
