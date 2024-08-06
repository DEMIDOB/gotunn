package client

import (
	"fmt"
	"log"
	"net"
	"os"
	"tunnel/config"
	"tunnel/types"
	"tunnel/util"
)

func prepareResponseData(response types.PublicResponse) []byte {
	return append([]byte(response.ID+config.ResponseIdSep), response.Data...)
}

func ClientPolling(ip net.IP, serverPort int, targetPort int) error {
	logger := log.New(os.Stdout, "[CLIENT] ", 0)
	logger.Println("Running as client talking to", fmt.Sprintf("%s:%d", ip, serverPort))

	serverAddr := net.TCPAddr{IP: ip, Port: serverPort}
	targetAddr := net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: targetPort}

	var currentRequest types.PublicRequest
	var currentResponse types.PublicResponse

	for {
		logger.Println("Dialing...")
		conn, err := net.DialTCP("tcp", nil, &serverAddr)
		if err != nil {
			return err
		}

		for {
			logger.Println("Sending the message...")

			_, err = conn.Write(prepareResponseData(currentResponse))
			if err != nil {
				return err
			}

			logger.Println("Sent!")

			data, err := util.ReadFromConnection(conn)
			currentRequest = types.ParsePublicRequest(data)
			logger.Println("Received something")

			currentResponse, err = AttackTarget(targetAddr, currentRequest)
			if err != nil {
				return err
			}

			logger.Println("Got a response from the target!")
		}

		err = conn.Close()
		if err != nil {
			return err
		}

	}

	return nil
}
