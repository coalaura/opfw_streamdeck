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
	log.MustPanic(err)

	go func() {
		for {
			message := make([]byte, 1024)

			length, _, err := conn.ReadFromUDP(message[:])
			if err != nil {
				log.Warning("Failed to read udp message: " + err.Error())

				continue
			}

			data := message[:length]

			log.DebugF("<- %s\n", strings.TrimSpace(string(data)))

			hub.broadcast <- data
		}
	}()
}
