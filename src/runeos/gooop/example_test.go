package gooop_test

import (
	"fmt"

	"runeos/gooop"
)

func Example_NewObject() {
	obj := gooop.NewObject("Example_Object")
	fmt.Printf("%v", obj.Get("_NAME_"))

	// Output:
	// Example_Object
}

func Example_NewNilObject() {
	obj := gooop.NewObject()
	fmt.Printf("%v", obj.Get("_NAME_"))

	// Output:
	//
}
