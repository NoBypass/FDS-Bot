package cmds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/consts"
	"github.com/nobypass/fds-bot/internal/pkg/discord"
)

var two = 2.0

var teamsPrinter = func(i *discord.InteractionCreate, teams [][]string) error {
	return i.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Teams",
					Color: consts.EmbedColor,
					Fields: func() []*discordgo.MessageEmbedField {
						var fields []*discordgo.MessageEmbedField
						for i, team := range teams {
							var val string
							for _, player := range team {
								val += fmt.Sprintf("` %v `\n", player)
							}
							fields = append(fields, &discordgo.MessageEmbedField{
								Inline: true,
								Name:   fmt.Sprintf("Team %v", i+1),
								Value:  val,
							})
						}
						return fields
					}(),
				},
			},
		},
	})
}
