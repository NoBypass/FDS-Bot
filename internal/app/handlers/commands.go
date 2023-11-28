package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/app/cmds"
	"log"
)

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error{
	"ping": cmds.PingHandler,
}

var commands = []*discordgo.ApplicationCommand{
	cmds.Ping,
}

func RegisterCommands(s *discordgo.Session) {
	for _, v := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", v.Name, err)
		}
	}
}

func Commands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		err := h(s, i)
		if err != nil {
			log.Printf("Cannot handle '%v' command: %v", i.ApplicationCommandData().Name, err)
			return
		}
	}
}
