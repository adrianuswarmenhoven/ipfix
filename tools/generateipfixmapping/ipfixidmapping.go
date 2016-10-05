//Generated, do not edit
package ipfixmessage

import (
	"fmt"
)

// NewFieldValueByID returns an empty FieldValue that matches the enterprise id and element id
func NewFieldValueByID(enterpriseid int, elementid int) (FieldValue, error) {
	switch enterpriseid {
{{range $_, $enterpriseid := .EnterpriseOrder}}
case {{$enterpriseid}}: // {{index $.Sources $enterpriseid}}
{{$elements := (index $.Elements $enterpriseid)}}
    switch elementid { {{range $_, $elementid:= (index $.ElementsOrder $enterpriseid)}}
        case {{$elementid}}:
        return &{{(index $elements $elementid).GoFieldValue}}{}, nil // {{(index $elements $elementid).Name}}{{end}}
        default:
           return nil,fmt.Errorf("No such element: E%did%d",enterpriseid,elementid)
    }
{{end}}
	default:
		return nil, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
	}
}

// FieldLengthByID returns the default length that matches the enterprise id and element id
func FieldLengthByID(enterpriseid int, elementid int) (uint16, error) {
	switch enterpriseid {
{{range $_, $enterpriseid := .EnterpriseOrder}}
case {{$enterpriseid}}: // {{index $.Sources $enterpriseid}}
{{$elements := (index $.Elements $enterpriseid)}}
    switch elementid { {{range $_, $elementid:= (index $.ElementsOrder $enterpriseid)}}
        case {{$elementid}}:
        return {{(index $elements $elementid).GoFieldLength}}, nil // {{(index $elements $elementid).Name}}{{end}}
        default:
           return 0,fmt.Errorf("No such element: E%did%d",enterpriseid,elementid)
    }
{{end}}
	default:
		return 0, fmt.Errorf("No such element: E%did%d", enterpriseid, elementid)
	}
    return 0,nil
}