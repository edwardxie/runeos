package cmd

import (
	"fmt"
	// "os"
	"runeos/cs"
	"runeos/ide"
)

var cmdWebServer = &Command{
	UsageLine: "webserver [port]",
	Short:     "start webserver for client connect.",
	Long: `
start webserver for client connect,

default port 8080, set [port].
`,
}

func init() {
	cmdWebServer.Run = createWeb
}

func createWeb(cmd *Command, args []string) {
	print("Init ide env...")
	ide.InitIDE()
	var port string
	// fmt.Printf("Args: %#v", args)
	if len(args) == 0 {
		port = ":8080"
	} else {
		port = args[0]
	}
	println("[done]")
	fmt.Printf("Start web server...port [%v]\n", port)
	if err := cs.Server("web", port); err != nil {
		fmt.Printf("[ERROR]%s\n", err)
	}
}
