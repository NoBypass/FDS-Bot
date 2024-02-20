package commands

import (
	"fmt"
	"github.com/NoBypass/fds/pkg/api"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io/ioutil"
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
	fontBytes, err := ioutil.ReadFile("assets/font/Inter-Bold.ttf")
	if err != nil {
		return err
	}
	fontParsed, err := opentype.Parse(fontBytes)
	if err != nil {
		return err
	}
	face, err := opentype.NewFace(fontParsed, &opentype.FaceOptions{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return err
	}

	commands := []Command{
		&Ping{},
		&Play{},
		&Teams{},
		&Admin{},
		&VCTeams{},
		&Help{cm.m},
		&Leaderboard{cm.api},
		&Profile{cm.api, face},
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
