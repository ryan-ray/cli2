package main

import (
	"fmt"
	"os"

	"github.com/ryan-ray/cli/cli2"
)

type Main struct{}

func (m Main) Name() string        { return "Cli2" }
func (m Main) Description() string { return "Description for Cli2" }

func (m Main) Execute(args []string) error {
	fmt.Fprintf(os.Stdout, "Executing %s with args %v\n", m.Name(), args)
	return nil
}

type Foo struct{}

func (b Foo) Name() string        { return "foo" }
func (b Foo) Description() string { return "Runs foo" }

func (b Foo) Execute(args []string) error {
	fmt.Fprintf(os.Stdout, "Executing %s with args %v\n", b.Name(), args)
	return nil
}

type Bar struct{}

func (f Bar) Name() string        { return "bar" }
func (f Bar) Description() string { return "Runs bar" }

func (f Bar) Execute(args []string) error {
	fmt.Fprintf(os.Stdout, "Executing %s with args %v\n", f.Name(), args)
	return nil
}

type Baz struct{}

func (f Baz) Name() string        { return "baz" }
func (f Baz) Description() string { return "Runs baz" }

func (f Baz) Execute(args []string) error {
	fmt.Fprintf(os.Stdout, "Executing %s with args %v\n", f.Name(), args)
	return nil
}

func main() {

	app := cli2.NewApp(&Main{})

	foo := app.AddSub("foo", &Foo{})
	foo.AddSub("baz", &Baz{})

	app.AddSub("bar", &Bar{})

	app.Run(os.Args[1:])

}
