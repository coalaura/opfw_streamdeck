package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Command struct {
	Callback func()
}

type Event struct {
	Name     string
	Commands []Command

	mutex sync.Mutex
}

var (
	userHomeDir string
)

func loadConfig() ([]*Event, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.New("failed to get user home directory")
	}

	userHomeDir = dirname

	path := dirname + "/opfw.yaml"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = ioutil.WriteFile(path, []byte("myEvent:\n  - run me hello\n  - wait 1500\n  - run me bye"), 0777)

		return nil, errors.New("config file not found")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("failed to read config file")
	}

	events := parseConfig(string(data))
	if events == nil {
		return nil, errors.New("failed to parse config file")
	}

	return events, nil
}

func parseConfig(data string) []*Event {
	lines := strings.Split(data, "\n")

	events := make([]*Event, 0)

	var event *Event

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "- ") {
			line = strings.Replace(line, "- ", "", 1)

			cmd := parseCommand(line)
			if cmd == nil {
				log.Warning("Failed to parse command: " + line)

				return nil
			}

			if event == nil {
				log.Warning("No event to add command to: " + line)

				return nil
			}

			event.Commands = append(event.Commands, *cmd)
		} else if strings.HasSuffix(line, ":") {
			if event != nil {
				events = append(events, event)
			}

			eventName := strings.Replace(line, ":", "", 1)

			event = &Event{
				Name:     eventName,
				Commands: make([]Command, 0),
			}
		}
	}

	if event != nil {
		events = append(events, event)
	}

	return events
}

func parseCommand(line string) *Command {
	parts := strings.Split(line, " ")

	command := parts[0]
	data := strings.Join(parts[1:], " ")

	switch command {
	case "run":
		return &Command{
			Callback: func() {
				hub.broadcast <- []byte(data)
			},
		}
	case "wait":
		ms, err := strconv.ParseInt(data, 10, 64)
		if err != nil {
			return nil
		}

		return &Command{
			Callback: func() {
				time.Sleep(time.Duration(ms) * time.Millisecond)
			},
		}
	case "key":
		cb, err := parseKeypressCallback(data)
		if err != nil {
			return nil
		}

		return &Command{
			Callback: cb,
		}
	case "screenshot":
		return &Command{
			Callback: screenshotCallback(data),
		}
	}

	return nil
}
