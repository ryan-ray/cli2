package cli2

import (
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

func (tc TestCmd) Name() string        { return name }
func (tc TestCmd) Description() string { return description }
func (tc *TestCmd) Execute(args []string) error {
	tc.executionSuccess = true
	tc.args = args
	return nil
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
