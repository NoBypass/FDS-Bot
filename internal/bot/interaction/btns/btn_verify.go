package btns

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/interaction/modals"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/opentracing/opentracing-go"
)

type verify struct {
	fds *session.FDSConnection
}

func Verify(fds *session.FDSConnection) event.Button {
	return &verify{fds}
}

func (v *verify) ID() string {
	return "btn_verify"
}

func (v *verify) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ *event.Context, _ opentracing.Span) (*event.Context, error) {
	return nil, s.InteractionRespond(i.Interaction, modals.Verify(v.fds).Content(map[string]any{
		"username": i.Member.User.Username,
	}))
}

func (v *verify) Content() *discordgo.Button {
	return &discordgo.Button{
		CustomID: v.ID(),
		Style:    discordgo.SuccessButton,
		Label:    "Verify",
		Emoji: discordgo.ComponentEmoji{
			Name: "🔗",
		},
	}
}
