package cmds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/discord"
	"github.com/nobypass/fds-bot/internal/pkg/helpers"
	"math/rand"
)

var VCTeams = &discord.Command{
	ApplicationCommand: vcTeams,
	Handler:            vcTeamsHandler,
}

var vcTeams = &discordgo.ApplicationCommand{
	Name:        "vcteams",
	Description: "Generate random teams from the members in your voice channel",
	Version:     "v1.0.0",
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

func vcTeamsHandler(i *discord.InteractionCreate) error {
	om := helpers.OptionMap(i.ApplicationCommandData().Options)
	teamAmount, tOk := om["teams"].(float64)
	memberAmount, mOk := om["members"].(float64)

	if tOk && mOk {
		return fmt.Errorf("cannot define both memberAmount and teamAmount")
	} else if !tOk && !mOk {
		teamAmount = 2
	}

	userID := i.Member.User.ID
	guildID := i.GuildID
	voiceState, err := i.Session.State.VoiceState(guildID, userID)
	if err != nil {
		return err
	}
	voiceChannelID := voiceState.ChannelID
	guild, err := i.Session.State.Guild(guildID)
	if err != nil {
		return err
	}

	var members []string
	for _, vs := range guild.VoiceStates {
		if vs.ChannelID == voiceChannelID {
			member, err := i.Session.GuildMember(guildID, vs.UserID)
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

	return teamsPrinter(i, teams)
}
