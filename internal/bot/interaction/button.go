package interaction

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/opentracing/opentracing-go"
)

func AllButtons(fds *session.FDSConnection) []event.Event {
	return []event.Event{
		&buttonVerify{fds},
	}
}

type (
	btn struct {
		fds *session.FDSConnection
	}

	buttonVerify btn
)

func (b *buttonVerify) Content() any {
	return &discordgo.Button{
		CustomID: "btn_verify",
		Style:    discordgo.SuccessButton,
		Label:    "Verify",
		Emoji: discordgo.ComponentEmoji{
			Name: "🔗",
		},
	}
}

func (b *buttonVerify) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ opentracing.Span) error {
	return s.InteractionRespond(i.Interaction,
		(&modalVerify{b.fds, i.Member.User.Username}).Content().(*discordgo.InteractionResponse))
}
