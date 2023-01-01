package main

import (
	"github.com/vova616/screenshot"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func screenshotCallback(data string) func() {
	path := userHomeDir

	if data != "" {
		path = data
	}

	return func() {
		file := filepath.Join(path, "opfw_"+strconv.FormatInt(time.Now().UnixNano(), 16)+".png")

		img, err := screenshot.CaptureScreen()
		if err != nil {
			log.Warning("failed to capture screen: " + err.Error())

			return
		}

		f, err := os.Create(file)
		if err != nil {
			log.Warning("failed to create screenshot file: " + err.Error())

			return
		}

		_ = png.Encode(f, img)
		_ = f.Close()
	}
}
