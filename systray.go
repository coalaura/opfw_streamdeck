package main

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

var (
	//go:embed icon.ico
	iconData []byte

	//go:embed icons/info.ico
	infoIcon []byte

	//go:embed icons/lamp_off.ico
	lampOffIcon []byte

	//go:embed icons/lamp.ico
	lampIcon []byte

	//go:embed icons/next.ico
	nextIcon []byte

	//go:embed icons/power.ico
	powerIcon []byte

	//go:embed icons/previous.ico
	previousIcon []byte

	//go:embed icons/reload.ico
	reloadIcon []byte

	systrayExit chan bool
)

func initSystray() {
	log.Info("Preparing systray...")
	go systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(iconData)

	systray.SetTitle("OP-FW Streamdeck")
	systray.SetTooltip("OP-FW Streamdeck")

	addSystrayItem("Debug", "Creates a debug dump", infoIcon, dumpDebugData)

	if isAutoStart() {
		addSystrayItem("Remove Auto-Start", "Remove the application from auto-start.", lampOffIcon, func() {
			err := removeFromAutoStart()

			if err != nil {
				alert("Unable to remove from auto-start. Please try again as administrator.")
			} else {
				alert("Removed from auto-start.")
			}
		})
	} else {
		addSystrayItem("Add Auto-Start", "Add the application to auto-start.", lampIcon, func() {
			err := addToAutoStart()

			if err != nil {
				alert("Unable to add to auto-start. Please try again as administrator.")
			} else {
				alert("Added to auto-start.")
			}
		})
	}

	addSystrayItem("Reload Config", "Reload the configuration file", reloadIcon, reloadConfig)

	addSystrayItem("Open Log-File", "Opens the current log file", nextIcon, func() {
		path, _ := filepath.Abs("op-fw_streamdeck.log")

		_ = open.Start(path)
	})

	addSystrayItem("Open Previous Log-File", "Opens the most recent log file (before the current one)", previousIcon, func() {
		path, _ := filepath.Abs("op-fw_streamdeck.old.log")

		_ = open.Start(path)
	})

	addSystrayItem("Quit", "Quit the integration", powerIcon, func() {
		systray.Quit()

		os.Exit(0)
	})
}

func onExit() {
	systrayExit <- true

	println("Exiting...")
}

func addSystrayItem(name, description string, icon []byte, callback func()) {
	item := systray.AddMenuItem(name, description)

	item.SetIcon(icon)

	go func() {
		for {
			select {
			case <-systrayExit:
				return
			case <-item.ClickedCh:
				callback()
			}
		}
	}()
}
