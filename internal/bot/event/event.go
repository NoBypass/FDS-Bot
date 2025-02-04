package event

import (
	"context"
	"github.com/NoBypass/mincache"
	"github.com/bwmarrin/discordgo"
	"github.com/labstack/gommon/log"
	"github.com/opentracing/opentracing-go"
)

type (
	Event interface {
		ID() string
		Exec(*discordgo.Session, *discordgo.InteractionCreate, *Context, opentracing.Span) (*Context, error)
	}

	Manager struct {
		Events map[string]Event
		tracer opentracing.Tracer
		logger *log.Logger
		s      *discordgo.Session
		cache  *mincache.CacheInstance
	}

	Context struct {
		Ctx    context.Context
		Member *discordgo.Member
	}

	Command interface {
		Event
		Content() *discordgo.ApplicationCommand
	}

	Modal interface {
		Event
		Content(map[string]any) *discordgo.InteractionResponse
	}

	Button interface {
		Event
		Content() *discordgo.Button
	}
)

func NewManager(s *discordgo.Session, tracer opentracing.Tracer, c *mincache.CacheInstance, l *log.Logger) *Manager {
	return &Manager{
		Events: make(map[string]Event),
		tracer: tracer,
		logger: l,
		cache:  c,
		s:      s,
	}
}

func InitContext(member *discordgo.Member) *Context {
	return &Context{
		Ctx:    context.Background(),
		Member: member,
	}
}

func (c *Context) Set(key string, value any) {
	c.Ctx = context.WithValue(c.Ctx, key, value)
}

func (c *Context) Get(key string) any {
	return c.Ctx.Value(key)
}

func (m *Manager) Add(e ...Event) {
	for i, ev := range e {
		name := ev.ID()
		m.Events[name] = ev
		switch ev.(type) {
		case Command:
			_, err := m.s.ApplicationCommandCreate(m.s.State.User.ID, "", ev.(Command).Content())
			if err != nil {
				log.Fatal(err)
			}
		}
		m.logger.Infof("Registered event: %s (%d/%d)", name, i+1, len(e))
	}
}

func (m *Manager) Remove() {
	cmds, err := m.s.ApplicationCommands(m.s.State.User.ID, "")
	if err != nil {
		log.Fatal(err)
	}

	for i, cmd := range cmds {
		err := m.s.ApplicationCommandDelete(m.s.State.User.ID, "", cmd.ID)
		if err != nil {
			log.Fatal(err)
		}
		m.logger.Infof("Deleted command: %s (%d/%d)", cmd.Name, i+1, len(cmds))
	}
}
