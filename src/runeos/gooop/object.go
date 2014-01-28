package gooop

import (
	// "container/list"
	"fmt"
	// "reflect"
	// "runeos/utils"
)

func NewObject(params ...interface{}) Object {
	isName := false
	obj := object{}
	if len(params) == 0 {
		return Object(&obj)
	}
	for inx, param := range params {
		switch typ := param.(type) {
		case string:
			if isName {
				return Object(ErrParamTooMany)
			}
			obj.name = param.(string)
			isName = true
		default:
			fmt.Printf("Params for switch is %#v [index:type] => [%v:%T]\n ", params, inx, typ)
		}
	}
	return Object(&obj)
}

func (o *object) Set(params ...interface{}) error {
	if len(params) == 0 {
		return ErrParamTooFew
	}
	for inx, param := range params {
		switch param.(type) {
		case string:
			if param == "_NAME_" {
				if v, ok := params[inx+1].(string); ok {
					o.name = v
				} else {
					return ErrParamInvalid //| ErrParamTooFew
				}
			}
		default:
		}
	}
	return nil
}

func (o *object) Get(params ...interface{}) interface{} {
	if len(params) == 0 {
		return ErrParamNil
	}
	for _, param := range params {
		switch param.(type) {
		case string:
			switch param {
			case "_THIS_":
				return &o
			case "_NAME_":
				return o.name
			case "_TYPE_":
				return fmt.Sprintf("%T", o)
			case "_PRT_":
				return o
			}
		default:

		}
	}
	return nil
}

func (o *object) Super(super *Object) { o.parent = super }

func (o *object) Flush() {}

func (o *object) AddEvent() {}

func (o *object) RemoveEvent() {}

func (o *object) SeekEvent() {}
