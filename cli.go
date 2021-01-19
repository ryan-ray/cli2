package cli2

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type Command interface {
	Name() string
	Description() string
	DefineFlags(*flag.FlagSet)
	Execute([]string) error
}

type Node struct {
	Command     Command
	SubCommands map[string]*Node
}

func NewNode(c Command) *Node {
	return &Node{
		Command:     c,
		SubCommands: make(map[string]*Node),
	}
}

func (n *Node) AddSub(name string, c Command) *Node {
	sc := NewNode(c)
	// TODO: Sanitize string to remove spaces
	n.SubCommands[name] = sc

	return sc
}

type App struct {
	Node
	fs *flag.FlagSet
}

func NewApp(c Command) *App {
	n := NewNode(c)
	return &App{
		Node: *n,
		fs:   flag.NewFlagSet(c.Name(), flag.ExitOnError),
	}
}

func (a App) Root() (*Node, error) {
	if a.Node.Command == nil {
		return nil, errors.New("No root node assigned. Did you use NewApp?")
	}
	return &a.Node, nil
}

// Run looks through our
func (a App) Run(args []string) error {
	// Ditch our first arg, root node is always
	// selected by default
	args = args[1:]

	cursor, err := a.Root()
	if err != nil {
		return err
	}

	// DefineFlags our global/root level flags
	cursor.Command.DefineFlags(a.fs)
	i := 0

	cmdChain := []string{}
	cmdChain = append(cmdChain, cursor.Command.Name())

	for _, arg := range args {
		if sc, ok := cursor.SubCommands[arg]; ok {
			cursor = sc
			cmdChain = append(cmdChain, cursor.Command.Name())
			i += 1
		} else {
			break
		}
	}

	// DefineFlags sub command flags if we have a sub command.
	if i > 0 {
		cursor.Command.DefineFlags(a.fs)
	}

	a.fs.Usage = usage(
		strings.Join(cmdChain, " "),
		cursor.Command.Description(),
		a.fs,
	)

	a.fs.Parse(args[i:])

	return cursor.Command.Execute(a.fs.Args())
}

func usage(name string, description string, fs *flag.FlagSet) func() {
	// Sneak our description in there
	return func() {
		w := fs.Output()
		fmt.Fprintf(w, "Usage of %s:\n\n", name)
		fmt.Fprintf(w, "%s\n\n", description)
		fs.PrintDefaults()
	}
}
