package helpers

import (
	"github.com/nobypass/fds-bot/internal/pkg/discord"
	"log"
)

func Shutdown(s *discord.Session) {
	log.Println("Removing commands...")
	userID := s.State.User.ID

	registeredCommands, err := s.ApplicationCommands(userID, "")
	if err != nil {
		log.Fatalf("Could not fetch registered commands: %v", err)
	}

	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(userID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
		log.Printf("Deleted '%v' command", v.Name)
	}

	log.Println("Gracefully shutting down.")
}
