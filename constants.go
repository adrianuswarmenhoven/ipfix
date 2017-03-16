package ipfixmessage

const (
	// IPFIXVersion denotes the version of Flow Record format exported in this message.
	// The value of this field is 10 for the current version, incrementing by one the version used in the NetFlow services export version 9 [RFC3954].
	IPFIXVersion = 10

	//EnterpriseBitSet is the base value if the Enterprise Bit is set
	EnterpriseBitSet = 32768
)

// SetID values
const (
	SetIDInvalid1       = 0 // RFC3954
	SetIDInvalid2       = 1 // RFC3954
	SetIDTemplate       = 2 // Denotes the set is a template
	SetIDOptionTemplate = 3 // Denotes the set is an option template
)

const (
	//VariableLength - The value 65535 is reserved for variable-length Information Elements (see Section 7).
	VariableLength = 65535
)

// BoolTrue & BoolFalse provide the boolean data type as it
// is specified according to the TruthValue in [RFC2579]:
// it is an integer with the value 1 for true and a value 2
// for false.  Every other value is undefined.  The boolean data type
// MUST be encoded in a single octet.
const (
	BoolTrue  = 1
	BoolFalse = 2
)

//'Package internal' constants
const (
	ipfixMessageHeaderLength = 16 //Length of the message header. For calculations.
	ipfixSetHeaderLength     = 4  // Length of a set header, For calculations.
)

//The Semantic field indicates the relationship among the different Information Element values within this Structured Data Information Element.
//Refer to IANA's "IPFIX Structured Data Types Semantics" registry.
const (
	NoneOf       = 0x00 //The "noneOf" structured data type semantic specifies that none of the elements are actual properties of the Data Record.	[RFC6313]
	ExactlyOneOf = 0x01 //The "exactlyOneOf" structured data type semantic specifies that only a single element from the structured data is an actual property of the Data Record. This is equivalent to a logical XOR operation.	[RFC6313]
	OneOrMoreOf  = 0x02 //The "oneOrMoreOf" structured data type semantic specifies that one or more elements from the list in the structured data are actual properties of the Data Record. This is equivalent to a logical OR operation.	[RFC6313]
	AllOf        = 0x03 //The "allOf" structured data type semantic specifies that all of the list elements from the structured data are actual properties of the Data Record.	[RFC6313]
	Ordered      = 0x04 //The "ordered" structured data type semantic specifies that elements from the list in the structured data are ordered.	[RFC6313]
	Undefined    = 0xFF //The "undefined" structured data type semantic specifies that the semantic of the list elements is not specified and that, if a semantic exists, then it is up to the Collecting Process to draw its own conclusions. The "undefined" structured data type semantic is the default structured data type semantic.	[RFC6313]
)
