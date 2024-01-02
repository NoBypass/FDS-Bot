package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/nobypass/fds-bot/internal/app/handlers"
	"github.com/nobypass/fds-bot/internal/app/lifecycle"
	"github.com/nobypass/fds-bot/internal/pkg/consts"
	"log"
	"os"
	"os/signal"
)

var (
	s              *discordgo.Session
	BotToken       = os.Getenv("token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
	b              = &lifecycle.Bot{
		Token:          BotToken,
		RemoveCommands: RemoveCommands,
	}
)

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + BotToken)
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

const VERSION = "v3.1.0"

func main() {
	defer s.Close()

	fmt.Println(`
   _______  ____   ___       __
  / __/ _ \/ __/  / _ )___  / /_
 / _// // /\ \   / _  / _ \/ __/
/_/ /____/___/  /____/\___/\__/   ` + consts.Purple.Sprint(VERSION) + `
The FDS Discord bot written in    ` + consts.WhiteOnCyan.Sprint(" GO ") + `
________________________________________________
`)

	handlers.RegisterCommands(s)

	s.AddHandler(handlers.Ready)
	s.AddHandler(handlers.Interactions)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	b.Shutdown()
}
