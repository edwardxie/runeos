package core

import (
	"testing"
)

func Test_New(t *testing.T) {
	obj := NewObject("NewObject")
	tmp := NewObject("TempObject")
	tmp.Set(1)
	tmp.Set(2)
	tmp.Set(3)
	obj.Set("testData1")
	obj.Set("testData2")
	obj.Set("testData3")
	obj.Set(tmp)
	// tmp.Flush()
	t.Logf("Object: %#v\n", obj)
	// t.Logf("GetValue: %#v\n", obj.Get())
	// t.Logf("GetValue: %#v\nSubObject: %#v\n", obj.Get(),
	// obj.Get().([]interface{})[3].(Object).Get())
}
