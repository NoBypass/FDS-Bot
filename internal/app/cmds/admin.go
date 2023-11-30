package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/app/consts"
	"github.com/nobypass/fds-bot/internal/pkg/helpers"
)

var adminPerms = int64(discordgo.PermissionAdministrator)

var Admin = &discordgo.ApplicationCommand{
	Name:                     "admin",
	Description:              "Admin utilities",
	Version:                  "v1.0.0",
	DefaultMemberPermissions: &adminPerms,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "embed",
			Description: "Write an embed to the channel",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "region_selector",
					Value: "region_selector",
				},
			},
		},
	},
}

func AdminHandler(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	om := helpers.OptionMap(i.ApplicationCommandData().Options)
	modal := om["modal"].(string)

	var res *discordgo.MessageSend

	switch modal {
	case "region_selector":
		res = &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Region Selector",
					Color:       consts.Purple,
					Description: "*Pick the region closest to you.*",
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.SelectMenu{
					CustomID: "region_selector",
					MenuType: discordgo.StringSelectMenu,
					Options: []discordgo.SelectMenuOption{
						{
							Label: "EU",
							Value: "eu",
						},
						{
							Label: "NA",
							Value: "na",
						},
					},
				},
			},
		}
	default:
		res = &discordgo.MessageSend{
			Content: "Invalid modal",
		}
	}

	_, err := s.ChannelMessageSendComplex(i.ChannelID, res)
	return err
}
