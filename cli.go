package cli2

import (
	"flag"
)

type Command interface {
	Name() string
	Description() string
	Register(*flag.FlagSet)
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

func (a App) Root() *Node {
	return &a.Node
}

func (a App) Run(args []string) error {
	var cursor *Node = a.Root()
	// Register our global/root level flags
	cursor.Command.Register(a.fs)
	i := 0

	for _, arg := range args {
		if sc, ok := cursor.SubCommands[arg]; ok {
			cursor = sc
			i += 1
		} else {
			break
		}
	}

	// Register sub command flags
	if i > 0 {
		cursor.Command.Register(a.fs)
	}

	a.fs.Parse(args[i:])
	return cursor.Command.Execute(a.fs.Args())
}
