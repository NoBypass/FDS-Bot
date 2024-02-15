package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
)

type Help struct {
	cmds map[string]Command
}

func (h *Help) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedHelp(s, func() map[string]*discordgo.ApplicationCommand {
					m := make(map[string]*discordgo.ApplicationCommand)
					for name, cmd := range h.cmds {
						m[name] = cmd.Content()
					}
					return m
				}()),
			},
		},
	})
}

func (h *Help) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "help",
		Description: "View the help menu",
		Version:     "v1.0.0",
	}
}
