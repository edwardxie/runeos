package core

import (
	// "container/list"
	"fmt"
	// "reflect"
	// "runeos/utils"
)

func (o *object) method(params ...interface{}) interface{} {
	return nil
}

func NewObject(params ...interface{}) Object {
	isName := false
	obj := object{} //oid: utils.NewID().UUID()
	if len(params) == 0 {
		return Object(&obj)
	}
	for inx, param := range params {
		switch typ := param.(type) {
		case string:
			if isName {
				panic("Too many string params set object name.")
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
		err := fmt.Errorf("Param is nil, I not know you need what.")
		return err
	}
	for inx, param := range params {
		switch param.(type) {
		case string:
			if param == "_NAME_" {
				if v, ok := params[inx+1].(string); ok {
					o.name = v
				} else {
					return fmt.Errorf("the second param not be string or too few param.")
				}
			}
		default:
		}
	}
	return nil
}

func (o *object) Get(params ...interface{}) interface{} {
	if len(params) == 0 {
		err := fmt.Errorf("Param is nil, I not knowe you need what.")
		return err
	}
	for _, param := range params {
		switch param.(type) {
		case string:
			switch param {
			case "_ID_":
				return fmt.Sprintf("%v", o)
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
