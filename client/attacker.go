package client

import (
	"net"
	"tunnel/types"
	"tunnel/util"
)

func AttackTarget(targetAddr net.TCPAddr, request types.PublicRequest) (types.PublicResponse, error) {
	response := types.PublicResponse{Data: make([]byte, 0), ID: request.ID}

	conn, err := net.DialTCP("tcp", nil, &targetAddr)
	if err != nil {
		return response, err
	}
	defer conn.Close()

	_, err = conn.Write(request.Data)
	if err != nil {
		return response, err
	}

	data, err := util.ReadFromConnection(conn)

	response.Data = data
	return response, nil
}
