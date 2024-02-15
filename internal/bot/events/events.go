package events

import (
	"context"
	"log"
)

type Event struct {
	context.Context
	logger *log.Logger
}

func New(logger *log.Logger) *Event {
	return &Event{
		logger:  logger,
		Context: context.Background(),
	}
}
