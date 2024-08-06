package main

import (
	"flag"
	"log"
	"net"
	"tunnel/client"
	"tunnel/server"
	"tunnel/types"
)

func main() {
	isServerPtr := flag.Bool("s", false, "Whether to run in server mode")
	port := flag.Int("p", 0, "Port to accept targets in server mode")
	flag.Parse()

	isServer := *isServerPtr

	if isServer {
		publicRequestCh := make(chan types.PublicRequest)
		publicResponseCh := make(chan types.PublicResponse)

		go func() {
			err := server.ListenToTargets(*port, publicRequestCh, publicResponseCh)
			if err != nil {
				log.Panic(err)
			}
		}()

		err := server.ListenToPublic(*port+1, publicRequestCh, publicResponseCh)
		if err != nil {
			log.Panic(err)
		}
	} else {
		// temporarily assume there is only one server with only one port to listen
		ip := net.IPv4(109, 201, 65, 61)
		//ip := net.IPv4(10, 8, 0, 4)
		//ip := net.IPv4(127, 0, 0, 1)

		serverPort := 2004
		targetPort := *port

		err := client.ClientPolling(ip, serverPort, targetPort)
		if err != nil {
			log.Panic(err)
		}
	}
}
