package ipfixmessage

//An enumeration of the types, mainly for testing purposes
const (
	t_unknown = iota
	t_octetArray
	t_unsigned8
	t_unsigned16
	t_unsigned32
	t_unsigned64
	t_signed8
	t_signed16
	t_signed32
	t_signed64
	t_float32
	t_float64
	t_boolean
	t_macAddress
	t_string
	t_dateTimeSeconds
	t_dateTimeMilliseconds
	t_dateTimeMicroseconds
	t_dateTimeNanoseconds
	t_ipv4Address
	t_ipv6Address
	t_basicList
	t_subTemplateList
	t_subTemplateMultiList

	tg_uint8
	tg_uint16
	tg_uint32
	tg_uint64
	tg_int8
	tg_int16
	tg_int32
	tg_int64
	tg_float32
	tg_float64
	tg_bool
	tg_mac
	tg_byteslice
	tg_string
	tg_time
	tg_ipaddress
)

var (
	StringToType map[string]int = map[string]int{
		"octetArray":           t_octetArray,
		"unsigned8":            t_unsigned8,
		"unsigned16":           t_unsigned16,
		"unsigned32":           t_unsigned32,
		"unsigned64":           t_unsigned64,
		"signed8":              t_signed8,
		"signed16":             t_signed16,
		"signed32":             t_signed32,
		"signed64":             t_signed64,
		"float32":              t_float32,
		"float64":              t_float64,
		"boolean":              t_boolean,
		"macAddress":           t_macAddress,
		"string":               t_string,
		"dateTimeSeconds":      t_dateTimeSeconds,
		"dateTimeMilliseconds": t_dateTimeMilliseconds,
		"dateTimeMicroseconds": t_dateTimeMicroseconds,
		"dateTimeNanoseconds":  t_dateTimeNanoseconds,
		"ipv4Address":          t_ipv4Address,
		"ipv6Address":          t_ipv6Address,
		"basicList":            t_basicList,
		"subTemplateList":      t_subTemplateList,
		"subTemplateMultiList": t_subTemplateMultiList,
	}
)
