package gooop

import (
	"reflect"
)

type Value struct {
	value map[reflect.Type]reflect.Value
}

type Values []Value
