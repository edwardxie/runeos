package ide

import (
// "fmt"
)

type Value struct {
	v interface{}
}

func (v Value) String() string {
	return v.v.(string)
}
