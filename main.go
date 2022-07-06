package main

import (
	"gitlab.com/milan44/logger-v2"
	"os"
	"os/signal"
	"syscall"
)

var (
	hub *SocketHub

	log = logger_v2.NewColored()
)

func main() {
	log.Info("Starting socket hub...")
	hub = NewSocketHub()

	log.Info("Starting UDP server...")
	StartUDPServer()

	log.Info("Startup complete")

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	<-sigc
	log.Warning("Signal received, shutting down")
}
