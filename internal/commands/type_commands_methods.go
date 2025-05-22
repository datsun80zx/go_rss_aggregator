package commands

import (
	"errors"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
)

func (c *Commands) run(s *internal.State, cmd Command) error {
	commandFunc, exists := c.Handlers[cmd.Name]
	if !exists {
		return errors.New("Command Not Found")
	}
	return commandFunc(s, cmd)
}

func (c *Commands) register(name string, f func(*internal.State, Command) error) error {
	_, exists := c.Handlers[name]
	if exists {
		return errors.New("Command already registered")
	}
	c.Handlers[name] = f
	return nil
}
