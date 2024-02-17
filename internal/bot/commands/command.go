package commands

import (
	"fmt"
	"github.com/NoBypass/fds/pkg/api"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Command interface {
	Run(s *discordgo.Session, i *discordgo.InteractionCreate) error
	Content() *discordgo.ApplicationCommand
}

type CommandManager struct {
	m      map[string]Command
	api    *api.Client
	logger *log.Logger
}

func NewManager(logger *log.Logger, apiClient *api.Client) *CommandManager {
	return &CommandManager{
		api:    apiClient,
		m:      make(map[string]Command),
		logger: logger,
	}
}

func (cm *CommandManager) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	name := i.ApplicationCommandData().Name
	c, ok := cm.m[name]
	if !ok {
		return fmt.Errorf("command not found: %s", name)
	}
	return c.Run(s, i)
}

func (cm *CommandManager) RegisterAll(s *discordgo.Session) error {
	commands := []Command{
		&Ping{},
		&Help{cm.m},
		&Admin{},
		&Play{},
		&Teams{},
		&Profile{cm.api},
		&VCTeams{},
	}

	for i, c := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", c.Content())
		if err != nil {
			return err
		}
		cm.m[c.Content().Name] = c
		if err != nil {
			return err
		}
		cm.logger.Printf("Registered command: %s (%d/%d)", c.Content().Name, i+1, len(commands))
	}

	return nil
}
