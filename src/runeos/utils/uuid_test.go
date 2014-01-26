package utils

import (
	"testing"
)

func TestNewId(t *testing.T) {
	r := NewID()
	t.Logf("%x\n", r)
}

func TestMD5(t *testing.T) {
	r := NewRand().MD5()
	t.Logf("Len: %v\n%x\n", len(r), r)
}

func TestURL(t *testing.T) {
	r := NewRand().URL(1)
	t.Logf("Len: %v\n%s\n", len(r), r)
}

func TestUUID(t *testing.T) {
	r := NewRand().UUID()
	t.Logf("%s\n", r)
}
