package cli2

type Command interface {
	Name() string
	Description() string
	Execute(args []string) error
}

type Node struct {
	Command     Command
	SubCommands map[string]*Node
}

func (n *Node) AddSub(name string, c Command) *Node {
	sc := &Node{
		Command:     c,
		SubCommands: make(map[string]*Node),
	}

	n.SubCommands[name] = sc

	return sc
}

type App struct {
	Node
}

func NewApp(c Command) *App {
	return &App{
		Node{
			Command:     c,
			SubCommands: make(map[string]*Node),
		},
	}
}

func (a App) Run(args []string) error {
	var cursor *Node = &a.Node
	i := 0

	for _, arg := range args {
		if sc, ok := cursor.SubCommands[arg]; ok {
			cursor = sc
			i += 1
		}
	}

	return cursor.Command.Execute(args[i:])
}
