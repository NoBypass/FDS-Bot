package events

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/commands"
)

func (e *Event) OnInteraction(cm *commands.CommandManager) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	e.Context = context.WithValue(e.Context, "commandManager", cm)
	return e.interactionHandler
}

func (e *Event) interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		err := e.Context.Value("commandManager").(*commands.CommandManager).Run(i.ApplicationCommandData().Name, s, i)
		if err != nil {
			fallbackErr := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: err.Error(),
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			if fallbackErr != nil {
				e.logger.Printf("Failed to respond to command: %s", err)
			}
		}
	case discordgo.InteractionMessageComponent:
	case discordgo.InteractionModalSubmit:
	default:
		e.logger.Println("Unknown interaction type")
	}
}
