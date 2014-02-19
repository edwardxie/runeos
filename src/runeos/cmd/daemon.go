package cmd

import (
	// "flag"
	"fmt"
	"os"
)

var (
	cmdDaemon = &Command{
		UsageLine: "daemon",
		Short:     "Application daemon process.",
		Long: `
Application daemon process, including hot config and server operation.
`,
		Pid: os.Getpid(),
	}
)

func init() {
	cmdDaemon.Run = registry_command_daemon
}

func registry_command_daemon(cmd *Command, args []string) {
	fmt.Printf("Daemon args: %v\n", args)
	switch {
	case len(args) > 0:
		fmt.Printf("\n[调试]Exec application daemon process -> %v is run? -> %v\n", args[0], cmd.Runnable())
	default:
		fmt.Printf("\n[调试]Exec application daemon process, is run? -> %v\nConfig path: %v\n", cmd.Runnable(), *flag_config_path_file)
	}

}
