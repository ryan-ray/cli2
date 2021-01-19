package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ryan-ray/cli2"
)

type BasicCommand struct {
	name        string
	description string
}

func (cmd BasicCommand) Name() string {
	return cmd.name
}

func (cmd BasicCommand) Description() string {
	return cmd.description
}

func (cmd BasicCommand) DefineFlags(fs *flag.FlagSet) {}

func (cmd BasicCommand) Execute(args []string) error {
	fmt.Println("Hello from our Basic Command!")
	return nil
}

func main() {

	b := &BasicCommand{
		name:        "Basic Command",
		description: "Description of our basic command",
	}

	app := cli2.NewApp(b)

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
