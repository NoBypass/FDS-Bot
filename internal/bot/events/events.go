package events

import "log"

type Event struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Event {
	return &Event{logger: logger}
}
