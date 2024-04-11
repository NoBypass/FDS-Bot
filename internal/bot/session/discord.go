package session

import (
	"github.com/bwmarrin/discordgo"
	"github.com/labstack/gommon/log"
	"os"
)

func ConnectToDiscord() *discordgo.Session {
	s, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	err = s.Open()
	if err != nil {
		log.Fatal(err)
	}

	return s
}
