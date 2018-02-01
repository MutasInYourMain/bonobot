package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// Command interface represents functions that should be implemented by any command
type Command interface {
	description() string
	syntax() string
	call(params []string, s *discordgo.Session, m *discordgo.MessageCreate) error
}

// commandList is list of all Bonobot commands
var commandList = map[string]Command{
	"bonobot_help": HelpCommand{},
}

// ParseCommand parses given string into Command and slice of strings which contains parameters.
// If given comand does not exist, it returns null command
func ParseCommand(toParse string) (Command, []string) {
	splitCommand := strings.Split(toParse, " ")
	commandName := strings.TrimLeft(splitCommand[0], "/")
	command, exists := commandList[commandName]
	if exists {
		return command, splitCommand[1:]
	}
	return nil, nil
}
