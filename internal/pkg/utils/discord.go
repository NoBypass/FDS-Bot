package utils

import "github.com/bwmarrin/discordgo"

func Username(member *discordgo.Member) string {
	if member.Nick != "" {
		return member.Nick
	}
	return member.User.Username
}
