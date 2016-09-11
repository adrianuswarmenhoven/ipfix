package ipfixmessage

type IPFIXFieldSpecifier struct {
	E                            bool   //Enterprise bit.  This is the first bit of the Field Specifier.  If this bit is zero, the Information Element Identifier identifies an IETF-specified Information Element, and the four-octet Enterprise Number field MUST NOT be present.  If this bit is one, the Information Element identifier identifies an enterprise-specific Information Element, and the Enterprise Number filed MUST be present.
	InformationElementIdentifier uint16 //A numeric value that represents the type of Information Element. Refer to [RFC5102].
	FieldLength                  uint16 //The length of the corresponding encoded Information Element, in octets. The value 65535 is reserved for variable- length Information Elements (see Section 7).
	EnterpriseNumber             uint32 //IANA enterprise number [PEN] of the authority defining the Information Element identifier in this Template Record.
}
