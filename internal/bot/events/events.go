package events

import (
	"github.com/nobypass/fds-bot/internal/bot/commands"
	"github.com/nobypass/fds-bot/internal/bot/message_components"
	"github.com/nobypass/fds-bot/internal/bot/modals"
	"log"
)

type Event struct {
	messageComponentManager *message_components.MessageComponentManager
	commandManager          *commands.CommandManager
	modalManager            *modals.ModalManager
	logger                  *log.Logger
}

func New(logger *log.Logger, commandManager *commands.CommandManager, messageComponentManager *message_components.MessageComponentManager, modalManager *modals.ModalManager) *Event {
	return &Event{
		logger:                  logger,
		modalManager:            modalManager,
		commandManager:          commandManager,
		messageComponentManager: messageComponentManager,
	}
}
