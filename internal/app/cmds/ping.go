package cmds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/discord"
)

var Ping = &discord.Command{
	ApplicationCommand: ping,
	Handler:            pingHandler,
}

var ping = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Get the bot's ping",
	Version:     "v1.1.1",
}

func pingHandler(i *discord.InteractionCreate) error {
	latency := i.Session.HeartbeatLatency().Milliseconds()

	return i.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Your ping to the bot (EU) is %vms", latency),
		},
	})
}
