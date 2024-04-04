package cmds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/opentracing/opentracing-go"
)

type ping struct {
}

func Ping() event.Command {
	return &ping{}
}

func (p *ping) ID() string {
	return "ping"
}

func (p *ping) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ *event.Context, _ opentracing.Span) (*event.Context, error) {
	latency := s.HeartbeatLatency().Milliseconds()
	return nil, s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("The bots latency is %vms", latency),
		},
	})
}

func (p *ping) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        p.ID(),
		Description: "Ping the bot",
		Version:     "v1.2.0",
	}
}
