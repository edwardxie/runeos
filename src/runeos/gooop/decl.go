package gooop

import (
	"reflect"
)

type Object interface {
	Propertyer
	Methoder
}

type Propertyer interface {
	Super(*Object)
}

type Methoder interface {
	Geter
	Seter
}

type Eventer interface {
	AddEvent()
	RemoveEvent()
	SeekEvent()
}

type Geter interface {
	Get(...interface{}) interface{}
}

type Seter interface {
	Set(...interface{}) error
}

type Flusher interface {
	Flush()
}

type Type struct{}

type Node struct{}

type Variant struct {
	metatype MetaType
	value    interface{}
}

type Value struct{}

type Values []Value

type object struct {
	// oid    string
	name    string
	setflag string
	parent  *Object
	lnext   *Object
	rnext   *Object
	member  []*Variant
	inject  map[reflect.Type]reflect.Value
}

type event struct {
	Object
	Methods map[string]interface{}
}

type trigger struct{}

type MetaTypeKind int

type MetaTyper interface{}

type MetaType struct {
	kind map[string]interface{}
	info map[int]string
	flag int
}
