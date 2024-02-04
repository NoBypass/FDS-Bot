package discord

import "github.com/bwmarrin/discordgo"

func (s *Session) toInteractionCreate(i *discordgo.InteractionCreate) *InteractionCreate {
	return &InteractionCreate{
		InteractionCreate: i,
		Session:           s,
	}
}
