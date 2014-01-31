package gooop

import (
// "reflect"
)

// type Pusher interface {
//  Push(...interface{})
// }

// type Poper interface {
//  Pop(interface{}) *Object
// }

// type Remover interface {
//  Remove(interface{}) error
// }

type Variant struct {
	name string
	obj  *Object
	flag *Flag
	typ  *MetaType
	data *Value
}

func NewVariant(params ...interface{}) *Variant {
	// tmp := make(map[reflect.Type]reflect.Value, 0)
	// for _, v := range vs {
	// tmp[reflect.TypeOf(v)] = reflect.ValueOf(v)
	// }
	return &Variant{}
}
