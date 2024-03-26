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

func ComponentName(obj any) string {
	switch obj.(type) {
	case *discordgo.Button:
		return obj.(*discordgo.Button).CustomID
	case *discordgo.SelectMenu:
		return obj.(*discordgo.SelectMenu).CustomID
	case *discordgo.ApplicationCommand:
		return obj.(*discordgo.ApplicationCommand).Name
	}
	return ""
}
