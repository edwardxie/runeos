package cmd

import (
	"fmt"
	"os"
	"runeos/cs"
	"runeos/ide"
)

var (
	flag_http_addr = cmdWebide.Flag.String("http", ":8080", cmdWebide.Long)

	cmdWebide = &Command{
		UsageLine: "webide [-http=:port]",
		Short:     "start web ide server for client connect.",
		Long: `
start web ide server for client connect,

default port 8080, set -http=[:port].
`,
		Pid: os.Getpid(),
	}
)

func init() {
	cmdWebide.Run = registry_command_ide
}

func registry_command_ide(cmd *Command, args []string) {
	fmt.Printf("Init ide env...[PID:%v]", cmdWebide.Pid)
	ide.InitIDE()
	println("[done]")
	fmt.Printf("Start web ide server... [%v]\n", *flag_http_addr)
	if err := cs.Server("webide", *flag_http_addr); err != nil {
		fmt.Printf("[ERROR]%s\n", err)
	}
}
