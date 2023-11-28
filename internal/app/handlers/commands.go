package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/app/cmds"
	"log"
)

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error{
	"ping":  cmds.PingHandler,
	"teams": cmds.TeamsHandler,
}

var commands = []*discordgo.ApplicationCommand{
	cmds.Ping,
	cmds.Teams,
}

func RegisterCommands(s *discordgo.Session) {
	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commands)
	if err != nil {
		log.Fatalf("Cannot register commands: %v", err)
	}
}

func Commands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	defer func() {
		if r := recover(); r != nil {
			respondErr(s, i, fmt.Errorf("(recovered) panic: %v", r))
		}
	}()

	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		err := h(s, i)
		if err != nil {
			respondErr(s, i, err)
			return
		}
	}
}

func respondErr(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	e := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: fmt.Sprintf("Oops, something went wrong: %v\n\nIf this keeps happening, please contact staff as this is likely an easy fix", err),
		},
	})
	if e != nil {
		log.Printf("Cannot send error message: %v", e)
	}
}
