package commands

import (
	"github.com/NoBypass/fds/pkg/api"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
)

type Profile struct {
	api *api.Client
}

func (p *Profile) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	option := i.ApplicationCommandData().Options[0]
	var id string
	if option == nil {
		id = i.Member.User.ID
	} else {
		id = option.UserValue(s).ID
	}

	member, err := p.api.Member(id)
	if err != nil {
		return err
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedProfile(member),
			},
		},
	})
}

func (p *Profile) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "profile",
		Description: "View the profile of a member",
		Version:     "v1.0.0",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionUser,
				Name:        "member",
				Description: "The member to view the profile of",
				Required:    false,
			},
		},
	}
}
