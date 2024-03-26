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
		name := utils.InteractionName(i)
		sp := m.tracer.StartSpan(name)
		defer sp.Finish()

		e, ok := m.Events[name]
		if !ok {
			return
		}

		start := time.Now()
		err := e.Exec(s, i, sp)
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

		monitoring.LogEvent(i, sp, time.Now().Sub(start), err)
	}()
}
