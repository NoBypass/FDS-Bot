package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/opentracing/opentracing-go"
)

type help struct {
	fds *session.FDSConnection
}

func Help() event.Command {
	return &help{}
}

func (h *help) ID() string {
	return "help"
}

func (h *help) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ *event.Context, _ opentracing.Span) (*event.Context, error) {
	return nil, s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedHelp(s, func() map[string]*discordgo.ApplicationCommand {
					m := make(map[string]*discordgo.ApplicationCommand)
					for _, cmd := range AllCommands(h.fds) {
						content := cmd.(event.Command).Content()
						m[content.Name] = content
					}
					return m
				}()),
			},
		},
	})
}

func (h *help) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        h.ID(),
		Description: "Get help",
		Version:     "v1.0.0",
	}
}
