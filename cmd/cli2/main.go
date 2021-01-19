package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ryan-ray/cli2"
)

type Main struct {
	debug   bool
	verbose bool
}

func (m Main) Name() string { return "Cli2" }
func (m Main) Description() string {
	return `This is the description text for the root or main command node.

Available commands:

sub	Runs the sub command
`
}

func (m *Main) DefineFlags(fs *flag.FlagSet) {
	fs.BoolVar(&m.debug, "debug", false, "Set debug")
	fs.BoolVar(&m.verbose, "v", false, "Set verbose")
}

func (m Main) Execute(args []string) error {
	fmt.Fprintf(os.Stdout, "Executing %s with args %v\n", m.Name(), args)
	fmt.Fprintf(os.Stdout, "Debug: %t\n", m.debug)
	fmt.Fprintf(os.Stdout, "Verbose: %t\n", m.verbose)
	return nil
}

type Sub struct {
	filename string
}

func (s Sub) Name() string        { return "Sub Command" }
func (s Sub) Description() string { return "Sub Command for Main" }

func (s *Sub) DefineFlags(fs *flag.FlagSet) {
	fs.StringVar(&s.filename, "subname", "", "Set subname")
}

func (s Sub) Execute(args []string) error {
	fmt.Fprintf(os.Stdout, "Executing %s with args %v\n", s.Name(), args)
	return nil
}

func main() {

	app := cli2.NewApp(&Main{})
	app.AddSub("sub", &Sub{})
	app.Run(os.Args)

}
