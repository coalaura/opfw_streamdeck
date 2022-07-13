package main

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/go-ps"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type DebugInfo struct {
	ReceivedCommands int64
	SentCommands     int64

	FailedRead int64
	FailedSend int64

	ActiveConnections int

	LastError error
}

func dumpDebugData() {
	path, _ := filepath.Abs("op-fw_debug.log")

	var buf bytes.Buffer

	buf.WriteString("--- Debug Dump ---\r\n")

	buf.WriteString(fmt.Sprintf("Active: %d\r\n", info.ActiveConnections))
	buf.WriteString("\r\n")

	buf.WriteString(fmt.Sprintf("Received:   %d\r\n", info.ReceivedCommands))
	buf.WriteString(fmt.Sprintf("Sent:       %d\r\n", info.SentCommands))

	buf.WriteString(fmt.Sprintf("FailedRead: %d\r\n", info.FailedRead))
	buf.WriteString(fmt.Sprintf("FailedSent: %d\r\n", info.FailedSend))

	buf.WriteString("\r\n")

	buf.WriteString(FindProcess("SDeck", "StreamDeck.exe"))
	buf.WriteString(FindProcess("FiveM", "FiveM.exe"))
	buf.WriteString(FindProcess("GTA-V", "GTAProcess.exe"))

	buf.WriteString("\r\n")

	if info.LastError != nil {
		buf.WriteString(fmt.Sprintf("LastError: %s", info.LastError.Error()))
	} else {
		buf.WriteString("LastError: N/A")
	}

	_ = ioutil.WriteFile(path, buf.Bytes(), 0777)

	_ = open.Start(path)
}

func FindProcess(label, name string) string {
	pid, name, err := _findProc(name)

	if err != nil {
		return fmt.Sprintf("%s PID: %s\r\n", label, err.Error())
	}

	return fmt.Sprintf("%s PID: %d (%s)\r\n", label, pid, name)
}

func _findProc(name string) (int, string, error) {
	proc, err := ps.Processes()
	if err != nil {
		return 0, "", err
	}

	name = strings.ToLower(name)

	for _, p := range proc {
		if strings.Contains(strings.ToLower(p.Executable()), name) {
			return p.Pid(), p.Executable(), nil
		}
	}

	return 0, "", nil
}
