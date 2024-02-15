package message_components

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
)

type VerifyClick struct {
}

func (v *VerifyClick) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return s.InteractionRespond(i.Interaction, components.ModalVerify(i.Member.User))
}

func (v *VerifyClick) ComponentID() string {
	return components.ButtonVerify.CustomID
}
