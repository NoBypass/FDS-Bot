package cmds

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/interaction/modals"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/opentracing/opentracing-go"
)

type revoke struct {
	fds *session.FDSConnection
}

func Revoke(fds *session.FDSConnection) event.Command {
	return &revoke{fds: fds}
}

func (r *revoke) ID() string {
	return "revoke"
}

func (r *revoke) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, ctx *event.Context, _ opentracing.Span) (*event.Context, error) {
	ctx.Set("revoke_interaction", i)

	return ctx, s.InteractionRespond(i.Interaction,
		modals.Revoke(r.fds).Content(map[string]any{
			"name": i.Member.Nick,
		}))
}

func (r *revoke) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:                     "revoke",
		Version:                  "v1.0.0",
		DefaultMemberPermissions: &utils.AdminPerms,
		Type:                     discordgo.UserApplicationCommand,
	}
}
