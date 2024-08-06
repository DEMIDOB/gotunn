package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"tunnel/types"
	"tunnel/util"
)

func ListenToTargets(port int, publicRequestCh chan types.PublicRequest, publicResponseCh chan types.PublicResponse) error {
	// listens to target and provides it with data from the public queue when available
	logger := log.New(os.Stdout, "[TARGETS SERVER] ", 0)
	logger.Println("Running as server listening to port", port)

	listenAddr := fmt.Sprintf("0.0.0.0:%d", port)

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		logger.Println("Accepted a new connection!")

		for {
			logger.Println("Waiting for data from the client...")
			data, err := util.ReadFromConnection(conn)
			logger.Println("Got it!")
			//dataStr := string(data)
			//logger.Println("Received", dataStr)

			// publicResponse may or may not be an actual response to the public
			// if there is a non-empty ID, it is an actual response
			// and needs to be placed to the channel
			publicResponse := types.ParsePublicResponse(data)
			if len(publicResponse.ID) > 0 {
				logger.Println("Putting to the channel...")
				publicResponseCh <- publicResponse
			} else {
				logger.Println("No id!")
			}

			logger.Println("Waiting for a public request...")
			// wait for an incoming request to the target
			publicRequest := <-publicRequestCh
			logger.Println("Received a public request")

			// forward the request
			_, err = conn.Write(publicRequest.Data)
			if err != nil {
				logger.Println("Failed to forward the public request to the target:", err)
			}

			logger.Println("Successfully forwarded the request!")
		}

		err = conn.Close()
		if err != nil {
			return err
		}
	}
}
