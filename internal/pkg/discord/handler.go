package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

func (s *Session) interactionHandler(i *discordgo.InteractionCreate) {
	defer func() {
		if r := recover(); r != nil {
			r = fmt.Errorf("(recovered) panic: %v", r)
			respondErr(s, i, r.(error))
			log.Print(r)
		}
	}()

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if h, ok := s.commands[i.ApplicationCommandData().Name]; ok {
			err := h.Handler(s.toInteractionCreate(i))
			if err != nil {
				panic(err)
			}
		}
	case discordgo.InteractionMessageComponent:
		err := s.interactions[i.MessageComponentData().CustomID](s.toInteractionCreate(i))
		if err != nil {
			panic(err)
		}
	default:
		log.Printf("Unknown interaction type: %v", i.Type)
		respondErr(s, i, fmt.Errorf("unknown interaction type: %v", i.Type))
	}
}

func respondErr(s *Session, i *discordgo.InteractionCreate, err error) {
	e := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: fmt.Sprintf("Oops, something went wrong: %v\n\nIf this keeps happening, please contact staff.", err),
		},
	})
	if e != nil {
		log.Printf("Cannot send error message: %v", e)
	}
}
