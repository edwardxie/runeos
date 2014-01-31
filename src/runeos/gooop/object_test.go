package gooop

import (
	"reflect"
	"testing"
)

type SpecialString interface{}

type SpecialObject interface{}

type TestStruct struct {
	Dep1 string        `inject`
	Dep2 SpecialString `inject`
	Dep3 string
}

var (
	testdata = [...]interface{}{"NewObject", "Too", "many",
		112, true, []byte("Hello,world")}
)

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func equal(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) {
		t.Errorf("Did not expect DeepEqual %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func notequal(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Did not expect Not DeepEqual %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

/* Test Items */
func Test_NewObject(t *testing.T) {
	obj := NewObject("NewObject", "Too", "many",
		112, true, []byte("Hello,world"))
	refute(t, obj, new(Object))
	//Param len=0 is nil and name is "", create no name object.
	refute(t, NewObject(), new(Object))
	//Too many string is error
	refute(t, NewObject("Too", "many"), new(Object)) //ErrParamTooMany
	t.Logf("\nName: \t\t%v\nAddress: \t%v\nType: \t\t%v\nPRT: \t\t%#v\n", obj.Get("_NAME_"), obj.Get("_THIS_"), obj.Get("_TYPE_"), obj.Get("_MEMBER_"))
	// t.Logf("\nName: \t\t%v\nAddress: \t%v\nType: \t\t%v\nPRT: \t\t%#v\n", obj.Get("_NAME_"), obj.Get("_THIS_"), obj.Get("_TYPE_"), obj.Get("_MEMBER_").([]*Variant)[4].data[reflect.TypeOf([]byte("d"))].Bytes()) //[3].data[reflect.TypeOf(true)].Bool()
}

func Test_SuperObject(t *testing.T) {
	obj := NewObject("NewObject")
	sub := NewObject("subObject")
	sub.Super(&obj)
}

func Test_Methoder(t *testing.T) {
	obj := NewObject("NewObject")
	obj.Set("testData1")
	obj.Set("testData2")
	obj.Set("testData3")
	obj.Set("testData1", "testData1")
	// if err := obj.Set("_NAME_", "testData1"); err != ErrNo {
	// 	t.Error(err)
	// }
	// if err := obj.Set("_NAME_", 12); err != ErrParamTooFew {
	// 	t.Error(err)
	// }
	// obj.Set(sub)
	// sub.Flush()
	// t.Logf("Master: %#v\nSubject: %#v\n", obj, sub)
	// t.Logf("GetValue: %#v\n", obj.Get())
	// t.Logf("GetValue: %#v\nSubObject: %#v\n", obj.Get(),
	// obj.Get().([]interface{})[3].(Object).Get())
}
