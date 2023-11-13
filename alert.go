package main

import (
	"os"
	"path/filepath"

	"gopkg.in/toast.v1"
)

func alert(title, msg string) {
	icon := filepath.Join(os.TempDir(), "opfw_streamdeck.ico")

	os.WriteFile(icon, iconData, 0644)

	notification := toast.Notification{
		AppID:   "OP-FW Streamdeck",
		Title:   title,
		Message: msg,
		Icon:    icon,
	}

	notification.Push()
}
