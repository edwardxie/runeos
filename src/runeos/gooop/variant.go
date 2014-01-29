package gooop

import (
	"reflect"
)

type Variant struct {
	flag *Flag
	typ  *MetaType
	data map[reflect.Type]reflect.Value
}

func newVariant(v interface{}) *Variant {
	tmp := make(map[reflect.Type]reflect.Value, 0)
	// for _, v := range vs {
	tmp[reflect.TypeOf(v)] = reflect.ValueOf(v)
	// }
	return &Variant{data: tmp}
}
