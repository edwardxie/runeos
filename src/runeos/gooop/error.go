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
	ErrNo           = &Error{ErrNoKind, "No error.", nil}
	ErrCustom       = &Error{ErrCustomKind, "Custom error.", nil}
	ErrObject       = &Error{ErrObjectKind, "Object internal error.", nil}
	ErrEvent        = &Error{ErrEventKind, "Event internal error.", nil}
	ErrMethod       = &Error{ErrMethodKind, "Method internal error.", nil}
	ErrParamNil     = &Error{ErrParamNilKind, "Param is nil.", nil}
	ErrParamInvalid = &Error{ErrParamInvalidKind, "Param is invalid.", nil}
	ErrParamTooFew  = &Error{ErrParamTooFewKind, "Too few param.", nil}
	ErrParamTooMany = &Error{ErrParamTooManyKind, "Too many param.", nil}
)

type Error struct {
	code   int
	errmsg string
	rptr   *Object //Retrieve a pointer, reservation no implementation.
}

func NewError(msg string) *Error { return &Error{ErrCustomKind, msg, nil} }

func (e *Error) Error() string { return fmt.Sprintf("%v: %v", e.code, e.errmsg) }

// func (e *Error) String() string { return e.errmsg }

func (e *Error) Set(...interface{}) error { return nil }

func (e *Error) Get(params ...interface{}) interface{} { return e.rptr }

func (e *Error) Super(super *Object) { e.rptr = super }
