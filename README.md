# cli2
![Build and Test](https://github.com/ryan-ray/cli2/workflows/Build%20and%20Test/badge.svg)

cli2 is a _very_ light-weight package to help with bootstrapping cli based applications. Particularly those that have multiple, chained commands like ```app sub do cmd -f -g -h```. There are a million of these already, but they are quite complex and can be a little over engineered<sup>1</sup> for what amounts to a very simple problem. I've designed this package to so that it resembles, rightly or wrongly, the way that I typically organize my cli applications. I've also tried to make sure it won't get in the way of how you write go code, and as far as I can see, doesn't try and shoehorn in an OOP style framework where it doesn't belong.

_<sup>1. I glanced at a couple of packages, nowhere near in depth enough to make this statement. But honest to god, one of them used a radix tree (whereas this uses a much more intuitive multiple linked list :neutral_face: )</sup>_

### Why is this called 'cli2'?

I threw together a few roughed out packages that all tackled this same problem, swinging between various degrees of simplicity and complexity. The second one, sitting in a folder called 'cli2', was the one that I prefered and honestly trying to come up with cool project names is exhausting. I'm in my mid thirties, and I don't have time for that.

### Tutorial

Every "command", including the root command (i.e., the psuedo main command) is a struct that implements the ```cli2.Command``` interface, like so;

```
type Main struct {
  debug bool
}
func (m Main) Name() string { return "app" }
func (m Main) Description() string { return "Description of execution. Displayed when -h/-help is flagged" }
func (m *Main) Register(fs *flag.FlagSet) {
  fs.BoolVar(&m.debug, "debug", false, "Set the app to debug mode")
}

func (m Main) Execute(args []string) error {
  ...
}
```

Then kick everything off with a ```cli2.App``` struct;

```
app := cli2.NewApp(&Main{})
if err := app.Run(os.Args); err != nil {
  ...
}
```
