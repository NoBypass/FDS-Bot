package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
)

const (
	ChannelLTP                = "❓│looking-to-play"
	ChoiceGamemodeScrims      = "scrims"
	ChoiceGamemodeBridgeQueue = "bridge_queue"
	ChoiceGamemodeBedwars     = "bedwars"
	ChoiceGamemodeOther       = "other"
)

type Play struct {
}

func (p *Play) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	options := i.ApplicationCommandData().Options
	gamemode := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "gamemode"
	})
	description := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "description"
	})

	channels, err := s.GuildChannels(i.GuildID)
	if err != nil {
		return err
	}
	ltpChannel := utils.Find(channels, func(c *discordgo.Channel) bool {
		return c.Name == ChannelLTP // TODO: per guild channel name
	})
	if ltpChannel == nil {
		return fmt.Errorf("LTP channel not found")
	}
	if i.ChannelID != ltpChannel.ID {
		return fmt.Errorf("command can only be used in %v", ltpChannel.Mention())
	}

	var roleName string // TODO: per guild role names
	switch gamemode.Value {
	case ChoiceGamemodeScrims:
		roleName = "Notify Bridge Scrims"
	case ChoiceGamemodeBridgeQueue:
		roleName = "Notify Bridge Queue"
	case ChoiceGamemodeBedwars:
		roleName = "Notify BedWars"
	case ChoiceGamemodeOther:
		roleName = "Notify Random"
	}
	roles, err := s.GuildRoles(i.GuildID)
	if err != nil {
		return err
	}
	gamemodeRole := utils.Find(roles, func(r *discordgo.Role) bool {
		return r.Name == roleName
	})
	if gamemodeRole == nil {
		return fmt.Errorf("gamemode role not found")
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("By: %s to: %s", i.Member.Mention(), gamemodeRole.Mention()),
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedPlay(i.Member, gamemodeRole, description, func() (name string) {
					switch gamemode.Value {
					case ChoiceGamemodeScrims:
						name = "Scrims"
					case ChoiceGamemodeBridgeQueue:
						name = "Bridge Queue"
					case ChoiceGamemodeBedwars:
						name = "BedWars"
					case ChoiceGamemodeOther:
						name = "Something"
					}
					return
				}()),
			},
		},
	})
}

func (p *Play) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "play",
		Description: "Ask the server to play any gamemode with you",
		Version:     "v1.1.0",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "gamemode",
				Description: "The gamemode you want to play",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Value: ChoiceGamemodeScrims,
						Name:  "Scrims",
					},
					{
						Value: ChoiceGamemodeBridgeQueue,
						Name:  "Bridge Queue",
					},
					{
						Value: ChoiceGamemodeBedwars,
						Name:  "BedWars",
					},
					{
						Value: ChoiceGamemodeOther,
						Name:  "Other",
					},
				},
			},
			{
				Name:        "description",
				Description: "A description of what you want to play (e.g. submode, players, map etc.)",
				Type:        discordgo.ApplicationCommandOptionString,
				MaxLength:   512,
				Required:    false,
			},
		},
	}
}
