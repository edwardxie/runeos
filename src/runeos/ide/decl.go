package ide

// Environment type declaration
type Enver interface{}

type Config struct {
	Setting
	// Plugins
}

type Setting struct {
	Project
	IDE
}

type IDE struct {
	Appearance
	Editor  `json:"Editor"`
	KeyMaps `json:"KeyMaps"`
}

type Env struct {
	Os       string `json:"Os"`
	Language string `json:"Language"`
	Arch     string `json:"Arch"`

	// Sysenv []string
	Sdk
	Goenv
	Cenv
	Cppenv
	Luaenv
	Javaenv
	Nodejsenv
	Pythonenv
	Rubyenv
	Bashenv
}

type Appearance struct{}

type Sdk struct{}

type Goenv struct {
	Goroot string `json:"GOROOT",xml:"GOROOT`
	Gopath string `json:"GOPATH"`
}

type Cenv struct{}

type Cppenv struct{}

type Luaenv struct{}

type Javaenv struct{}

type Nodejsenv struct{}

type Pythonenv struct{}

type Rubyenv struct{}

type Bashenv struct{}

type Webenv struct {
	Address string
	Port    string
}

// Workspace type declaration
type Workspace struct {
	Env      `json:"Env"`
	Setting  `json:"Setting"`
	Projects []Project `json:"Projects"`
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

// Project type declaration
type Project struct {
	Name     string
	Type     string
	Path     string
	FileList string
	FileType string
	LibList  string
}
