package cmds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/consts"
	"github.com/nobypass/fds-bot/internal/pkg/discord"
)

var Daily = &discord.Command{
	ApplicationCommand: daily,
	Handler:            dailyHandler,
}

var daily = &discordgo.ApplicationCommand{
	Name:        "daily",
	Description: "Claim some xp daily",
	Version:     "v1.0.1",
}

func dailyHandler(i *discord.InteractionCreate) error {
	daily, err := i.Session.Core.Daily(i.Member.User.ID)
	if err != nil {
		return err
	}

	return i.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: fmt.Sprintf("You get %v xp as a daily reward!", daily.Gained),
					Color: consts.EmbedColor,
					// Image: &discordgo.MessageEmbedImage{

					// }
				},
			},
		},
	})
}
