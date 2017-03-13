package ipfixmessage

import "fmt"

const (
	//MaxMoreErrors - Maximum of stacked errors before 'More Errors' is used. This is to prevent out-of-memory issues
	MaxMoreErrors = 5
)

//ErrXXX are the various severities of the errors
const (
	ErrINFO     = iota //Not really an error
	ErrFailure         //When unmarshalling a single element failed
	ErrCritical        //When code can not or should not continue
)

//ProtocolError is a custom error message that can stack multiple errors
type ProtocolError struct {
	SubError    []ProtocolError
	Severity    int
	Description string
	MoreErrors  int
}

//Error implements the error interface
func (err *ProtocolError) Error() string {
	ret := fmt.Sprintf("%d - %s ", err.Severity, err.Description)
	for _, sube := range err.SubError {
		ret += "{" + sube.Error() + "}"
	}
	if err.MoreErrors > 0 {
		ret += fmt.Sprintf("{%d more errors}", err.MoreErrors)
	}
	return ret
}

//NewError returns a new protocol error
func NewError(desc string, sev int) *ProtocolError {
	return &ProtocolError{
		SubError:    []ProtocolError{},
		Severity:    sev,
		Description: desc,
	}
}

//Stack stacks an error on top of the current error
func (err *ProtocolError) Stack(stackerr interface{}) {
	if len(err.SubError) >= MaxMoreErrors {
		err.MoreErrors++
		return
	}
	switch stackerr.(type) {
	case ProtocolError:
		err.SubError = append(err.SubError, stackerr.(ProtocolError))
	case *ProtocolError:
		err.SubError = append(err.SubError, *(stackerr.(*ProtocolError)))
	default:
		err.SubError = append(err.SubError, *(NewError(fmt.Sprintf("%+v", err), ErrINFO)))
	}
}
