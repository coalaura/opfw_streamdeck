package main

import (
	_ "embed"
	"os"
	"os/signal"
	"syscall"

	logger_v2 "gitlab.com/milan44/logger-v2"
)

var (
	hub *SocketHub

	log *logger_v2.Logger

	info = DebugInfo{}

	config *Config
)

func main() {
	_, err := CreateMutex("OPFW_STREAMDECK")
	if err != nil {
		alert("Startup Error", "OP-FW Streamdeck is already running. You can only run one instance at a time.")

		panic(err)
	}

	_ = os.RemoveAll("op-fw_streamdeck.old.log")
	_ = os.Rename("op-fw_streamdeck.log", "op-fw_streamdeck.old.log")

	file, err := os.OpenFile("op-fw_streamdeck.log", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		alert("Startup Error", "Unable to create log file. Please check your permissions and try again. You may have to check \"unblock\" in the exe's properties.")

		panic(err)
	}

	log = logger_v2.New(false, file)

	reloadConfig()

	initSystray()

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

func reloadConfig() {
	var err error

	config, err = loadConfig()
	if err != nil {
		log.WarningE(err)
	}

	if len(config.Events) > 0 {
		log.InfoF("Loaded %d event listener(s)\n", len(config.Events))
	}
}
