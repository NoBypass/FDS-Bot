package components

import "github.com/bwmarrin/discordgo"

var (
	ModalVerify = func(user *discordgo.User) *discordgo.InteractionResponse {
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{
				CustomID: "verify_modal_submit",
				Title:    "Verify " + user.Username,
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
)
