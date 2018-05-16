package cli

import (
	"errors"
	"os"
)

// Cli holds commands and other app info
type Cli struct {
	Name        string
	Description string
	Commands    []Command
}

func New(name string, desc string) *Cli {
	return &Cli{
		Name:        name,
		Description: desc,
		Commands:    []Command{},
	}
}

func (c *Cli) Register(cmd Command) error {
	for _, command := range c.Commands {
		if cmd.Name == command.Name {
			return errors.New("Can't register duplicate commands")
		}
	}

	c.Commands = append(c.Commands, cmd)
	return nil
}

func (c *Cli) Run() {
	if len(os.Args) < 2 {
		panic(errors.New("Invalid command"))
		return
	}

	command := os.Args[1]

	for _, cmd := range c.Commands {
		if cmd.Name == command {
			err := cmd.Action()
			if err != nil {
				panic(err)
			}
		}
	}
}

type Action func() error

// Command represents a cli command
type Command struct {
	Name        string
	Description string
	Action      Action
}
