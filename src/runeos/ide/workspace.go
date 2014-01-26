package ide

import (

// "fmt"
// "os"
// "strings"
)

type Workspace struct {
	Env      `json:"Env"`
	Setting  `json:"Setting"`
	Projects []Project `json:"Projects"`
}

type Setting struct {
	Editor  `json:"Editor"`
	KeyMaps `json:"KeyMaps"`
}

type Editor struct {
	FontFace    string
	FontSize    int
	TabSize     int
	Theme       string
	ColorScheme string

	LineNumber bool
}

type KeyMaps struct{}

func NewWorkspace() *Workspace { return &Workspace{} }
