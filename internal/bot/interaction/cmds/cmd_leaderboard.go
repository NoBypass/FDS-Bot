package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/opentracing/opentracing-go"
)

type leaderboard struct {
	fds *session.FDSConnection
}

func Leaderboard(fds *session.FDSConnection) event.Command {
	return &leaderboard{fds}
}

func (l *leaderboard) ID() string {
	return "leaderboard"
}

func (l *leaderboard) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ *event.Context, sp opentracing.Span) (*event.Context, error) {
	lb, err := l.fds.Leaderboard(sp, 0) // TODO: implement pagination
	if err != nil {
		return nil, err
	}

	return nil, s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedLeaderboard(s, lb, 0),
			},
		},
	})
}

func (l *leaderboard) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        l.ID(),
		Description: "View the server leaderboard",
		Version:     "v1.0.0",
	}
}
