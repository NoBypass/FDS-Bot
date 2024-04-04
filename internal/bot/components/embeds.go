package components

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/bot/model"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/nobypass/fds-bot/internal/pkg/version"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"time"
)

func EmbedRevoked(member *model.MemberResponse) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Verification of %s revoked", member.Nick),
		Description: fmt.Sprintf("Previous data: %+v", member),
	}
}

func EmbedProfile(member *model.MemberResponse, url string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Profile of %s", member.Nick),
		Color: 0x2B2D31,
		Fields: func() []*discordgo.MessageEmbedField {
			startDate := time.Now().Add(-time.Duration(member.Streak) * 24 * time.Hour).Format("2006-01-02")
			dailyAt, err := time.Parse(time.RFC3339, member.LastDailyAt)
			if err != nil {
				dailyAt = time.Time{}
			}
			if member.Streak == 0 {
				startDate = "Never"
			}

			return []*discordgo.MessageEmbedField{
				{
					Name:  "Last daily was claimed at",
					Value: fmt.Sprintf("`%s`\n(`%s` ago)", dailyAt.Format("2006-01-02 15:04:05"), utils.StrAgo(dailyAt)),
				},
				{
					Name:  "Streak",
					Value: fmt.Sprintf("Current Streak `%d`\nStarted at `%s`\n Best Streak `TODO`", member.Streak, startDate),
				},
			}
		}(),
		Image: &discordgo.MessageEmbedImage{
			URL: url,
		},
	}
}

func EmbedVerificationDone(resp *model.VerifyResponse) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "You are verified!",
		Color: 0x2B2D31,
		Description: fmt.Sprintf("This discord account has been linked to `%v` via Hypixel.\n\n"+
			"Info: you will soon not be able to see this channel anymore.", resp.Actual),
	}
}

func EmbedTeams(teams [][]string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Teams",
		Color: 0x2B2D31,
		Fields: func() []*discordgo.MessageEmbedField {
			var fields []*discordgo.MessageEmbedField
			for i, team := range teams {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name: fmt.Sprintf("Team %d", i+1),
					Value: func() (r string) {
						for _, player := range team {
							r += fmt.Sprintf("%s\n", player)
						}
						return
					}(),
					Inline: true,
				})
			}
			return fields
		}(),
	}
}

func EmbedPlay(member *discordgo.Member, desc *discordgo.ApplicationCommandInteractionDataOption, mode string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s is looking to play %s", member.Nick, mode),
		Description: func() string {
			if desc != nil {
				return fmt.Sprintf("**Description**\n%v", desc.Value)
			}
			return ""
		}(),
		Color: 0x2B2D31,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    member.Nick,
			IconURL: member.User.AvatarURL("64"),
		},
	}
}

func EmbedVerify() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Verify",
		Color: 0x2B2D31,
		Description: `Verify your Discord account by linking it to Hypixel.

If it does not work, it might be because you have not linked your Discord account to your Hypixel account. [This video](https://www.youtube.com/watch?v=UresIQdoQHk) will show you how to do it.

**THIS STEP IS REQUIRED TO ACCESS THE SERVER**`,
	}
}

func EmbedHelp(s *discordgo.Session, cmds map[string]*discordgo.ApplicationCommand) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("FDS Bot %s", version.VERSION),
		Description: "List of available commands",
		Color:       0x2B2D31,
		Fields: func() []*discordgo.MessageEmbedField {
			fields := make([]*discordgo.MessageEmbedField, 0, len(cmds))
			for name, cmd := range cmds {
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
				user, err := s.User("672835870080106509")
				if err != nil {
					user = s.State.User
				}
				return user.AvatarURL("64")
			}(),
		},
	}
}

func EmbedLeaderboard(s *discordgo.Session, lb *model.LeaderboardResponse, page int) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Leaderboard (Page %v)", page+1),
		Color: 0x2B2D31,
		Description: func() (final string) {
			for i, player := range *lb {
				user, err := s.User(player.DiscordID)
				if err != nil {
					user = &discordgo.User{ID: "Unknown"}
				}
				final += fmt.Sprintf("%d. %s - Level: %d | XP: %f\n", i+1, user.Mention(), player.Level, player.XP)
			}
			return
		}(),
	}
}

func EmbedDaily(before, after *model.MemberResponse) *discordgo.MessageEmbed {
	xpDiff := after.XP - before.XP
	lvlDiff := after.Level - before.Level
	var lvlup string
	if lvlDiff > 0 {
		xpDiff += before.GetNeededXP() - before.XP
		lvlup = fmt.Sprintf("**You have leveled up to %d!**\n", after.Level)
	}

	return &discordgo.MessageEmbed{
		Title: fmt.Sprintf("Daily reward for %s", before.Nick),
		Color: 0x2B2D31,
		Description: fmt.Sprintf("You have claimed your daily reward!\n"+
			"%s\nReveived **%d**xp\nCurrent streak: **%d**\n"+
			"Need **%d** for next level\nLevel **%d**", lvlup, int(xpDiff), after.Streak, int(after.GetNeededXP()-after.XP), after.Level),
	}
}

func EmbedError(err error, sp opentracing.Span) *discordgo.MessageEmbed {
	traceID := sp.Context().(jaeger.SpanContext).TraceID().String()
	return &discordgo.MessageEmbed{
		Title:       "An error occurred",
		Color:       0x2B2D31,
		Description: fmt.Sprintf("## `%v`\n\n\nIf you think that this is not intended behaviour, please send this ID `%s` to an admin", err, traceID),
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("TraceID: %s", traceID),
		},
	}
}
