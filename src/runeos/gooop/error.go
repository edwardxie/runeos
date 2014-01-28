package gooop

const (
	ErrNoKind = iota
	ErrCustomKind
	ErrObjectKind
	ErrEventKind
	ErrMethodKind
	ErrParamNilKind
	ErrParamInvalidKind
	ErrParamTooFewKind
	ErrParamTooManyKind
)

var (
	ErrNo           = &Error{ErrNoKind, "No error."}
	ErrCustom       = &Error{ErrCustomKind, "Custom error."}
	ErrObject       = &Error{ErrObjectKind, "Object internal error."}
	ErrEvent        = &Error{ErrEventKind, "Event internal error."}
	ErrMethod       = &Error{ErrMethodKind, "Method internal error."}
	ErrParamNil     = &Error{ErrParamNilKind, "Param is nil."}
	ErrParamInvalid = &Error{ErrParamInvalidKind, "Param is invalid."}
	ErrParamTooFew  = &Error{ErrParamTooFewKind, "Too few param."}
	ErrParamTooMany = &Error{ErrParamTooManyKind, "Too many param."}
)

func NewError(msg string) *Error {
	return &Error{ErrCustomKind, msg}
}

func ContainError(err ...Error) bool {
	r := false
	return r
}

func (e *Error) String() string {
	return e.errmsg
}
