package core

import (
	// "container/list"
	// "fmt"
	"runeos/utils"
)

func NewObject(name string) Object {
	obj := object{
		oid:  utils.NewID().UUID(),
		name: name,
		// data: list.New(),
	}
	return Object(&obj)
}

func (o *object) Id() string { return o.oid }

func (o *object) Name() string { return o.name }

func (o *object) Set(v interface{}) {}

func (o *object) Get() interface{} { return nil }

func (o *object) Flush() {}

func (o *object) AddEvent() {}

func (o *object) RemoveEvent() {}

func (o *object) SeekEvent() {}
