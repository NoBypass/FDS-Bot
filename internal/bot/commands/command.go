package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Command interface {
	Run(s *discordgo.Session, i *discordgo.InteractionCreate) error
	Content() *discordgo.ApplicationCommand
}

type CommandManager struct {
	m      map[string]Command
	logger *log.Logger
}

func NewCommandManager(logger *log.Logger) *CommandManager {
	return &CommandManager{
		m:      make(map[string]Command),
		logger: logger,
	}
}

func (cm *CommandManager) Run(name string, s *discordgo.Session, i *discordgo.InteractionCreate) error {
	c, ok := cm.m[name]
	if !ok {
		return fmt.Errorf("command not found: %s", name)
	}
	return c.Run(s, i)
}

func (cm *CommandManager) Register(s *discordgo.Session, c Command) error {
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", c.Content())
	if err != nil {
		return err
	}
	cm.m[c.Content().Name] = c
	return nil
}

func (cm *CommandManager) RegisterAll(s *discordgo.Session) error {
	commands := []Command{
		&Ping{},
		&Help{cm.m},
		&Admin{},
		&Play{},
	}

	for _, c := range commands {
		err := cm.Register(s, c)
		if err != nil {
			return err
		}
		cm.logger.Printf("Registered command: %s", c.Content().Name)
	}

	return nil
}
