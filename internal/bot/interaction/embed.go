package interaction

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/models"
	"github.com/nobypass/fds-bot/internal/pkg/version"
)

type (
	embed struct{}

	embedProfile          struct{ member *models.MemberResponse }
	embedVerificationDone struct{ resp models.VerifyResponse }
	embedTeams            struct{ teams [][]string }
	embedPlay             struct {
		member *discordgo.Member
		to     *discordgo.Role
		desc   *discordgo.ApplicationCommandInteractionDataOption
		mode   string
	}
	embedVerify embed
	embedHelp   struct {
		s    *discordgo.Session
		cmds map[string]*discordgo.ApplicationCommand
	}
)

func (e *embedProfile) Content() any {
	return &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Profile of %s", e.member.Nick),
		Color: 0x2B2D31,
		// TODO: Add additional information
	}
}

func (e *embedVerificationDone) Content() any {
	return &discordgo.MessageEmbed{
		Title: "You are verified!",
		Color: 0x2B2D31,
		Description: fmt.Sprintf("This discord account has been linked to `%v` via Hypixel.\n\n"+
			"Info: you will soon not be able to see this channel anymore.", e.resp.Actual),
	}
}

func (e *embedTeams) Content() any {
	return &discordgo.MessageEmbed{
		Title: "Teams",
		Color: 0x2B2D31,
		Fields: func() []*discordgo.MessageEmbedField {
			var fields []*discordgo.MessageEmbedField
			for i, team := range e.teams {
				var val string
				for _, player := range team {
					val += fmt.Sprintf("` %v `\n", player)
				}
				fields = append(fields, &discordgo.MessageEmbedField{
					Inline: true,
					Name:   fmt.Sprintf("Team %v", i+1),
					Value:  val,
				})
			}
			return fields
		}(),
	}
}

// TODO: add use for to
func (e *embedPlay) Content() any {
	return &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s is looking to play %s", e.member.User.Username, e.mode),
		Description: func() string {
			if e.desc != nil {
				return fmt.Sprintf("**Description**\n%v", e.desc.Value)
			}
			return ""
		}(),
		Color: 0x2B2D31,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    e.member.User.Username,
			IconURL: e.member.User.AvatarURL("64"),
		},
	}
}

func (e *embedVerify) Content() any {
	return &discordgo.MessageEmbed{
		Title: "Verify",
		Color: 0x2B2D31,
		Description: `Verify your Discord account by linking it to Hypixel.

If it does not work, it might be because you have not linked your Discord account to your Hypixel account. [This video](https://www.youtube.com/watch?v=UresIQdoQHk) will show you how to do it.

**THIS STEP IS REQUIRED TO ACCESS THE SERVER**`,
	}
}
func (e *embedHelp) Content() any {
	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("FDS Bot %s", version.VERSION),
		Description: "List of available commands",
		Color:       0x00ff00,
		Fields: func() []*discordgo.MessageEmbedField {
			fields := make([]*discordgo.MessageEmbedField, 0, len(e.cmds))
			for name, cmd := range e.cmds {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name:   fmt.Sprintf("%s %s", name, cmd.Version),
					Value:  fmt.Sprintf("`/%s`: %s", cmd.Name, cmd.Description),
					Inline: false,
				})
			}
			return fields
		}(),
		URL: "https://github.com/NoBypass/fds-bot",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Bot by NoBypass",
			IconURL: func() string {
				user, err := e.s.User("672835870080106509")
				if err != nil {
					user = e.s.State.User
				}
				return user.AvatarURL("64")
			}(),
		},
	}
}
