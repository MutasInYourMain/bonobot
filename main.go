package main;
import (
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

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

	for {
		dc.ChannelMessageSendTTS("273429211610480643", "Bonobo nigdy nie miał kobiety")
		time.Sleep(time.Second * 30)
	}

}