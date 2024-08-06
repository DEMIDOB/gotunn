package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"tunnel/types"
	"tunnel/util"
)

func ListenToPublic(port int, publicRequestCh chan types.PublicRequest, publicResponseCh chan types.PublicResponse) error {
	logger := log.New(os.Stdout, "[PUBLIC SERVER] ", 0)
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

		data, err := util.ReadFromConnection(conn)
		//logger.Println("Received", string(data))

		publicRequest := types.NewPublicRequest(data)
		publicRequestCh <- publicRequest

		publicResponse := <-publicResponseCh
		if publicResponse.ID != publicRequest.ID {
			log.Panic("Response-Request identifiers do not match!")
		}

		_, err = conn.Write(publicResponse.Data)
		if err != nil {
			logger.Println("Failed to forward the public request to the target:", err)
		}

		err = conn.Close()
		if err != nil {
			return err
		}
	}
}
