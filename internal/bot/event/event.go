package event

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

type Event interface {
	Exec(*discordgo.Session, *discordgo.InteractionCreate, opentracing.Span) error
	Content() any
}

type Manager struct {
	Events map[string]Event
	tracer opentracing.Tracer
	s      *discordgo.Session
}

func NewManager(s *discordgo.Session, tracer opentracing.Tracer) *Manager {
	return &Manager{
		Events: make(map[string]Event),
		tracer: tracer,
		s:      s,
	}
}

func (m *Manager) Add(e ...Event) {
	for _, ev := range e {
		content := ev.Content()
		name := utils.ComponentName(content)
		m.Events[name] = ev
		switch content.(type) {
		case *discordgo.ApplicationCommand:
			_, err := m.s.ApplicationCommandCreate(m.s.State.User.ID, "", content.(*discordgo.ApplicationCommand))
			if err != nil {
				panic(err)
			}
		}
		println("Registered event: " + name)
	}
}
