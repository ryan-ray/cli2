# cli2
![Build and Test](https://github.com/ryan-ray/cli2/workflows/Build%20and%20Test/badge.svg)

cli2 is a _very_ light-weight package to help with bootstrapping cli based applications. Particularly those that have multiple, chained commands like ```app sub do cmd -f -g -h```. There are a million of these already, but there are quite complex and can be a little over engineered for what amounts to a very simple problem.

### Tutorial

Every "command", including the root command (i.e., the psuedo main command) is a struct that implements the ```Command``` interface.

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
