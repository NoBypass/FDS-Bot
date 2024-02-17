package main

import (
	"fmt"
	"github.com/NoBypass/fds/pkg/api"
	"github.com/fatih/color"
	"github.com/nobypass/fds-bot/internal/bot/commands"
	"github.com/nobypass/fds-bot/internal/bot/core"
	"github.com/nobypass/fds-bot/internal/bot/events"
	"github.com/nobypass/fds-bot/internal/bot/message_components"
	"github.com/nobypass/fds-bot/internal/bot/modals"
	"github.com/nobypass/fds-bot/internal/pkg/version"
	"log"
	"os"
	"os/signal"
)

func init() {
	fmt.Println(`
   _______  ____   ___       __
  / __/ _ \/ __/  / _ )___  / /_
 / _// // /\ \   / _  / _ \/ __/
/_/ /____/___/  /____/\___/\__/   ` + color.New(color.FgMagenta).Sprint(version.VERSION) + `
The FDS Discord bot written in    ` + color.New(color.BgHiCyan).Add(color.FgHiWhite).Sprint(" GO ") + `
________________________________________________
`)
}

func main() {
	session := core.NewSession()
	logger := log.New(os.Stdout, "fds-bot: ", log.Ldate|log.Ltime|log.LstdFlags)

	apiClient := api.NewFDSClient(os.Getenv("API_URL"))
	resp, err := apiClient.BotLogin(&api.DiscordBotLoginRequest{Pwd: os.Getenv("PASSWORD")})
	if err != nil {
		logger.Fatalf("failed to login to bot: %v", err)
	}
	apiClient.SetToken(resp.Token)

	componentManager := message_components.NewManager(logger)
	modalManager := modals.NewManager(logger, apiClient)
	cmdManager := commands.NewManager(logger, apiClient)
	modalManager.RegisterAll()
	componentManager.RegisterAll()
	err = cmdManager.RegisterAll(session)
	if err != nil {
		logger.Fatal(err)
	}

	event := events.New(logger, cmdManager, componentManager, modalManager)
	defer session.Close()

	session.AddHandler(event.OnInteraction)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
