package types

import (
	"strings"
	"tunnel/config"
)

// stores data that is to be sent back to the public
type PublicResponse struct {
	Data []byte
	ID   string
}

func ParsePublicResponse(rawData []byte) PublicResponse {
	rawDataStr := string(rawData)

	publicResponse := PublicResponse{Data: rawData}
	split := strings.Split(rawDataStr, config.ResponseIdSep)

	if len(split) > 1 {
		publicResponse.ID = split[0]
		publicResponse.Data = rawData[len(publicResponse.ID)+config.ResponseIdSepLen:]
	}

	return publicResponse
}
