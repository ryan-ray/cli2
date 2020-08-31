package cli2

import (
	"errors"
	"flag"
	"testing"
)

type TestCmd struct {
	executionSuccess bool
	args             []string
}

func (tc TestCmd) Name() string              { return "TestCmd" }
func (tc TestCmd) Description() string       { return "TestDescription" }
func (tc TestCmd) Register(fs *flag.FlagSet) {}
func (tc *TestCmd) Execute(args []string) error {
	tc.executionSuccess = true
	tc.args = args
	return nil
}

var errCmdReturn error = errors.New("ErrCmdRet")

type ErrorCmd struct{}

func (ec ErrorCmd) Name() string              { return "ErrorCmd" }
func (ec ErrorCmd) Description() string       { return "ErrorDescription" }
func (ec ErrorCmd) Register(fs *flag.FlagSet) {}
func (ec ErrorCmd) Execute(args []string) error {
	return errCmdReturn
}

func TestNewApp(t *testing.T) {
	tc := &TestCmd{false, []string{}}
	app := NewApp(tc)

	tests := []struct {
		name string
		got  string
		want string
	}{
		{name: "CmdName", got: app.Command.Name(), want: "TestCmd"},
		{name: "CmdDescription", got: app.Command.Description(), want: "TestDescription"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("Got %s, want %s", tt.got, tt.want)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name string
		cmd  *TestCmd
		args []string
		err  error
	}{
		{name: "EmptyArgs", cmd: &TestCmd{}, args: []string{}},
		{name: "SingleArg", cmd: &TestCmd{}, args: []string{"arg1"}},
		{name: "MultipleArgs", cmd: &TestCmd{}, args: []string{"arg1", "arg2", "arg3"}},
		{name: "DuplicateArgs", cmd: &TestCmd{}, args: []string{"arg", "arg", "arg"}},
		{name: "ArgAfterFlag", cmd: &TestCmd{}, args: []string{"-f", "arg"}},
		{name: "OnlyFlags", cmd: &TestCmd{}, args: []string{"-a", "-b"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := NewApp(tt.cmd)
			app.Command.Execute(tt.args)
			if !tt.cmd.executionSuccess {
				t.Errorf("Command did not successfully execute")
			}

			for i, arg := range tt.args {
				if arg != tt.args[i] {
					t.Errorf("Got %s, want %s", arg, tt.args[i])
				}
			}
		})
	}
}
