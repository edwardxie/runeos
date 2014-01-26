package golua

import (
	"fmt"
	"github.com/aarzilli/golua/lua"
)

var L = lua.NewState()

func init() {
	L.OpenLibs()
}

func PrintInfo() {
	fmt.Print("Lua version: ")
	L.DoString("print(_VERSION)")
}

func Close() {
	L.Close()
}
