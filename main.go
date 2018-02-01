package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var dc *discordgo.Session

func GetApiKey() string {
	key, exists := os.LookupEnv("DISCORD_API_KEY")
	if exists {
		return key
	} else {
		fmt.Fprintln(os.Stderr, "Cannot find DISCORD_API_KEY enviroment variable")
		os.Exit(1)
	}
	return ""
}

func main() {
	api_key := GetApiKey()
	fmt.Println("hello, go!")

	dc, err := discordgo.New("Bot " + api_key)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot create Discord session: ", err)
	}

	err = dc.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error during opening connection: ", err)
		return
	}

	dc.AddHandler(onMessage)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dc.Close()
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// If it is message (starts with /), parse and execute it
	if m.Content[0] == '/' {
		command, params := ParseCommand(m.Content)
		err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
		if err != nil {
			fmt.Println("Error deleting message: ", err)
		}

		if command != nil {
			command.call(params, s, m)
		} else {
			message := fmt.Sprintf("%s, nieznana komenda. Aby uzyskać listę komend, użyj /bonobot_help", m.Author.Mention())
			s.ChannelMessageSend(m.ChannelID, message)
		}
	}
}
