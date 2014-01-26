package gooop

import (
	"fmt"
)

func NewEvent(obj Object) Eventer {
	et := &event{
		Object:  obj,
		Methods: make(map[string]interface{}, 0),
	}
	return Eventer(et)
}

func (e *event) AddEvent() { fmt.Println("Add event.") }

func (e *event) RemoveEvent() {}

func (e *event) SeekEvent() {}
