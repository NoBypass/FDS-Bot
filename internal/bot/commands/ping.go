package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Ping struct {
}

func (p *Ping) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	latency := s.HeartbeatLatency().Milliseconds()
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("The bots latency is %vms", latency),
		},
	})
}

func (p *Ping) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping the bot",
		Version:     "v1.2.0",
	}
}
