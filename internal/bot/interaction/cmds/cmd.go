package cmds

import (
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
)

func AllCommands(fds *session.FDSConnection) []event.Event {
	return []event.Event{
		Admin(fds),
		Help(),
		Leaderboard(fds),
		Ping(),
		Play(),
		Profile(fds),
		Revoke(fds),
		Teams(),
		VCTeams(),
	}
}
