package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/interaction/btns"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

type admin struct {
	fds *session.FDSConnection
}

func Admin(fds *session.FDSConnection) event.Command {
	return &admin{fds}
}

func (a *admin) ID() string {
	return "admin"
}

func (a *admin) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ *event.Context, _ opentracing.Span) (*event.Context, error) {
	options := i.ApplicationCommandData().Options
	var message *discordgo.MessageSend
	for _, option := range options {
		switch option.Name {
		case "embed":
			switch option.StringValue() {
			case "verify":
				message = &discordgo.MessageSend{
					Embeds: []*discordgo.MessageEmbed{
						components.EmbedVerify(),
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								btns.Verify(a.fds).Content(),
							},
						},
					},
				}
			}
		}
	}

	_, err := s.ChannelMessageSendComplex(i.ChannelID, message)
	if err != nil {
		return nil, err
	}
	return nil, s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message was sent to channel",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (a *admin) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     a.ID(),
		Description:              "Admin utilities",
		Version:                  "v1.0.3",
		DefaultMemberPermissions: &utils.AdminPerms,
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
