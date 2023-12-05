package main

import (
	"flag"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/app/handlers"
	"github.com/nobypass/fds-bot/internal/app/lifecycle"
	"log"
	"os"
	"os/signal"
)

var (
	s              *discordgo.Session
	BotToken       = flag.String("token", lifecycle.Config(), "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
	b              = &lifecycle.Bot{
		Token:          BotToken,
		RemoveCommands: RemoveCommands,
	}
)

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	b.Session = s
	log.Println("Session created")

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	log.Println("Session opened")
}

func main() {
	defer s.Close()
	handlers.RegisterCommands(s)

	s.AddHandler(handlers.Ready)
	s.AddHandler(handlers.Commands)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	b.Shutdown()
}
