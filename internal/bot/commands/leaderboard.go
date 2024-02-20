package commands

import (
	"github.com/NoBypass/fds/pkg/api"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
)

type Leaderboard struct {
	api *api.Client
}

func (l *Leaderboard) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	lb, err := l.api.Leaderboard(0) // TODO: implement pagination
	if err != nil {
		return err
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedLeaderboard(s, lb, 0),
			},
		},
	})
}

func (l *Leaderboard) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "leaderboard",
		Description: "View the server leaderboard",
		Version:     "v1.0.0",
	}
}
