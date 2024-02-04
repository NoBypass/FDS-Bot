package interactions

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/consts"
	"github.com/nobypass/fds-bot/internal/pkg/discord"
)

func VerifyHandler(i *discord.InteractionCreate) error {
	return i.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "verify_modal_submit",
			Title:    "Verify " + i.Interaction.Member.User.Username,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "mc_name",
							Label:       "What time is your Minecraft name?",
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
	})
}

func VerifyModalSubmitHandler(i *discord.InteractionCreate) error {
	mcName := i.Interaction.ModalSubmitData().Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	resp, err := i.Session.Core.Verify(i.Member.User.ID, i.Member.User.Username, mcName)
	if err != nil {
		return err
	}

	desc := fmt.Sprintf("This discord account has been linked to `%v` via Hypixel.\n\nInfo: you will soon not be able to see this channel anymore.", resp.Actual)

	return i.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: i.Member.User.Mention(),
			Flags:   discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "You have been verified!",
					Description: desc,
					Color:       consts.EmbedColor,
				},
			},
		},
	})
}
