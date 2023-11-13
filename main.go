package main

import (
	_ "embed"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
	logger_v2 "gitlab.com/milan44/logger-v2"
)

var (
	hub *SocketHub

	log *logger_v2.Logger

	info = DebugInfo{}

	config *Config

	//go:embed icon.ico
	iconData []byte

	//go:embed icons/debug.ico
	debugIcon []byte

	//go:embed icons/config.ico
	configIcon []byte

	//go:embed icons/quit.ico
	quitIcon []byte

	//go:embed icons/oldLogs.ico
	oldLogsIcon []byte

	//go:embed icons/currentLogs.ico
	currentLogsIcon []byte
)

func main() {
	_, err := CreateMutex("OPFW_STREAMDECK")
	if err != nil {
		alert("OP-FW Streamdeck is already running. You can only run one instance at a time.")

		panic(err)
	}

	_ = os.RemoveAll("op-fw_streamdeck.old.log")
	_ = os.Rename("op-fw_streamdeck.log", "op-fw_streamdeck.old.log")

	file, err := os.OpenFile("op-fw_streamdeck.log", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		alert("Unable to create log file. Please check your permissions and try again. You may have to check \"unblock\" in the exe's properties.")

		panic(err)
	}

	log = logger_v2.New(false, file)

	reloadConfig()

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

func onReady() {
	systray.SetIcon(iconData)

	systray.SetTitle("OP-FW Streamdeck")
	systray.SetTooltip("OP-FW Streamdeck")

	mDebug := systray.AddMenuItem("Debug", "Creates a debug dump")

	mDebug.SetIcon(debugIcon)

	go func() {
		for {
			<-mDebug.ClickedCh

			dumpDebugData()
		}
	}()

	mConfig := systray.AddMenuItem("Reload Config", "Reload the configuration file")

	mConfig.SetIcon(configIcon)

	go func() {
		<-mConfig.ClickedCh

		log.Info("Reloading config...")

		reloadConfig()
	}()

	mLogs := systray.AddMenuItem("Show current Logs", "Opens the current log file")

	mLogs.SetIcon(currentLogsIcon)

	go func() {
		for {
			<-mLogs.ClickedCh

			path, _ := filepath.Abs("op-fw_streamdeck.log")

			_ = open.Start(path)
		}
	}()

	mOldLogs := systray.AddMenuItem("Show previous Logs", "Opens the most recent log file")

	mOldLogs.SetIcon(oldLogsIcon)

	go func() {
		for {
			<-mOldLogs.ClickedCh

			path, _ := filepath.Abs("op-fw_streamdeck.old.log")

			_ = open.Start(path)
		}
	}()

	mQuit := systray.AddMenuItem("Quit", "Quit the integration")

	mQuit.SetIcon(quitIcon)

	go func() {
		<-mQuit.ClickedCh

		systray.Quit()

		os.Exit(0)
	}()
}

func onExit() {}
