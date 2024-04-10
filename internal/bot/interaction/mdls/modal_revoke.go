package mdls

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/opentracing/opentracing-go"
	"strings"
)

type modalRevoke struct {
	fds *session.FDSConnection
}

func Revoke(fds *session.FDSConnection) event.Modal {
	return &modalRevoke{fds: fds}
}

func (m *modalRevoke) ID() string {
	return "modal_revoke"
}

func (m *modalRevoke) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, ctx *event.Context, sp opentracing.Span) (*event.Context, error) {
	typedName := i.Interaction.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	revokeInteraction, ok := ctx.Get("revoke_interaction").(*discordgo.InteractionCreate)
	if !ok {
		return nil, fmt.Errorf("took too long to respond to the modal, please try again")
	}

	expectedName := revokeInteraction.Member.Nick
	if strings.ToLower(typedName) != strings.ToLower(expectedName) {
		return nil, s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   discordgo.MessageFlagsEphemeral,
				Content: fmt.Sprintf("The name you typed does not match the user you want to revoke. (%s != %s)", typedName, expectedName),
			},
		})
	}

	resp, err := m.fds.Revoke(sp, i.Member.User.ID)
	if err != nil {
		return nil, err
	}

	// TODO: use different way to get role ID
	memberID := revokeInteraction.Member.User.ID
	err = s.GuildMemberRoleRemove(i.GuildID, memberID, "1179714940395864134") // TODO: use linked role
	if err != nil {
		return nil, err
	}

	err = s.GuildMemberNickname(i.GuildID, memberID, resp.Name)
	if err != nil {
		var dcErr *discordgo.RESTError
		ok := errors.As(err, &dcErr)
		if !ok {
			return nil, err
		}
	}

	return nil, s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedRevoked(resp),
			},
		},
	})
}

func (m *modalRevoke) Content(params map[string]any) *discordgo.InteractionResponse {
	name := params["name"].(string)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: m.ID(),
			Title:    "Confirm Revocation",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    m.ID(),
							Label:       "Type their username to confirm",
							Placeholder: name,
							Style:       discordgo.TextInputShort,
							Required:    true,
						},
					},
				},
			},
		},
	}
}
