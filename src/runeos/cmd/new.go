package cmd

import (
	"fmt"
	"os"
)

var cmdNew = &Command{
	UsageLine: "new [appname]",
	Short:     "create an application base on project",
	Long: `
create an application base on project,

which in the current path with folder named [appname].
`,
}

func init() {
	cmdNew.Run = createApp
}

func createApp(cmd *Command, args []string) {
	currpath, _ := os.Getwd()
	fmt.Printf("\n[调试]Create application %v\n", currpath)
}
