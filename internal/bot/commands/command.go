package commands

import "github.com/bwmarrin/discordgo"

type Command interface {
	Run(s *discordgo.Session, i *discordgo.InteractionCreate) error
	Content() *discordgo.ApplicationCommand
}

func Register(s *discordgo.Session, c Command) error {
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", c.Content())
	if err != nil {
		return err
	}

	return nil
}

func RegisterAll(s *discordgo.Session) error {
	commands := []Command{
		&Ping{},
	}

	for _, c := range commands {
		err := Register(s, c)
		if err != nil {
			return err
		}
	}

	return nil
}
