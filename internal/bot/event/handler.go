package event

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/monitoring"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"time"
)

func (m *Manager) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	go func() {
		// TODO recover from panics
		name := utils.InteractionName(i)
		sp := m.tracer.StartSpan(name)
		defer sp.Finish()
		start := time.Now()

		ev, ok := m.Events[name]
		if !ok {
			return
		}

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
					Content: fmt.Sprintf("An error occurred: %s\n"+
						"If you think that this is not intended behaviour, "+
						"please send the following id to an admin: `%s`", err.Error(), sp.Context().(jaeger.SpanContext).TraceID().String()),
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
		}

		monitoring.LogEvent(i, sp, time.Since(start), err)
	}()
}
