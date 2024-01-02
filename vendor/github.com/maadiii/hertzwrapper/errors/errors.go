package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

type Error struct {
	Message string `json:"message,omitempty"`
	Params  []any  `json:"params,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

var _ error = (*Error)(nil)

// example: e.Pattern("foo field must be in [0] and [1]", 10, 20)
func (e *Error) Pattern(value string, params ...any) error {
	stack := stack()
	e.Message = fmt.Sprintf("%s\n%s", value, strings.Join(stack, "\n"))
	e.Params = append(e.Params, params...)

	return e
}

func Is(err, target error) bool {
	var errMsg, targetMsg string

	if err != nil {
		errMsg = err.Error()
	}

	if target != nil {
		targetMsg = target.Error()
	}

	return strings.Contains(errMsg, targetMsg)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func New(text string) *Error {
	stack := stack()

	return &Error{
		Message: fmt.Sprintf("%s\n%s", errors.New(text), strings.Join(stack, "\n")),
	}
}

func Wrap(err error) *Error {
	stack := stack()
	e := &Error{
		Message: fmt.Sprintf("%s\n%s", err.Error(), strings.Join(stack, "\n")),
	}

	return e
}

// func Join(errs ...error) *Error {
// 	return Wrap(&Error{LowLevel: errors.Join(errs...)})
// }
//
// func Unwrap(err error) *Error {
// 	if e, ok := err.(*Error); ok {
// 		e.LowLevel = errors.Unwrap(e.LowLevel)
//
// 		return e
// 	}
//
// 	e := &Error{LowLevel: errors.Unwrap(err)}
//
// 	return e
// }

func stack() []string {
	buf := make([]byte, 2048)
	runtime.Stack(buf, false)

	stack := strings.Split(string(buf), "\n")

	return stack[6:]
}

var (
	BadRequest                   = &Error{} //nolint
	Unauthorized                 = &Error{} //nolint
	PaymentRequired              = &Error{} //nolint
	Forbidden                    = &Error{} //nolint
	NotFound                     = &Error{} //nolint
	MethodNotAllowed             = &Error{} //nolint
	NotAcceptable                = &Error{} //nolint
	ProxyAuthRequired            = &Error{} //nolint
	RequestTimeout               = &Error{} //nolint
	Conflict                     = &Error{} //nolint
	Gone                         = &Error{} //nolint
	LengthRequired               = &Error{} //nolint
	PreconditionFailed           = &Error{} //nolint
	RequestEntityTooLarge        = &Error{} //nolint
	RequestURITooLong            = &Error{} //nolint
	UnsupportedMediaType         = &Error{} //nolint
	RequestedRangeNotSatisfiable = &Error{} //nolint
	ExpectationFailed            = &Error{} //nolint
	Teapot                       = &Error{} //nolint
	MisdirectedRequest           = &Error{} //nolint
	UnprocessableEntity          = &Error{} //nolint
	Locked                       = &Error{} //nolint
	FailedDependency             = &Error{} //nolint
	TooEarly                     = &Error{} //nolint
	UpgradeRequired              = &Error{} //nolint
	PreconditionRequired         = &Error{} //nolint
	TooManyRequests              = &Error{} //nolint
	RequestHeaderFieldsTooLarge  = &Error{} //nolint
	UnavailableForLegalReasons   = &Error{} //nolint

	InternalServerError           = &Error{} //nolint
	NotImplemented                = &Error{} //nolint
	BadGateway                    = &Error{} //nolint
	ServiceUnavailable            = &Error{} //nolint
	GatewayTimeout                = &Error{} //nolint
	HTTPVersionNotSupported       = &Error{} //nolint
	VariantAlsoNegotiates         = &Error{} //nolint
	InsufficientStorage           = &Error{} //nolint
	LoopDetected                  = &Error{} //nolint
	NotExtended                   = &Error{} //nolint
	NetworkAuthenticationRequired = &Error{} //nolint
)
