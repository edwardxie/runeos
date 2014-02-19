package cmd

import (
	"fmt"
	"os"
)

var (
	runeos_root string

	flag_config_path_file = cmdConfig.Flag.String("path-file", "conf/runeos.json", cmdConfig.Long)
	cmdConfig             = &Command{
		UsageLine: "config [-p=conf]",
		Short:     "Handle IDE config folder and files manager.",
		Long: `
Application handle IDE config folder and files manager.
`,
		Pid: os.Getpid(),
	}
)

func init() { cmdConfig.Run = registry_command_config }

func registry_command_config(cmd *Command, args []string) {
	err := init_config()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("Args: %#v\n", args)
	fmt.Printf("Config command. path => %#v\n", *flag_config_path_file)
}

func init_config() error {
	runeos_root = os.Getenv("RUNEOS_ROOT")
	if len(runeos_root) == 0 {
		return fmt.Errorf("[ERROR]Environment variable RUNEOS_ROOT no set.")
	}
	if _, err := os.Stat(*flag_config_path_file); err != nil {
		return err
	}
	return nil
}
