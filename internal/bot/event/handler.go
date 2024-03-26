package event

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/monitoring"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/opentracing/opentracing-go/ext"
	"time"
)

func (m *Manager) Handle(s *discordgo.Session, i *discordgo.InteractionCreate) {
	go func() {
		name := utils.InteractionName(i)
		sp := m.tracer.StartSpan(name)
		defer sp.Finish()

		for _, e := range m.Events {
			if e.Name() == name {

				start := time.Now()
				err := e.Exec(s, i)
				if err != nil {
					// TODO handle error
					ext.LogError(sp, err)
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: err.Error(),
							Flags:   discordgo.MessageFlagsEphemeral,
						},
					})
				}

				monitoring.LogEvent(i, sp, time.Now().Sub(start), err)
				return
			}
		}
	}()
}
