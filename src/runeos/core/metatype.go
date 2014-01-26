package core

type MetaTypeKind int

const (
	Array MetaTypeKind = 100 << iota
	ArrayPrt
	Func
	FuncPrt
	Struct
	StructPrt
	Interface
	Map
	Chan
)

const (
	PropertyType MetaTypeKind = 1000 << iota
	EventType
	MethodType
)

type MetaTyper interface{}

type MetaType struct {
	kind map[string]interface{}
	info map[int]string
	flag int
}

var mt = &MetaType{
	kind: make(map[string]interface{}, 0),
	info: make(map[int]string, 0),
}

func (m *MetaType) RegisterMetaType(name string, typ interface{}) {

}

func (m *MetaType) UnregisterMetaType(name string) {

}
