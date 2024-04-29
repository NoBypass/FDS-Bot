package event

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"time"
)

func (m *Manager) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				m.logger.Errorf("panic: %v", r)
			}
		}()

		name := utils.InteractionName(i)
		ev, ok := m.Events[name]
		if !ok {
			return
		}

		sp := m.tracer.StartSpan(name)
		defer sp.Finish()
		start := time.Now()

		untypedCtx, ok := m.cache.Get(i.Member.User.ID)
		ctx := InitContext(i.Member)
		if ok {
			ctx = untypedCtx.(*Context)
		}

		newCtx, err := ev.Exec(s, i, ctx, sp)
		if newCtx != nil {
			m.cache.Set(i.Member.User.ID, newCtx, 2*time.Minute)
		}

		if err != nil {
			ext.LogError(sp, err)
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						components.EmbedError(err, sp),
					},
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
		}

		logData := map[string]any{
			"type":             i.Type,
			"interaction_name": name,
			"latency":          time.Since(start).String(),
			"channel":          i.ChannelID,
			"member":           i.Member.User.ID,
			"username":         i.Member.User.Username,
			"interaction_id":   i.Interaction.ID,
			"guild":            i.GuildID,
			"message": func() string {
				if err != nil {
					return fmt.Sprintf("error: %v", err)
				}
				return fmt.Sprintf("interaction: %s (type %d) by %s (%s)", i.Interaction.ID, i.Type, i.Member.User.ID, i.Member.User.Username)
			}(),
			"level": func() string {
				if err != nil {
					return "ERROR"
				}
				return "INFO"
			}(),
			"trace_id": sp.Context().(jaeger.SpanContext).TraceID().String(),
		}

		sp.LogKV(
			"type", logData["type"],
			"name", logData["name"],
			"guild", logData["guild"],
			"channel", logData["channel"],
			"member", logData["member"],
			"username", logData["username"],
			"interaction_id", logData["interaction_id"],
		)

		m.logger.Infoj(logData)
	}()
}
