package message_components

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

type MessageComponent interface {
	Run(s *discordgo.Session, i *discordgo.InteractionCreate) error
	ComponentID() string
}

type MessageComponentManager struct {
	mcMap  map[string]MessageComponent
	logger *log.Logger
}

func NewManager(logger *log.Logger) *MessageComponentManager {
	return &MessageComponentManager{
		mcMap:  make(map[string]MessageComponent),
		logger: logger,
	}
}

func (mcm *MessageComponentManager) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	customID := i.MessageComponentData().CustomID
	if ar, ok := mcm.mcMap[customID]; ok {
		return ar.Run(s, i)
	}
	return fmt.Errorf("message component not found: %s", customID)
}

func (mcm *MessageComponentManager) RegisterAll() {
	messageComponents := []MessageComponent{
		&VerifyClick{},
	}

	for i, mc := range messageComponents {
		mcm.mcMap[mc.ComponentID()] = mc
		mcm.logger.Printf("Registered message component: %s (%d/%d)", mc.ComponentID(), i+1, len(messageComponents))
	}
}
