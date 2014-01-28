package gooop

import (
	"fmt"
)

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

type Error struct {
	code   int
	errmsg string
}

func NewError(msg string) *Error { return &Error{ErrCustomKind, msg} }

func (e *Error) Error() string { return fmt.Sprintf("%v: %v", e.code, e.errmsg) }

// func (e *Error) String() string { return e.errmsg }

func (e *Error) Set(params ...interface{}) error { return nil }

func (e *Error) Get(params ...interface{}) interface{} { return nil }

func (e *Error) Super(super *Object) {}
