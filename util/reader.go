package util

import (
	"net"
	"time"
)

const (
	BufferSize  = 1
	ReadTimeout = time.Millisecond
)

func ReadFromConnection(conn net.Conn) ([]byte, error) {
	data := make([]byte, 0, BufferSize)

	buff := make([]byte, BufferSize)
	bytesRead, err := conn.Read(buff)
	if err != nil {
		return data, err
	}

	for bytesRead > 0 && bytesRead == BufferSize {
		if bytesRead < 1 {
			break
		}

		data = append(data, buff...)
		buff = make([]byte, BufferSize)
		err := conn.SetReadDeadline(time.Now().Add(ReadTimeout))
		if err != nil {
			return nil, err
		}

		bytesRead, _ = conn.Read(buff)
	}

	data = append(data, buff...)

	return data, nil
}
