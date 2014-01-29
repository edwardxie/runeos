package gooop

import (
	"fmt"
)

type Eventer interface {
	On()
}

type event struct {
	Object
	Methods map[string]interface{}
}

func NewEvent(obj Object) Eventer {
	et := &event{
		Object:  obj,
		Methods: make(map[string]interface{}, 0),
	}
	return Eventer(et)
}
func (e *event) On() {}

func AddEvent() { fmt.Println("Add event.") }

func RemoveEvent() {}

func SeekEvent() {}
