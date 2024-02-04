package interactions

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/discord"
)

func VerifyHandler(i *discord.InteractionCreate) error {
	return i.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "verify_modal_" + i.Interaction.Member.User.ID,
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
	if i.Type != discordgo.InteractionModalSubmit {
		return fmt.Errorf("unexpected interaction type: %v", i.Type)
	}

	// values := i.MessageComponentData().Values
	// minecraftName := values[0]

	// Perform necessary actions with minecraftName

	return nil
}
