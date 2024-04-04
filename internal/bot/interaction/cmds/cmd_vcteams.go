package cmds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/components"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/opentracing/opentracing-go"
	"math/rand"
)

type vcTeams struct {
}

func VCTeams() event.Command {
	return &vcTeams{}
}

func (v *vcTeams) ID() string {
	return "vcteams"
}

func (v *vcTeams) Exec(s *discordgo.Session, i *discordgo.InteractionCreate, _ *event.Context, _ opentracing.Span) (*event.Context, error) {
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
		return nil, fmt.Errorf("you need to be in a voice channel to use this command")
	}
	voiceChannelID := voiceState.ChannelID
	guild, err := s.State.Guild(guildID)
	if err != nil {
		return nil, err
	}

	var members []string
	for _, vs := range guild.VoiceStates {
		if vs.ChannelID == voiceChannelID {
			member, err := s.GuildMember(guildID, vs.UserID)
			if err != nil {
				return nil, err
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

	return nil, s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				components.EmbedTeams(teams),
			},
		},
	})
}

var two = 2.0

func (v *vcTeams) Content() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        v.ID(),
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
