package gooop

import ()

type Flusher interface {
	Flush()
}

type Type struct{}

type Node struct{}

type trigger struct{}

type MetaTypeKind int

type MetaTyper interface{}

type MetaType struct {
	kind map[string]interface{}
	info map[int]string
	flag int
}
