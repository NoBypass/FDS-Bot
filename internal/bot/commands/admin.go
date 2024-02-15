package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
)

type Admin struct {
}

func (a *Admin) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	options := i.ApplicationCommandData().Options
	var message *discordgo.MessageSend
	for _, option := range options {
		switch option.Name {
		case "embed":
			switch option.StringValue() {
			case "verify":
				message = &discordgo.MessageSend{
					Embeds: []*discordgo.MessageEmbed{
						components.EmbedVerify,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								components.ButtonVerify,
							},
						},
					},
				}
			}
		}
	}

	_, err := s.ChannelMessageSendComplex(i.ChannelID, message)
	if err != nil {
		return err
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message was sent to channel",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

var adminPerms = int64(discordgo.PermissionAdministrator)

func (a *Admin) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     "admin",
		Description:              "Admin utilities",
		Version:                  "v1.0.3",
		DefaultMemberPermissions: &adminPerms,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "embed",
				Description: "Write an embed to a channel",
				Required:    false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "verify",
						Value: "verify",
					},
				},
			},
		},
	}
}
