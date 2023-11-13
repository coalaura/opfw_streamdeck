package main

import (
	"net"
	"strings"
)

func StartUDPServer() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: 42069,
		IP:   net.ParseIP("127.0.0.1"),
	})
	if err != nil {
		alert("Failed to start UDP server. Make sure no other application is using port 42069. You may have to check \"unblock\" in the exe's properties.")

		log.MustPanic(err)
	}

	go func() {
		for {
			message := make([]byte, 1024)

			length, _, err := conn.ReadFromUDP(message[:])
			if err != nil {
				log.Warning("Failed to read udp message: " + err.Error())

				info.LastError = err

				info.FailedRead++

				continue
			}

			data := message[:length]

			info.ReceivedCommands++

			log.DebugF("<- %s\n", strings.TrimSpace(string(data)))

			hub.broadcast <- data
		}
	}()
}
