package gooop

import (
	"reflect"
)

const (
	FlagCreateKind = iota
	FlagAppendKind
	FlagModifyKind
	FlagUpdateKind
	FlagChangeKind
	FlagErrorKind
	FlagFatalkind
	FlagDumpKind
	FlagCloseKind
	FlagRemoveKind
	FlagDestoryKind
)

type Flag struct {
	kind   int
	info   string
	invoke func(interface{}) ([]reflect.Value, error)
}

var (
	FlagCreate  = &Flag{FlagCreateKind, "Create", nil}
	FlagAppend  = &Flag{FlagAppendKind, "Append", nil}
	FlagModify  = &Flag{FlagModifyKind, "Modify", nil}
	FlagUpdate  = &Flag{FlagUpdateKind, "Update", nil}
	FlagChange  = &Flag{FlagChangeKind, "Change", nil}
	FlagDump    = &Flag{FlagDumpKind, "Dump", nil}
	FlagClose   = &Flag{FlagCloseKind, "Close", nil}
	FlagRemove  = &Flag{FlagRemoveKind, "Remove", nil}
	FlagDestory = &Flag{FlagDestoryKind, "Destory", nil}
)

func (f *Flag) String() string {
	return f.info
}
