package ipfixmessage

//--- Data Record

// DataRecord - A Data Record is a record that contains values of the parameters corresponding to a Template Record.
// The Data Records are sent in Data Sets which consist only of one or more Field Values.
// The Template ID to which the Field Values belong is encoded in the Set Header field "Set ID", i.e., "Set ID" == "Template ID".
// Note that Field Values do not necessarily have a length of 16 bits.
// Field Values are encoded according to their data type as specified in [RFC7012].
// Interpretation of the Data Record format can be done only if the Template Record corresponding to the Template ID is available at the Collecting Process.
type DataRecord struct {
	FieldValues []FieldValue //Note that Field Values do not necessarily have a length of 16 bits. Field Values are encoded according to their data type specified in [RFC5102].
}
