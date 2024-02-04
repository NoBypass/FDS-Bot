package discord

import "github.com/bwmarrin/discordgo"

func (i *InteractionCreate) Respond(response *discordgo.InteractionResponse) error {
	return i.Session.InteractionRespond(i.InteractionCreate.Interaction, response)
}

func (i *InteractionCreate) ChannelMessageSendComplex(channelID string, data *discordgo.MessageSend) (*discordgo.Message, error) {
	return i.Session.ChannelMessageSendComplex(channelID, data)
}
