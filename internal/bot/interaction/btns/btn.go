package btns

import (
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
)

func AllButtons(fds *session.FDSConnection) []event.Event {
	return []event.Event{
		Verify(fds),
	}
}
