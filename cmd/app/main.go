package main

import (
	"fmt"
	"github.com/nobypass/fds-bot/internal/app/cmds"
	"github.com/nobypass/fds-bot/internal/app/interactions"
	"github.com/nobypass/fds-bot/internal/pkg/consts"
	"github.com/nobypass/fds-bot/internal/pkg/discord"
	"github.com/nobypass/fds-bot/internal/pkg/helpers"
	"log"
	"os"
	"os/signal"
)

var s *discord.Session

func init() {
	fmt.Println(`
   _______  ____   ___       __
  / __/ _ \/ __/  / _ )___  / /_
 / _// // /\ \   / _  / _ \/ __/
/_/ /____/___/  /____/\___/\__/   ` + consts.Purple.Sprint(VERSION) + `
The FDS Discord bot written in    ` + consts.WhiteOnCyan.Sprint(" GO ") + `
________________________________________________
`)
}

func init() {
	var err error
	s, err = discord.Connect(os.Getenv("token"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Session created")
}

const VERSION = "v3.2.1"

func main() {
	defer s.Close()
	defer helpers.Shutdown(s)

	s.RegisterCommand(cmds.Admin)
	s.RegisterCommand(cmds.Ping)
	s.RegisterCommand(cmds.Daily)
	s.RegisterCommand(cmds.Admin)
	s.RegisterCommand(cmds.VCTeams)
	s.RegisterCommand(cmds.Play)

	s.RegisterInteraction("verify", interactions.VerifyHandler)
	s.RegisterInteraction("verify_modal_submit", interactions.VerifyModalSubmitHandler)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
