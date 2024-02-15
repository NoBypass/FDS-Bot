package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"math/rand"
)

type VCTeams struct {
}

func (v *VCTeams) Run(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	options := i.ApplicationCommandData().Options
	teamsInput := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "teams"
	})
	membersInput := utils.Find(options, func(o *discordgo.ApplicationCommandInteractionDataOption) bool {
		return o.Name == "members"
	})

	var teamAmount int64
	memberAmount := int64(2)
	if teamsInput != nil {
		teamAmount = teamsInput.IntValue()
	}
	if membersInput != nil {
		memberAmount = membersInput.IntValue()
	}

	userID := i.Member.User.ID
	guildID := i.GuildID
	voiceState, err := s.State.VoiceState(guildID, userID)
	if err != nil {
		return fmt.Errorf("you need to be in a voice channel to use this command")
	}
	voiceChannelID := voiceState.ChannelID
	guild, err := s.State.Guild(guildID)
	if err != nil {
		return err
	}

	var members []string
	for _, vs := range guild.VoiceStates {
		if vs.ChannelID == voiceChannelID {
			member, err := s.GuildMember(guildID, vs.UserID)
			if err != nil {
				return err
			}
			members = append(members, member.Nick)
		}
	}

	var teams [][]string
	if memberAmount != 0 {
		teams = make([][]string, len(members)/int(memberAmount))
	} else {
		teams = make([][]string, int(teamAmount))
	}

	rand.Shuffle(len(members), func(i, j int) {
		members[i], members[j] = members[j], members[i]
	})
	for i, player := range members {
		teams[i%len(teams)] = append(teams[i%len(teams)], player)
	}

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedTeams(teams),
			},
		},
	})
}

func (v *VCTeams) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "vcteams",
		Description: "Generate random teams from the members in your voice channel",
		Version:     "v1.0.1",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "teams",
				Description: "Number of teams (default: 2)",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    &two,
				Required:    false,
			},
			{
				Name:        "members",
				Description: "Number of members per team (takes priority over teams)",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    &two,
				Required:    false,
			},
		},
	}
}
