package cmd

import (
	"flag"
	"fmt"
	"os"
	"runeos/cs"
	"runeos/ide"
)

var (
	flagHttpAddr = flag.String("http", ":8080", cmdWebServer.Long)

	cmdWebServer = &Command{
		UsageLine: "webserver [port]",
		Short:     "start webserver for client connect.",
		Long: `
start webserver for client connect,

default port 8080, set [port].
`,
		Pid: os.Getpid(),
	}
)

func init() {
	cmdWebServer.Run = createWebserver
}

func createWebserver(cmd *Command, args []string) {
	fmt.Printf("Init ide env...[PID:%v]", cmdWebServer.Pid)
	ide.InitIDE()
	var port string
	if len(args) == 0 {
		port = ":8080"
	} else {
		port = args[0]
	}
	println("[done]")
	fmt.Printf("Start web server... [%v]\n", port)
	if err := cs.Server("web", port); err != nil {
		fmt.Printf("[ERROR]%s\n", err)
	}
}
