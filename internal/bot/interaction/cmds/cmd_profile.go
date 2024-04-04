package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/opentracing/opentracing-go"
)

type profile struct {
	fds *session.FDSConnection
}

func Profile(fds *session.FDSConnection) event.Command {
	return &profile{fds}
}

func (p *profile) ID() string {
	return "profile"
}

func (p *profile) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ *event.Context, sp opentracing.Span) (*event.Context, error) {
	option := i.ApplicationCommandData().Options[0]
	var id string
	if option == nil {
		id = i.Member.User.ID
	} else {
		id = option.UserValue(s).ID
	}

	member, err := p.fds.Member(sp, id)
	if err != nil {
		return nil, err
	}

	return nil, s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedProfile(member),
			},
		},
	})
}

func (p *profile) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "profile",
		Description: "View the profile of a member",
		Version:     "v1.0.0",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "member",
				Description: "The member to view the profile of",
				Required:    false,
			},
		},
	}
}
