package session

import (
	"github.com/bwmarrin/discordgo"
	"os"
)

func ConnectToDiscord() *discordgo.Session {
	s, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}
	err = s.Open()
	if err != nil {
		panic(err)
	}

	return s
}
