package ipfixmessage

const (
	testMarkerString = "**************************************************\n* %s\n**************************************************\n"

	errorPrefixMarker = "\033[0;31m###\033[0m"
)

//An enumeration of the types, mainly for testing purposes
const (
	tUnknown = iota
	tOctetArray
	tUnsigned8
	tUnsigned16
	tUnsigned32
	tUnsigned64
	tSigned8
	tSigned16
	tSigned32
	tSigned64
	tFloat32
	tFloat64
	tBoolean
	tmacAddress
	tString
	tDateTimeSeconds
	tDateTimeMilliseconds
	tDateTimeMicroseconds
	tDateTimeNanoseconds
	tIPv4Address
	tIPv6Address
	tBasicList
	tSubTemplateList
	tSubTemplateMultiList

	tgUint8
	tgUint16
	tgUint32
	tgUint64
	tgInt8
	tgInt16
	tgInt32
	tgInt64
	tgFloat32
	tgFloat64
	tgBool
	tgMAC
	tgByteslice
	tgString
	tgTime
	tgIPAddress
)
