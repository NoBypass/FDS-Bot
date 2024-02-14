package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/nobypass/fds-bot/internal/bot/commands"
	"github.com/nobypass/fds-bot/internal/bot/core"
	"github.com/nobypass/fds-bot/internal/bot/events"
	"log"
	"os"
	"os/signal"
)

func init() {
	fmt.Println(`
   _______  ____   ___       __
  / __/ _ \/ __/  / _ )___  / /_
 / _// // /\ \   / _  / _ \/ __/
/_/ /____/___/  /____/\___/\__/   ` + color.New(color.FgMagenta).Sprint(VERSION) + `
The FDS Discord bot written in    ` + color.New(color.BgHiCyan).Add(color.FgHiWhite).Sprint(" GO ") + `
________________________________________________
`)
}

const VERSION = "v3.4.0"

func main() {
	session := core.NewSession()
	logger := log.New(os.Stdout, "fds-bot: ", log.Ldate|log.Ltime|log.LstdFlags)
	event := events.New(logger)
	defer session.Close()

	err := commands.RegisterAll(session)
	if err != nil {
		logger.Println(err)
		return
	}

	session.AddHandler(event.OnInteraction)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
