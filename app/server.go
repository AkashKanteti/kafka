package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		defer conn.Close()

		req := make([]byte, 128)
		_, err = conn.Read(req)
		if err != nil {
			fmt.Println("Error reading data from connection: ", err.Error())
			os.Exit(1)
		}

		resp := make([]byte, 23)

		//header
		// message size
		binary.BigEndian.PutUint32(resp[0:4], uint32(19))
		// correlation id
		copy(resp[4:8], req[8:12])

		//body
		//error code
		api_version := binary.BigEndian.Uint16(req[6:8])
		if api_version >= 0 && api_version <= 4 {
			binary.BigEndian.PutUint16(resp[8:10], uint16(0))
		} else {
			binary.BigEndian.PutUint16(resp[8:10], uint16(35))
		}

		// array length
		copy(resp[10:11], []byte{2})

		//api key
		copy(resp[11:13], req[4:6])
		//min api version
		binary.BigEndian.PutUint16(resp[13:15], uint16(0))
		//max api version
		copy(resp[15:17], req[6:8])
		//tag buffer
		copy(resp[17:18], []byte{0})
		//throttle time
		binary.BigEndian.PutUint32(resp[18:22], uint32(0))
		//tag buffer
		copy(resp[22:23], []byte{0})

		_, err = conn.Write(resp)
		if err != nil {
			fmt.Println("Error writing data to connection: ", err.Error())
			os.Exit(1)
		}

		defer conn.Close()
	}
}
