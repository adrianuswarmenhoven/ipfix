package ipfix

import "fmt"

/*

Auxiliary functions

*/

// EncodeVariableLength returns the bytes for encoding a variable length as specified in RFC 7011, section 7
// In the Template Set, the Information Element Field Length is recorded as 65535.
// This reserved length value notifies the Collecting Process that the length value of the Information Element will be carried in the Information Element content itself.
// In most cases, the length of the Information Element will be less than 255 octets. In this case 1 byte is sufficient to encode the length.
// The length may also be encoded into 3 octets before the Information Element, allowing the length of the Information Element to be greater than or equal to 255 octets.
// In this case, the first octet of the Length field MUST be 255, and the length is carried in the second and third octets.
// The octets carrying the length (either the first or the first three octets) MUST NOT be included in the length of the Information Element.
// If rfc6313recommended is true, then the length is always encoded as {255,high byte,low byte}, otherwise it is encoded as stated above
func EncodeVariableLength(content []byte, rfc6313recommended bool) ([]byte, error) {
	retval := []byte{}
	if len(content) < 255 && !rfc6313recommended {
		retval = []byte{uint8(len(content))}
	} else {
		if len(content) > 65535 {
			return []byte{}, NewError(fmt.Sprintf("Content too large, maximum of 65535 octets, but it is %d", len(content)), ErrCritical)
		}
		lengthBytes := []byte{255}
		lengthContentBytes, err := marshalBinarySingleValue(uint16(len(content)))
		if err != nil {
			return []byte{}, err
		}
		retval = append(lengthBytes, lengthContentBytes...)
	}
	return retval, nil
}

// DecodeVariableLength returns the length for decoding a variable length as specified in RFC 7011, section 7
// cursorshift denotes the number of positions needed to get the encoded length
func DecodeVariableLength(content []byte) (uint16, uint8, error) {
	cursorshift := uint8(0)
	retval := uint16(0)
	if content[0] == 0 {
		return 0, 0, NewError("Content can not be 0 in length.", ErrCritical)
	}
	if content[0] < 255 {
		retval = uint16(content[0])
		cursorshift = 1
	} else {
		retval = uint16(256*uint16(content[1])) + uint16(content[2])
		cursorshift = 3
	}
	return retval, cursorshift, nil
}
