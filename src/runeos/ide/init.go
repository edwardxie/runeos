package ide

import (
	// "fmt"
	"os"
	"runtime"
)

const ()

var (
	globalEnv map[string]interface{}
	ws        *Workspace
	rfs       []string
)

func init() {
	globalEnv = make(map[string]interface{})
	ws = &Workspace{}
	rfs = []string{"bin", "conf", "doc", "lib", "locale", "plugins", "scripts", "static", "template"}
}

func InitIDE() {
	if globalEnv["RUNEOS_ROOT"] = os.Getenv("RUNEOS_ROOT"); globalEnv["RUNEOS_ROOT"] == "" {
		println("[ERROR]Environment variable RUNEOS_ROOT no set.")
		os.Exit(1)
	}
	globalEnv["runpath"], _ = os.Getwd()
	confpath = globalEnv["RUNEOS_ROOT"].(string) + "/" + rfs[1]
	globalEnv["confpath"] = confpath
	ws.Os = runtime.GOOS
	ws.Arch = runtime.GOARCH
	ws.Goroot = runtime.GOROOT()
	ws.Gopath = os.Getenv("GOPATH")
	ws.LineNumber = true
	globalEnv["workspace"] = ws
	// ws.Sysenv = os.Environ()
}

func loadDefault() {}

func reload() {}

// func GetRuneosRoot() string {
// 	return env["RUNEOS_ROOT"]
// }

// func GetWorkspace() *Workspace {
// return env["workspace"]
// }
