package ipfixmessage

const (
	//Version of Flow Record format exported in this message.
	//The value of this field is 0x000a for the current version, incrementing by one the version used in the NetFlow services export version 9 [RFC3954].
	IPFIXVersion = 0x000a

	SetID_Invalid1       = 0 //RFC3954
	SetID_Invalid2       = 1 //RFC3954
	SetID_Template       = 2
	SetID_OptionTemplate = 3

	VariableLength = 65535 //The value 65535 is reserved for variable-length Information Elements (see Section 7).

	//The boolean data type is specified according to the TruthValue in
	//[RFC2579]: it is an integer with the value 1 for true and a value 2
	//for false.  Every other value is undefined.  The boolean data type
	//MUST be encoded in a single octet.
	BoolTrue  = 1
	BoolFalse = 2

	ipfixMessageHeaderLength = 128 //Length of the message header. For calculations.
	ipfixSetHeaderLength     = 32
)
