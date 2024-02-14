package events

import (
	"github.com/bwmarrin/discordgo"
)

func (e *Event) OnInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
	case discordgo.InteractionMessageComponent:
	case discordgo.InteractionModalSubmit:
	default:
		e.logger.Println("Unknown interaction type")
	}
}
