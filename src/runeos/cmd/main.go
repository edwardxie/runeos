package cmd

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	// "log"
	"os"
	"strings"
)

type Command struct {
	Run         func(cmd *Command, args []string)
	UsageLine   string
	Short       string
	Long        string
	Flag        flag.FlagSet
	CustomFlags bool
	Pid         int
}

func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}

func (c *Command) Runnable() bool {
	return c.Run != nil
}

var commands = []*Command{
	cmdNew,
	cmdWebServer,
	// cmdHelp,
}

func Main() {
	flag.Usage = usage
	flag.Parse()
	// log.SetFlags(0)

	args := flag.Args()
	// if len(args) < 1 {
	// 	usage()
	// }

	// if args[0] == "help" {
	// 	help(args[1:])
	// 	return
	// }

	httpMode := *flagHttpAddr != ""
	fmt.Printf("Args: %v\nFlag: %v httpmode: %v\n", args, *flagHttpAddr, httpMode)
	// if httpMode {
	// 	Command.Run(commands[1], args)
	// 	os.Exit(2)
	// 	// return
	// }
	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			os.Exit(2)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "runeos: unknown subcommand %q\nRun 'runeos help' for usage.\n", args[0])
	os.Exit(2)
}

var usageTemplate = `Runeos is a tool for ide.

Usage:

	runeos command [arguments]

The commands are:
{{range .}}{{if .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "runeos help [command]" for more information about a command.

Additional help topics:
{{range .}}{{if not .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "runeos help [topic]" for more information about that topic.

`
var helpTemplate = `{{if .Runnable}}usage: runeos {{.UsageLine}}

{{end}}{{.Long | trim}}
`

func usage() {
	tmpl(os.Stdout, usageTemplate, commands)
	os.Exit(2)
}

func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func help(args []string) {
	if len(args) == 0 {
		usage()
		// not exit 2: succeeded at 'go help'.
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stdout, "usage: runeos help command\n\nToo many arguments given.\n")
		os.Exit(2) // failed at 'runeos help'
	}

	arg := args[0]

	for _, cmd := range commands {
		if cmd.Name() == arg {
			tmpl(os.Stdout, helpTemplate, cmd)
			// not exit 2: succeeded at 'go help cmd'.
			return
		}
	}

	fmt.Fprintf(os.Stdout, "Unknown help topic %#q.  Run 'runeos help'.\n", arg)
	os.Exit(2) // failed at 'runeos help cmd'
}
