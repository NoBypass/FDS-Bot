package components

import "github.com/bwmarrin/discordgo"

var (
	ButtonVerify = &discordgo.Button{
		CustomID: "verify",
		Style:    discordgo.SuccessButton,
		Label:    "Verify",
		Emoji: discordgo.ComponentEmoji{
			Name: "🔗",
		},
	}
)
