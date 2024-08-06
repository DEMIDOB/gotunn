package types

import (
	"github.com/google/uuid"
	"strings"
	"tunnel/config"
)

// PublicRequest stores data received from the public
type PublicRequest struct {
	Data []byte
	ID   string
}

func NewPublicRequest(data []byte) PublicRequest {
	id := uuid.New().String()
	dataWithId := append([]byte(id+config.ResponseIdSep), data...)

	return PublicRequest{Data: dataWithId, ID: id}
}

func ParsePublicRequest(rawData []byte) PublicRequest {
	rawDataStr := string(rawData)

	publicResponse := PublicRequest{Data: rawData}
	split := strings.Split(rawDataStr, config.ResponseIdSep)

	if len(split) > 1 {
		publicResponse.ID = split[0]
		publicResponse.Data = rawData[len(publicResponse.ID)+config.ResponseIdSepLen:]
	}

	return publicResponse
}
