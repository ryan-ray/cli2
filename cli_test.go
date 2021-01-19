package cli2

import (
	"errors"
	"flag"
	"fmt"
	"testing"
)

var (
	name        = "TestCmd"
	description = "TestDescription"
)

type TestCmd struct {
	executionSuccess bool
	args             []string
}

func (tc TestCmd) Name() string                 { return "TestCmd" }
func (tc TestCmd) Description() string          { return "TestDescription" }
func (tc TestCmd) DefineFlags(fs *flag.FlagSet) {}
func (tc *TestCmd) Execute(args []string) error {
	tc.executionSuccess = true
	tc.args = args
	return nil
}

var errCmdReturn error = errors.New("ErrCmdRet")

type ErrorCmd struct{}

func (ec ErrorCmd) Name() string                 { return "ErrorCmd" }
func (ec ErrorCmd) Description() string          { return "ErrorDescription" }
func (ec ErrorCmd) DefineFlags(fs *flag.FlagSet) {}
func (ec ErrorCmd) Execute(args []string) error {
	return errCmdReturn
}

func TestNewApp(t *testing.T) {
	tc := &TestCmd{false, []string{}}
	app := NewApp(tc)

	got := app.Command.Name()
	want := "TestCmd"
	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}

func TestAppRootUnassigned(t *testing.T) {
	app := &App{}
	n, err := app.Root()

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	if n != nil {
		t.Errorf("Got %v, want nil", n)
	}
}

func TestExecute(t *testing.T) {
	tc := &TestCmd{false, []string{}}
	app := NewApp(tc)

	args := []string{"one", "two", "three"}
	app.Command.Execute(args)

	if !tc.executionSuccess {
		t.Errorf("Expected true, got false")
	}

	for i, v := range tc.args {
		if v != args[i] {
			t.Errorf("Got %s, want %s", v, args[i])
		}
	}
}

type MainCmd struct{}

func (c MainCmd) Name() string                 { return "Main" }
func (c MainCmd) Description() string          { return "MainDescription" }
func (c MainCmd) DefineFlags(fs *flag.FlagSet) {}
func (c MainCmd) Execute(args []string) error  { return nil }

type SubCmd struct{}

func (c SubCmd) Name() string                 { return "Sub" }
func (c SubCmd) Description() string          { return "SubDescription" }
func (c SubCmd) DefineFlags(fs *flag.FlagSet) {}
func (c SubCmd) Execute(args []string) error {
	return fmt.Errorf("SubCmd executed")
}

func TestExecuteSubCommand(t *testing.T) {
	app := NewApp(&MainCmd{})
	app.AddSub("sub", &SubCmd{})

	err := app.Run([]string{"./test", "sub"})
	if err == nil {
		t.Errorf("Sub command not executed")
	}
}
