package interaction

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
)

func AllCommands(fds *session.FDSConnection) []event.Event {
	return []event.Event{
		&ping{fds},
	}
}

type (
	cmd struct {
		fds *session.FDSConnection
	}

	ping cmd
)

func (p *ping) Exec(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	latency := s.HeartbeatLatency().Milliseconds()
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("The bots latency is %vms", latency),
		},
	})
}

func (p *ping) Register(s *discordgo.Session) {
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping the bot",
		Version:     "v1.2.0",
	})
	if err != nil {
		panic(err)
	}
}

func (p *ping) Name() string {
	return "ping"
}
