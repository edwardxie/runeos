package gooop

import (
	// "container/list"
	"fmt"
	// "reflect"
	// "runeos/utils"
)

type Object interface {
	// This() Object
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

type Geter interface {
	Get(...interface{}) interface{}
}

type Seter interface {
	Set(...interface{}) error
}

type object struct {
	name   string
	flag   *Flag
	parent *Object
	lnext  *Object
	rnext  *Object
	member []*Variant
}

func NewObject(params ...interface{}) Object {
	isName := false
	obj := object{member: make([]*Variant, 0)}
	if len(params) == 0 {
		return Object(&obj)
	}
	for _, param := range params {
		switch param.(type) {
		case string:
			if isName {
				// return Object(ErrParamTooMany)
				obj.member = append(obj.member, newVariant(param))
				break
			}
			obj.name = param.(string)
			isName = true
		default:
			obj.member = append(obj.member, newVariant(param))
			// fmt.Printf("Params for switch is %#v [index:type] => [%v:%T]\n ", params, inx, typ)
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
			case "_MEMBER_":
				return o.member
			}
		default:

		}
	}
	return nil
}

func (o *object) Super(super *Object) { o.parent = super }

func (o *object) List() []*Variant { return o.member }

// func (o *object) Member() ([]reflect.Value, error) {
// 	i := 0
// 	var in []reflect.Value
// 	var f interface{}
// 	for k, v := range o.member {
// 		var in = make([]reflect.Value, k.NumIn())
// 		in[k] = v
// 		f = k
// 	}
// 	return reflect.ValueOf(f).Call(in), nil
// }

func (o *object) Map(key string) {}

func (o *object) Copy(dest *Object) {
	var tmp Object
	tmp = Object(o)
	dest = &tmp
}
