package components

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/pkg/version"
)

var (
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
