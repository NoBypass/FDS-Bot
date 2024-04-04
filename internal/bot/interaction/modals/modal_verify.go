package modals

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/model"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/opentracing/opentracing-go"
)

type verify struct {
	fds *session.FDSConnection
}

func Verify(fds *session.FDSConnection) event.Modal {
	return &verify{fds: fds}
}

func (v *verify) ID() string {
	return "modal_verify"
}

func (v *verify) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ *event.Context, sp opentracing.Span) (*event.Context, error) {
	ign := i.Interaction.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	resp, err := v.fds.Verify(sp, &model.VerifyRequest{
		ID:   i.Member.User.ID,
		Nick: ign,
		Name: i.Member.User.Username,
	})
	if err != nil {
		return nil, err
	}

	err = s.GuildMemberNickname(i.GuildID, i.Member.User.ID, resp.Actual)
	if err != nil {
		var dcErr *discordgo.RESTError
		ok := errors.As(err, &dcErr)
		if !ok {
			return nil, err
		}
	}

	// TODO: use better way to get role ID
	err = s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, "1179714940395864134") // TODO: use linked role
	if err != nil {
		return nil, err
	}

	return nil, s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: i.Member.User.Mention(),
			Flags:   discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedVerificationDone(resp),
			},
		},
	})
}

func (v *verify) Content(m map[string]any) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: v.ID(),
			Title:    "Verify " + m["username"].(string),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    v.ID(),
							Label:       "What is your Minecraft name?",
							Style:       discordgo.TextInputShort,
							Placeholder: "Your Minecraft name",
							Required:    true,
							MaxLength:   16,
							MinLength:   1,
						},
					},
				},
			},
		},
	}
}
