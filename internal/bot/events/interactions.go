package events

import (
	"github.com/bwmarrin/discordgo"
)

func (e *Event) OnInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		err := e.commandManager.Run(s, i)
		if err != nil {
			e.fallback(s, i, err)
		}
	case discordgo.InteractionMessageComponent:
		err := e.messageComponentManager.Run(s, i)
		if err != nil {
			e.fallback(s, i, err)
		}
	case discordgo.InteractionModalSubmit:
		err := e.modalManager.Run(s, i)
		if err != nil {
			e.fallback(s, i, err)
		}
	default:
		e.logger.Println("Unknown interaction type")
	}
}

func (e *Event) fallback(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
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
