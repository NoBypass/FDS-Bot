package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/opentracing/opentracing-go"
)

type daily struct {
	fds *session.FDSConnection
}

func Daily(fds *session.FDSConnection) event.Command {
	return &daily{fds: fds}
}

func (d *daily) ID() string {
	return "daily"
}

func (d *daily) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ *event.Context, sp opentracing.Span) (*event.Context, error) {
	before, err := d.fds.Member(sp, i.Member.User.ID)
	if err != nil {
		return nil, err
	}

	after, err := d.fds.Daily(sp, i.Member.User.ID)
	if err != nil {
		return nil, err
	}

	return nil, s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedDaily(before, after),
			},
		},
	})
}

func (d *daily) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        d.ID(),
		Description: "Claim your daily reward",
		Version:     "1.0.0",
	}
}
