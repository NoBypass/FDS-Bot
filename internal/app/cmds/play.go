package cmds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/consts"
	"github.com/nobypass/fds-bot/internal/pkg/helpers"
	"strings"
)

const (
	ChannelLTP                = "❓│looking-to-play"
	ChoiceGamemodeScrims      = "scrims"
	ChoiceGamemodeBridgeQueue = "bridge_queue"
	ChoiceGamemodeBedwars     = "bedwars"
	ChoiceGamemodeOther       = "other"
)

var Play = &discordgo.ApplicationCommand{
	Name:        "play",
	Description: "Ask the server to play any gamemode with you",
	Version:     "v1.0.0",
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
					Name:  "Bedwars",
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

func PlayHandler(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	options := i.ApplicationCommandData().Options
	gamemode := helpers.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "gamemode"
	})
	description := helpers.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "description"
	})

	channels, err := s.GuildChannels(i.GuildID)
	if err != nil {
		return err
	}
	ltpChannel := helpers.Find(channels, func(c *discordgo.Channel) bool {
		return c.Name == ChannelLTP
	})
	if ltpChannel == nil {
		return fmt.Errorf("LTP channel not found")
	}

	var roleName string
	switch gamemode.Value {
	case ChoiceGamemodeScrims:
		roleName = "Bridge Player"
	case ChoiceGamemodeBridgeQueue:
		roleName = "Bridge Player"
	case ChoiceGamemodeBedwars:
		roleName = "Bedwars Player"
	}
	roles, err := s.GuildRoles(i.GuildID)
	if err != nil {
		return err
	}
	gamemodeRole := helpers.Find(roles, func(r *discordgo.Role) bool {
		return r.Name == roleName
	})
	if gamemodeRole == nil {
		return fmt.Errorf("gamemode role not found")
	}

	_, err = s.ChannelMessageSendComplex(ltpChannel.ID, &discordgo.MessageSend{
		Content: fmt.Sprintf("By: %s to: %s", i.Member.Mention(), gamemodeRole.Mention()),
		Embed: &discordgo.MessageEmbed{
			Title: fmt.Sprintf("%s is looking to play %s", i.Member.Nick, strings.Title(strings.Replace(fmt.Sprint(gamemode.Value), "_", " ", -1))),
			Description: func() string {
				if description != nil {
					return fmt.Sprintf("**Description**\n%v", description.Value)
				}
				return ""
			}(),
			Color: consts.EmbedColor,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    i.Member.Nick,
				IconURL: i.Member.User.AvatarURL(""),
			},
		},
	})
	if err != nil {
		return err
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Ping was sent to %v", ltpChannel.Mention()),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
