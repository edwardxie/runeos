package ide

import ()

var (
	globalEnv map[string]interface{}
)

type Enver interface{}

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

// type Sysenv struct {
// 	Sysenv []string
// }

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

func SetEnv() {

}

func GetEnv(n string) interface{} {
	return globalEnv[n]
}

func GetEnvMap() map[string]interface{} {
	return globalEnv
}
