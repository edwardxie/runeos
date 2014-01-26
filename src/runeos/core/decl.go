package core

type Object interface {
	Propertyer
	Methoder
	Eventer
}

type Propertyer interface {
	Id() string
	Name() string
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
	Get() interface{}
}

type Seter interface {
	Set(interface{})
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
	oid    string
	name   string
	parent *Object
	data   *Variant
}

type event struct {
	Object
	Methods map[string]interface{}
}
