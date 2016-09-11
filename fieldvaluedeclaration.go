package ipfixmessage

type IPFIXFieldValue interface {
}

//The type "unsigned8" represents a non-negative integer value in the range of 0 to 255.
type IPFIXFieldValue_unsigned8 struct {
}

//The type "unsigned16" represents a non-negative integer value in the range of 0 to 65535.
type IPFIXFieldValue_unsigned16 struct {
}

//The type "unsigned32" represents a non-negative integer value in the range of 0 to 4294967295
type IPFIXFieldValue_unsigned32 struct {
}

//The type "unsigned64" represents a non-negative integer value in the range of 0 to 18446744073709551615
type IPFIXFieldValue_unsigned64 struct {
}

//The type "signed8" represents an integer value in the range of -128 to 127
type IPFIXFieldValue_signed8 struct {
}

//The type "signed16" represents an integer value in the range of -32768 to 32767.
type IPFIXFieldValue_signed16 struct {
}

//The type "signed32" represents an integer value in the range of -2147483648 to 2147483647.
type IPFIXFieldValue_signed32 struct {
}

//The type "signed64" represents an integer value in the range of -9223372036854775808 to 9223372036854775807.
type IPFIXFieldValue_signed64 struct {
}

//The type "float32" corresponds to an IEEE single-precision 32-bit floating-point type as defined in [IEEE.754.2008].
type IPFIXFieldValue_float32 struct {
}

//The type "float64" corresponds to an IEEE double-precision 64-bit floating-point type as defined in [IEEE.754.2008].
type IPFIXFieldValue_float64 struct {
}

//The type "boolean" represents a binary value.  The only allowed values are "true" and "false".
type IPFIXFieldValue_boolean struct {
}

//The type "macAddress" represents a MAC-48 address as defined in [IEEE.802-3.2012].
type IPFIXFieldValue_macAddress struct {
}

//The type "octetArray" represents a finite-length string of octets.
type IPFIXFieldValue_octetArray struct {
}

//The type "string" represents a finite-length string of valid characters from the Unicode coded character set [ISO.10646].
//Unicode incorporates ASCII [RFC20] and the characters of many other international character sets.
type IPFIXFieldValue_string struct {
}

//The type "dateTimeSeconds" represents a time value expressed with second-level precision.
type IPFIXFieldValue_dateTimeSeconds struct {
}

//The type "dateTimeMilliseconds" represents a time value expressed with millisecond-level precision.
type IPFIXFieldValue_dateTimeMilliseconds struct {
}

//The type "dateTimeMicroseconds" represents a time value expressed with microsecond-level precision.
type IPFIXFieldValue_dateTimeMicroseconds struct {
}

//The type "dateTimeNanoseconds" represents a time value expressed with nanosecond-level precision.
type IPFIXFieldValue_dateTimeNanoseconds struct {
}

//The type "ipv4Address" represents an IPv4 address.
type IPFIXFieldValue_ipv4Address struct {
}

//The type "ipv6Address" represents an IPv6 address.
type IPFIXFieldValue_ipv6Address struct {
}

//The type "basicList" supports structured data export as described in [RFC6313];
//see Section 4.5.1 of that document for encoding details.
type IPFIXFieldValue_basicList struct {
}

//The type "subTemplateList" supports structured data export as described in [RFC6313];
//see Section 4.5.2 of that document for encoding details.
type IPFIXFieldValue_subTemplateList struct {
}

//The type "subTemplateMultiList" supports structured data export as described in [RFC6313];
//see Section 4.5.3 of that document for encoding details.
type IPFIXFieldValue_subTemplateMultiList struct {
}
