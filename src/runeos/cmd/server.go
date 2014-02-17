package cmd

import (
	// "flag"
	"fmt"
	"os"
)

var (
	cmdServer = &Command{
		UsageLine: "server [port]",
		Short:     "start server for client connect.",
		Long: `
start server for client connect,

default port 8080, set [port].
`,
		Pid: os.Getpid(),
	}
)

func init() {
	cmdServer.Run = registry_command_server
}

func registry_command_server(cmd *Command, args []string) {
	switch {
	case len(args) > 0:
		fmt.Printf("\n[调试]Exec backend server application -> %v is run? -> %v\n", args[0], cmd.Runnable())
	default:
		fmt.Printf("\n[调试]Exec backend server application is run? -> %v\n", cmd.Runnable())
	}

}
