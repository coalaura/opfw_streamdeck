package main

import (
	_ "embed"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
	"gitlab.com/milan44/logger-v2"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var (
	hub *SocketHub

	log *logger_v2.Logger

	//go:embed icon.ico
	iconData []byte
)

func main() {
	_, err := CreateMutex("OPFW_STREAMDECK")
	if err != nil {
		panic(err)
	}

	_ = os.RemoveAll("op-fw_streamdeck.log")

	file, err := os.OpenFile("op-fw_streamdeck.log", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	log = logger_v2.New(false, file)

	log.Info("Preparing systray...")
	go systray.Run(onReady, onExit)

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

	_ = file.Close()
}

func onReady() {
	systray.SetIcon(iconData)

	systray.SetTitle("OP-FW Streamdeck")
	systray.SetTooltip("OP-FW Streamdeck")

	mLogs := systray.AddMenuItem("Show Logs", "Opens the most recent log file")

	go func() {
		<-mLogs.ClickedCh

		path, _ := filepath.Abs("op-fw_streamdeck.log")

		_ = open.Start(path)
	}()

	mQuit := systray.AddMenuItem("Quit", "Quit the integration")

	go func() {
		<-mQuit.ClickedCh

		systray.Quit()

		os.Exit(0)
	}()
}

func onExit() {}
