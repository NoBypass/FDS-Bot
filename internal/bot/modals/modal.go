package modals

import (
	"github.com/NoBypass/fds/pkg/api"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Modal interface {
	Run(s *discordgo.Session, i *discordgo.InteractionCreate) error
	ModalID() string
}

type ModalManager struct {
	mMap   map[string]Modal
	logger *log.Logger
	api    *api.Client
}

func NewManager(logger *log.Logger, apiClient *api.Client) *ModalManager {
	return &ModalManager{
		api:    apiClient,
		mMap:   make(map[string]Modal),
		logger: logger,
	}
}

func (mm *ModalManager) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	modalID := i.ModalSubmitData().CustomID
	if m, ok := mm.mMap[modalID]; ok {
		return m.Run(s, i)
	}
	return nil
}

func (mm *ModalManager) RegisterAll() {
	modals := []Modal{
		&VerifySubmit{api: mm.api},
	}

	for i, m := range modals {
		mm.mMap[m.ModalID()] = m
		mm.logger.Printf("Registered modal: %s (%d/%d)", m.ModalID(), i+1, len(modals))
	}
}
