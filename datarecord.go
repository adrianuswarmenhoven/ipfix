package ipfixmessage

//--- Data Record

type IPFIXDataRecord struct {
	FieldValue []IPFIXFieldValue //Note that Field Values do not necessarily have a length of 16 bits. Field Values are encoded according to their data type specified in [RFC5102].
}
