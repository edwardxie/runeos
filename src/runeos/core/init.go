package core

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

var (
	globarVar map[string]interface{}
	mt        = &MetaType{
		kind: make(map[string]interface{}, 0),
		info: make(map[int]string, 0),
	}
)

func init() {
	globarVar = make(map[string]interface{})
}
