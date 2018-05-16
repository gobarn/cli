package cli

import (
	"os"
	"testing"
)

func Test_New(t *testing.T) {
	name := "app"
	desc := "desc"
	c := New(name, desc)

	if c.Name != name {
		t.Errorf("Expected %s, Got %s", name, c.Name)
	}

	if c.Description != desc {
		t.Errorf("Expected %s, Got %s", desc, c.Description)
	}
}

func Test_Run(t *testing.T) {
	c := New("app", "desc")

	got := false
	expected := true

	cmd := Command{
		Name:        "foo",
		Description: "Description",
		Action: func() error {
			got = expected
			return nil
		},
	}

	c.Register(cmd)

	old := os.Args
	os.Args = []string{"cli", "foo"}

	c.Run()

	if got != expected {
		t.Errorf("Expected %v, Got %v", expected, got)
	}

	os.Args = old
}

func Test_Register(t *testing.T) {
	c := New("app", "desc")

	cmd := Command{
		Name:        "foo",
		Description: "Description",
		Action:      func() error { return nil },
	}

	c.Register(cmd)

	got := len(c.Commands)
	expected := 1

	if got != expected {
		t.Errorf("Expected %v, Got %v", expected, got)
	}
}

func Test_Register_Duplicate(t *testing.T) {
	c := New("app", "desc")

	cmd := Command{
		Name:        "foo",
		Description: "Description",
		Action:      func() error { return nil },
	}

	err := c.Register(cmd)
	err = c.Register(cmd)

	got := len(c.Commands)
	expected := 1

	if got != expected {
		t.Errorf("Expected %v, Got %v", expected, got)
	}

	if err == nil {
		t.Errorf("Expected an error")
	}
}
