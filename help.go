package main

import (
	"fmt"
	"bytes"
	"github.com/bwmarrin/discordgo"
)

var helpString = `
**Bonobot, v. 0.0.1**
Dostępne komendy: 

`

// HelpCommand is a command that displays descriptions of all commands and Bonobot help string
type HelpCommand struct {
}

func (h HelpCommand) description() string {
	return "wyświetla tą wiadomość"
}

func (h HelpCommand) syntax() string {
	return ""
}

func (h HelpCommand) call(params []string, s *discordgo.Session, m *discordgo.MessageCreate) error {
	var helpMessageBuffer bytes.Buffer
	helpMessageBuffer.WriteString(helpString)
	for commandName, command := range commandList {
		commandHelp := fmt.Sprintf("**/%s** %s - %s\n", commandName, command.syntax(), command.description())
		helpMessageBuffer.WriteString(commandHelp)
	}
	_, err := s.ChannelMessageSend(m.ChannelID, helpMessageBuffer.String())
	return err
}