package gooop

import (
	"reflect"
	"testing"
)

type SpecialString interface{}

type TestStruct struct {
	Dep1 string        `inject`
	Dep2 SpecialString `inject`
	Dep3 string
}

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

func Test_NewObject(t *testing.T) {
	obj := NewObject("NewObject")
	tmp := NewObject("TempObject")
	//NewObject() //Prop if len=0 is nil
	// NewObject("Too", "many") //panic("Too many string set object name.")
	tmp.Set(1)
	tmp.Set(2)
	tmp.Set(3)
	tmp.Super(&obj)
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
	obj.Set(tmp)
	// tmp.Flush()
	// t.Logf("Master: %#v\nSubject: %#v\n", obj, tmp)
	// t.Logf("GetValue: %#v\n", obj.Get())
	// t.Logf("GetValue: %#v\nSubObject: %#v\n", obj.Get(),
	// obj.Get().([]interface{})[3].(Object).Get())
	t.Logf("\nAddress: \t%v\nType: \t%v\nPRT: %v\n", obj.Get("_THIS_"), obj.Get("_TYPE_"), obj.Get("_PRT_"))
}
