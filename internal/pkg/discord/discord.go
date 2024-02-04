package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/core"
	"log"
)

type interactionCreateFunc func(i *InteractionCreate) error

type InteractionCreate struct {
	*discordgo.InteractionCreate
	*Session
}

type Session struct {
	*discordgo.Session
	commands     map[string]*Command
	interactions map[string]interactionCreateFunc
	Core         *core.Api
}

type Command struct {
	*discordgo.ApplicationCommand
	Handler interactionCreateFunc
}

func Connect(token string) (*Session, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	err = s.Open()
	if err != nil {
		return nil, err
	}

	newS := &Session{
		Session: s,
	}

	newS.interactions = make(map[string]interactionCreateFunc)
	newS.commands = make(map[string]*Command)

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		newS.interactionHandler(i)
	})
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	newS.Core = core.NewCore()

	return newS, nil
}
