package main

import (
	"github.com/NoBypass/mincache"
	"github.com/fatih/color"
	"github.com/nobypass/fds-bot/internal/bot/event"
	"github.com/nobypass/fds-bot/internal/bot/interaction/btns"
	"github.com/nobypass/fds-bot/internal/bot/interaction/cmds"
	"github.com/nobypass/fds-bot/internal/bot/interaction/modals"
	"github.com/nobypass/fds-bot/internal/bot/session"
	"github.com/nobypass/fds-bot/internal/monitoring"
	"github.com/nobypass/fds-bot/internal/pkg/version"
	"os"
	"os/signal"
	"slices"
)

func init() {
	println(`
   _______  ____   ___       __
  / __/ _ \/ __/  / _ )___  / /_
 / _// // /\ \   / _  / _ \/ __/
/_/ /____/___/  /____/\___/\__/   ` + color.New(color.FgMagenta).Sprint(version.VERSION) + `
The FDS Discord bot written in    ` + color.New(color.BgHiCyan).Add(color.FgHiWhite).Sprint(" GO ") + `
________________________________________________
`)
}

func main() {
	s := session.ConnectToDiscord()
	tracer, closer := monitoring.CreateTracer()
	fds := session.ConnectToFDS(tracer)
	c := mincache.New()
	em := event.NewManager(s, tracer, c)
	defer closer.Close()
	defer s.Close()

	modals := modals.AllModals(fds)
	buttons := btns.AllButtons(fds)
	commands := cmds.AllCommands(fds)

	em.Add(slices.Concat(modals, buttons, commands)...)

	s.AddHandler(em.Handle)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	em.Remove()
}
