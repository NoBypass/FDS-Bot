package cmds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/app/consts"
	"strings"
)

var two = 2.0

var teamsPrinter = func(s *discordgo.Session, i *discordgo.InteractionCreate, teams [][]string) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Teams",
					Color: consts.Purple,
					Fields: func() []*discordgo.MessageEmbedField {
						var fields []*discordgo.MessageEmbedField
						for i, team := range teams {
							fields = append(fields, &discordgo.MessageEmbedField{
								Name:  fmt.Sprintf("Team %v", i+1),
								Value: strings.Join(team, "\n"),
							})
						}
						return fields
					}(),
				},
			},
		},
	})
}
