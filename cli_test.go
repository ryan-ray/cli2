package cli2

import (
	"flag"
	"fmt"
	"testing"
)

type MockCmd struct {
	vars        map[string]interface{}
	name        func() string
	description func() string
	defineFlags func(*flag.FlagSet)
	execute     func([]string) error
}

func (mc MockCmd) Name() string                 { return mc.name() }
func (mc MockCmd) Description() string          { return mc.description() }
func (mc MockCmd) DefineFlags(fs *flag.FlagSet) { mc.defineFlags(fs) }
func (mc MockCmd) Execute(args []string) error  { return mc.execute(args) }

var testCmd MockCmd = MockCmd{
	vars:        make(map[string]interface{}),
	name:        func() string { return "TestCmd" },
	description: func() string { return "TestCmd Description" },
	defineFlags: func(*flag.FlagSet) {},
	execute:     func(args []string) error { return nil },
}

func TestNewApp(t *testing.T) {
	app := NewApp(&testCmd)

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

func TestExecuteCmd(t *testing.T) {
	testCmd.execute = func(args []string) error {
		testCmd.vars["args"] = args
		testCmd.vars["executed"] = true

		return nil
	}

	app := NewApp(&testCmd)
	err := app.Run([]string{"./prog", "one", "two", "three"})
	if err != nil {
		t.Errorf("Error returned from execute")
	}

	v, ok := testCmd.vars["executed"].(bool)
	if !ok {
		t.Errorf("Error with bool")
	}

	if v != true {
		t.Errorf("Cmd did not execute")
	}
}

func TestExecuteSubCommand(t *testing.T) {

	mainCmd := MockCmd{
		name:        func() string { return "Main" },
		description: func() string { return "MainDescription" },
		defineFlags: func(fs *flag.FlagSet) {},
		execute:     func(args []string) error { return nil },
	}

	subCmd := MockCmd{
		name:        func() string { return "Sub" },
		description: func() string { return "SubDescription" },
		defineFlags: func(fs *flag.FlagSet) {},
		execute:     func(args []string) error { return fmt.Errorf("SubCmd executed") },
	}

	app := NewApp(&mainCmd)
	app.AddSub("sub", &subCmd)

	err := app.Run([]string{"./test", "sub"})
	if err == nil {
		t.Errorf("Sub command not executed")
	}
}

func TestRemoveDisallowedChars(t *testing.T) {
	tests := []struct {
		old  string
		want string
	}{
		{"cmdwithop%+-=/", "cmdwithop"},
		{" leadingtrailing ", "leadingtrailing"},
		{"h@ck th3 pl@n3t", "h@ckth3pl@n3t"},
		{"mess with the best, die like the rest", "messwiththebestdieliketherest"},
	}

	for _, tt := range tests {
		t.Run(tt.old, func(t *testing.T) {
			got := removeDisallowedChars(tt.old, " \t\n\r+-=/*%,")
			if got != tt.want {
				t.Errorf("Got %s, want %s", got, tt.want)
			}
		})
	}
}
