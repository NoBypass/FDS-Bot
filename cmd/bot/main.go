package main

import (
	"github.com/NoBypass/mincache"
	"github.com/fatih/color"
	"github.com/labstack/gommon/log"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/interaction/btns"
	"github.com/nobypass/fds-bot/internal/bot/interaction/cmds"
	"github.com/nobypass/fds-bot/internal/bot/interaction/mdls"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/nobypass/fds-bot/internal/monitoring"
	"github.com/nobypass/fds-bot/internal/pkg/version"
	"os"
	"os/signal"
	"slices"
)

func main() {
	logger := log.New("bot")
	logger.SetLevel(log.INFO)
	logger.Print(`:
   _______  ____   ___       __
  / __/ _ \/ __/  / _ )___  / /_
 / _// // /\ \   / _  / _ \/ __/
/_/ /____/___/  /____/\___/\__/   ` + color.New(color.FgMagenta).Sprint(version.VERSION) + `
The FDS Discord bot written in    ` + color.New(color.BgHiCyan).Add(color.FgHiWhite).Sprint(" GO ") + `
________________________________________________
`)

	tracer, closer := monitoring.CreateTracer()
	defer closer.Close()
	logger.Info("✓ Connected to SurrealDB")

	fds := session.ConnectToFDS(tracer)
	logger.Info("✓ Logged in to FDS API")

	s := session.ConnectToDiscord()
	defer s.Close()
	logger.Info("✓ Connected to Discord")

	c := mincache.New()
	logger.Info("✓ Started cache")

	em := event.NewManager(s, tracer, c, logger)

	modals := mdls.AllModals(fds)
	buttons := btns.AllButtons(fds)
	commands := cmds.AllCommands(fds)

	em.Add(slices.Concat(modals, buttons, commands)...)

	s.AddHandler(em.Handle)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	em.Remove()
}
