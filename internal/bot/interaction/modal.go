package interaction

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/models"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/opentracing/opentracing-go"
)

func AllModals(fds *session.FDSConnection) []event.Event {
	return []event.Event{
		&modalVerify{},
	}
}

type (
	modalVerify struct {
		fds  *session.FDSConnection
		user *discordgo.User
	}
)

func (m *modalVerify) Content() any {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "verify_modal_submit",
			Title:    "Verify " + m.user.Username,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "mc_name",
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

func (m *modalVerify) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, sp opentracing.Span) error {
	ign := i.Interaction.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	resp, err := m.fds.Verify(sp, &models.VerifyRequest{
		ID:   i.Member.User.ID,
		Nick: ign,
		Name: i.Member.User.Username,
	})
	if err != nil {
		return err
	}

	err = s.GuildMemberNickname(i.GuildID, i.Member.User.ID, resp.Actual)
	if err != nil {
		return err
	}

	err = s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, "1179714940395864134") // TODO: use linked role
	if err != nil {
		return err
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: i.Member.User.Mention(),
			Flags:   discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				(&embedVerificationDone{*resp}).Content().(*discordgo.MessageEmbed),
			},
		},
	})
}
