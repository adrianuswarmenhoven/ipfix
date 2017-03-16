package ipfix

import (
	"encoding/binary"
	"fmt"
)

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
func NewFieldSpecifier(enterpriseid uint32, informationelementid, fieldlength uint16) (*FieldSpecifier, error) {
	if informationelementid > 32767 {
		return &FieldSpecifier{}, NewError(fmt.Sprintf("Information Element ID can not be greater than 32767, but got %d", informationelementid), ErrCritical)
	}
	return &FieldSpecifier{
		E: (enterpriseid != 0),
		InformationElementIdentifier: informationelementid,
		FieldLength:                  fieldlength,
		EnterpriseNumber:             enterpriseid,
	}, nil
}

// String returns the string representation of the Field Specifier
func (fsp *FieldSpecifier) String() string {
	retstr := ""

	if fsp.E {
		fd, err := FieldDescriptionByID(fsp.EnterpriseNumber, fsp.InformationElementIdentifier)
		if err != nil {
			fd = "unknown"
		}
		retstr = fmt.Sprintf("enterprise bit=yes, enterprise number=%d, information element identifier=%d ('%s'), field length=%d, ", fsp.EnterpriseNumber, fsp.InformationElementIdentifier, fd, fsp.FieldLength)
	} else {
		fd, err := FieldDescriptionByID(0, fsp.InformationElementIdentifier)
		if err != nil {
			fd = "unknown"
		}
		retstr = fmt.Sprintf("enterprise bit=no, information element identifier=%d ('%s'), field length=%d, ", fsp.InformationElementIdentifier, fd, fsp.FieldLength)
	}
	if fsp.FieldLength == VariableLength {
		return retstr + "variable length=yes"
	}
	return retstr + "variable length=no"
}

// Len returns the length of the field specifier, in octets.
func (fsp *FieldSpecifier) Len() uint16 {
	if fsp.E {
		return 8 //If the Enterprise Bit is set, we have to add the Enterprise Number
	}
	return 4
}

// MarshalBinary satisfies the encoding/BinaryMarshaler interface
func (fsp *FieldSpecifier) MarshalBinary() (data []byte, err error) {
	marshalValue := make([]byte, 0, 0)
	marshalFieldID, err := marshalBinarySingleValue(fsp.InformationElementIdentifier)
	if err != nil {
		return nil, err
	}
	if fsp.E {
		marshalFieldID[0] = marshalFieldID[0] | 128 //Setting the EnterpriseID bit
	}
	marshalValue = append(marshalValue, marshalFieldID...)
	marshalFieldLength, err := marshalBinarySingleValue(fsp.FieldLength)
	if err != nil {
		return nil, err
	}
	marshalValue = append(marshalValue, marshalFieldLength...)

	if fsp.E {
		marshalEnterpriseID, err := marshalBinarySingleValue(fsp.EnterpriseNumber)
		if err != nil {
			return nil, err
		}
		marshalValue = append(marshalValue, marshalEnterpriseID...)
	}

	return marshalValue, nil
}

// UnmarshalBinary satisfies the encoding/BinaryUnmarshaler interface
func (fsp *FieldSpecifier) UnmarshalBinary(data []byte) error {
	if data == nil || len(data) == 0 {
		return NewError(fmt.Sprintf("Can not unmarshal, invalid data. %#v", data), ErrCritical)
	}
	if (data[0] & 128) != 0 {
		fsp.E = true
		data[0] = data[0] & 127 //Remove the bit
	} else {
		fsp.E = false
		fsp.EnterpriseNumber = 0
	}
	fsp.InformationElementIdentifier = binary.BigEndian.Uint16(data[0:2])
	if fsp.E {
		data[0] = data[0] | 128 //Restore the bit (so we keep the original datablob)
		fsp.EnterpriseNumber = binary.BigEndian.Uint32(data[4:])
	}
	fsp.FieldLength = binary.BigEndian.Uint16(data[2:4])
	return nil
}
