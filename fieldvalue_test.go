package ipfixmessage

import (
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"
)

func TestFieldValueValueGoTypes(t *testing.T) {
	testsetmatch := map[interface{}]int{
		FieldValueUnsigned8{}.Value:            tg_uint8,
		FieldValueUnsigned16{}.Value:           tg_uint16,
		FieldValueUnsigned32{}.Value:           tg_uint32,
		FieldValueUnsigned64{}.Value:           tg_uint64,
		FieldValueSigned8{}.Value:              tg_int8,
		FieldValueSigned16{}.Value:             tg_int16,
		FieldValueSigned32{}.Value:             tg_int32,
		FieldValueSigned64{}.Value:             tg_int64,
		FieldValueFloat32{}.Value:              tg_float32,
		FieldValueFloat64{}.Value:              tg_float64,
		FieldValueBoolean{}.Value:              tg_bool,
		FieldValueString{}.Value:               tg_string,
		FieldValueDateTimeSeconds{}.Value:      tg_time,
		FieldValueDateTimeMilliseconds{}.Value: tg_time,
		FieldValueDateTimeMicroseconds{}.Value: tg_time,
		FieldValueDateTimeNanoseconds{}.Value:  tg_time,
	}
	for testval, checkval := range testsetmatch {
		if goTypeName(testval) != checkval {
			t.Errorf("Should have gotten %d but got %d for %#v", checkval, goTypeName(testval), testval)
		}
	}

	//Now checking the unhashables
	x := reflect.TypeOf(FieldValueMacAddress{}.Value)
	fmt.Println(x)

}

type testfv struct {
	Value interface{}
}

func goTypeName(fv interface{}) int {
	switch fv.(type) {
	case uint8:
		return tg_uint8
	case uint16:
		return tg_uint16
	case uint32:
		return tg_uint32
	case uint64:
		return tg_uint64
	case int8:
		return tg_int8
	case int16:
		return tg_int16
	case int32:
		return tg_int32
	case int64:
		return tg_int64
	case float32:
		return tg_float32
	case float64:
		return tg_float64
	case bool:
		return tg_bool
	case net.HardwareAddr:
		return tg_mac
	case []byte:
		return tg_byteslice
	case string:
		return tg_string
	case time.Time:
		return tg_time
	case net.IP:
		return tg_ipaddress
	default:
		return t_unknown
	}
}
