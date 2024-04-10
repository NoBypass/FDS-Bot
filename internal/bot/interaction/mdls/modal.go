package mdls

import (
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
)

func AllModals(fds *session.FDSConnection) []event.Event {
	return []event.Event{
		Verify(fds),
		Revoke(fds),
	}
}
