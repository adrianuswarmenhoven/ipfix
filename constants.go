package ipfixmessage

const (
	// IPFIXVersion denotes the version of Flow Record format exported in this message.
	// The value of this field is 10 for the current version, incrementing by one the version used in the NetFlow services export version 9 [RFC3954].
	IPFIXVersion = 10

	//EnterpriseBitSet is the base value if the Enterprise Bit is set
	EnterpriseBitSet = 32768

	// SetID values
	SetIDInvalid1       = 0 // RFC3954
	SetIDInvalid2       = 1 // RFC3954
	SetIDTemplate       = 2 // Denotes the set is a template
	SetIDOptionTemplate = 3 // Denotes the set is an option template

	VariableLength = 65535 //The value 65535 is reserved for variable-length Information Elements (see Section 7).

	// BoolTrue & BoolFalse provide the boolean data type as it
	// is specified according to the TruthValue in [RFC2579]:
	// it is an integer with the value 1 for true and a value 2
	// for false.  Every other value is undefined.  The boolean data type
	// MUST be encoded in a single octet.
	BoolTrue  = 1
	BoolFalse = 2

	//'Package internal' constants
	ipfixMessageHeaderLength = 16 //Length of the message header. For calculations.
	ipfixSetHeaderLength     = 4  // Length of a set header, For calculations.
)
