package gooop

type MetaTypeKind int

type MetaTyper interface{}

type MetaType struct {
	kind map[string]interface{}
	info map[int]string
	flag int
}

func (m *MetaType) RegisterMetaType(name string, typ interface{}) {}

func (m *MetaType) UnregisterMetaType(name string) {}
