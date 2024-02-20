package components

import (
	"fmt"
	"github.com/NoBypass/fds/pkg/api"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/utils"
	"github.com/nobypass/fds-bot/internal/pkg/version"
	"time"
)

var (
	EmbedLeaderboard = func(s *discordgo.Session, lb *api.DiscordLeaderboardResponse, page int) *discordgo.MessageEmbed {
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

	EmbedProfile = func(dcMember *discordgo.Member, member *api.DiscordMemberResponse, url string) *discordgo.MessageEmbed {
		return &discordgo.MessageEmbed{
			Title: fmt.Sprintf("Profile of %s", utils.Username(dcMember)),
			Color: 0x2B2D31,
			Fields: func() []*discordgo.MessageEmbedField {
				dailyAt := utils.StrToTime(member.LastDailyAt)
				return []*discordgo.MessageEmbedField{
					{
						Name:  "Last daily was claimed at",
						Value: fmt.Sprintf("`%s`\n(`%s` ago)", dailyAt.Format("2006-01-02 15:04:05"), utils.StrAgo(dailyAt)),
					},
					{
						Name:  "Streak",
						Value: fmt.Sprintf("Current Streak `%d`\nStarted at `%s`\n Best Streak `TODO`", member.Streak, time.Now().Add(-time.Duration(member.Streak)*24*time.Hour).Format("2006-01-02")),
					},
				}
			}(),
			Image: &discordgo.MessageEmbedImage{
				URL: url,
			},
		}
	}

	EmbedVerificationDone = func(resp api.DiscordVerifyResponse) *discordgo.MessageEmbed {
		return &discordgo.MessageEmbed{
			Title:       "You are verified!",
			Color:       0x2B2D31,
			Description: fmt.Sprintf("This discord account has been linked to `%v` via Hypixel.\n\nInfo: you will soon not be able to see this channel anymore.", resp.Actual),
		}
	}

	EmbedTeams = func(teams [][]string) *discordgo.MessageEmbed {
		return &discordgo.MessageEmbed{
			Title: "Teams",
			Color: 0x2B2D31,
			Fields: func() []*discordgo.MessageEmbedField {
				var fields []*discordgo.MessageEmbedField
				for i, team := range teams {
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

	EmbedPlay = func(member *discordgo.Member, to *discordgo.Role, desc *discordgo.ApplicationCommandInteractionDataOption, mode string) *discordgo.MessageEmbed {
		return &discordgo.MessageEmbed{
			Title: fmt.Sprintf("%s is looking to play %s", member.User.Username, mode),
			Description: func() string {
				if desc != nil {
					return fmt.Sprintf("**Description**\n%v", desc.Value)
				}
				return ""
			}(),
			Color: 0x2B2D31,
			Author: &discordgo.MessageEmbedAuthor{
				Name:    member.User.Username,
				IconURL: member.User.AvatarURL("64"),
			},
		}
	}

	EmbedVerify = &discordgo.MessageEmbed{
		Title: "Verify",
		Color: 0x2B2D31,
		Description: `Verify your Discord account by linking it to Hypixel.

If it does not work, it might be because you have not linked your Discord account to your Hypixel account. [This video](https://www.youtube.com/watch?v=UresIQdoQHk) will show you how to do it.

**THIS STEP IS REQUIRED TO ACCESS THE SERVER**`,
	}

	EmbedHelp = func(s *discordgo.Session, cmds map[string]*discordgo.ApplicationCommand) *discordgo.MessageEmbed {
		return &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("FDS Bot %s", version.VERSION),
			Description: "List of available commands",
			Color:       0x00ff00,
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
)
