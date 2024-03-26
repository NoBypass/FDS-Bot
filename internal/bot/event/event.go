package event

import (
	"github.com/bwmarrin/discordgo"
	"github.com/opentracing/opentracing-go"
)

type Event interface {
	Exec(*discordgo.Session, *discordgo.InteractionCreate, opentracing.Span) error
	Content() any
	Name() string
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
		m.Events[ev.Name()] = ev
		content := ev.Content()
		switch content.(type) {
		case *discordgo.ApplicationCommand:
			_, err := m.s.ApplicationCommandCreate(m.s.State.User.ID, "", content.(*discordgo.ApplicationCommand))
			if err != nil {
				panic(err)
			}
		}
		println("Registered event: " + ev.Name())
	}
}
