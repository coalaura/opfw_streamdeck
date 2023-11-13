package main

import (
	"os"
	"path/filepath"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func startupPath() string {
	return filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs", "Startup", "opfw_streamdeck.lnk")
}

func isAutoStart() bool {
	_, err := os.Stat(startupPath())

	return err == nil
}

func removeFromAutoStart() error {
	return os.Remove(startupPath())
}

func addToAutoStart() error {
	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		return err
	}

	defer oleShellObject.Release()

	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}

	defer wshell.Release()

	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", startupPath())
	if err != nil {
		return err
	}

	idispatch := cs.ToIDispatch()

	defer idispatch.Release()

	oleutil.PutProperty(idispatch, "TargetPath", os.Args[0])
	oleutil.CallMethod(idispatch, "Save")

	return nil
}
