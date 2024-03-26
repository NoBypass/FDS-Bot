package utils

import "github.com/bwmarrin/discordgo"

func InteractionName(i *discordgo.InteractionCreate) string {
	var name string
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		name = i.ApplicationCommandData().Name
	case discordgo.InteractionMessageComponent:
		name = i.MessageComponentData().CustomID
	case discordgo.InteractionModalSubmit:
		name = i.ModalSubmitData().CustomID
	}
	return name
}
