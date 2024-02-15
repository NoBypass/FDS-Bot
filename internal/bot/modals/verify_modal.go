package modals

import (
	"github.com/NoBypass/fds/pkg/api"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
)

type VerifySubmit struct {
	api *api.Client
}

func (v *VerifySubmit) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	ign := i.Interaction.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	resp, err := v.api.Verify(&api.DiscordVerifyRequest{
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
	} // TODO: proper error handling

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
				components.EmbedVerificationDone(*resp),
			},
		},
	})
}

func (v *VerifySubmit) ModalID() string {
	return components.ModalVerify(&discordgo.User{Username: ""}).Data.CustomID
}
