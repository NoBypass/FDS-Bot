package discord

import "log"

func (s *Session) RegisterCommand(command *Command) {
	s.commands[command.Name] = command
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", command.ApplicationCommand)
	if err != nil {
		panic(err)
	}
	log.Println("Registered command:", command.Name)
}

func (s *Session) RegisterInteraction(name string, handler interactionCreateFunc) {
	s.interactions[name] = handler
}
