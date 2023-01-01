package main

import (
	"errors"
	"github.com/micmonay/keybd_event"
	"strings"
	"time"
)

func parseKeypressCallback(data string) (func(), error) {
	key := _getKey(data)
	if key == -1 {
		return nil, errors.New("invalid key")
	}

	return func() {
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			log.Warning("Failed to create key bonding: %s", err)

			return
		}

		kb.SetKeys(key)

		err = kb.Press()
		if err != nil {
			log.Warning("Failed to press key (%s): %s", data, err)

			return
		}

		time.Sleep(100 * time.Millisecond)

		err = kb.Release()
		if err != nil {
			log.Warning("Failed to release key (%s): %s", data, err)
		}

		log.Debug("Pressed key: %s", data)
	}, nil
}

func _getKey(data string) int {
	data = strings.ToLower(data)

	switch data {
	case "a":
		return keybd_event.VK_A
	case "b":
		return keybd_event.VK_B
	case "c":
		return keybd_event.VK_C
	case "d":
		return keybd_event.VK_D
	case "e":
		return keybd_event.VK_E
	case "f":
		return keybd_event.VK_F
	case "g":
		return keybd_event.VK_G
	case "h":
		return keybd_event.VK_H
	case "i":
		return keybd_event.VK_I
	case "j":
		return keybd_event.VK_J
	case "k":
		return keybd_event.VK_K
	case "l":
		return keybd_event.VK_L
	case "m":
		return keybd_event.VK_M
	case "n":
		return keybd_event.VK_N
	case "o":
		return keybd_event.VK_O
	case "p":
		return keybd_event.VK_P
	case "q":
		return keybd_event.VK_Q
	case "r":
		return keybd_event.VK_R
	case "s":
		return keybd_event.VK_S
	case "t":
		return keybd_event.VK_T
	case "u":
		return keybd_event.VK_U
	case "v":
		return keybd_event.VK_V
	case "w":
		return keybd_event.VK_W
	case "x":
		return keybd_event.VK_X
	case "y":
		return keybd_event.VK_Y
	case "z":
		return keybd_event.VK_Z

	case "1":
		return keybd_event.VK_1
	case "2":
		return keybd_event.VK_2
	case "3":
		return keybd_event.VK_3
	case "4":
		return keybd_event.VK_4
	case "5":
		return keybd_event.VK_5
	case "6":
		return keybd_event.VK_6
	case "7":
		return keybd_event.VK_7
	case "8":
		return keybd_event.VK_8
	case "9":
		return keybd_event.VK_9
	case "0":
		return keybd_event.VK_0

	case "enter":
		return keybd_event.VK_ENTER
	case "space":
		return keybd_event.VK_SPACE
	case "backspace":
		return keybd_event.VK_BACKSPACE
	case "tab":
		return keybd_event.VK_TAB

	case "f1":
		return keybd_event.VK_F1
	case "f2":
		return keybd_event.VK_F2
	case "f3":
		return keybd_event.VK_F3
	case "f4":
		return keybd_event.VK_F4
	case "f5":
		return keybd_event.VK_F5
	case "f6":
		return keybd_event.VK_F6
	case "f7":
		return keybd_event.VK_F7
	case "f8":
		return keybd_event.VK_F8
	case "f9":
		return keybd_event.VK_F9
	case "f10":
		return keybd_event.VK_F10
	case "f11":
		return keybd_event.VK_F11
	case "f12":
		return keybd_event.VK_F12
	}

	return -1
}
