package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to connect to port 9092")
		os.Exit(1)
	}
	defer conn.Close()

	req := make([]byte, 12)

	binary.BigEndian.PutUint32(req[0:4], uint32(42))
	binary.BigEndian.PutUint16(req[4:6], uint16(4))
	binary.BigEndian.PutUint16(req[6:8], uint16(4))
	binary.BigEndian.PutUint32(req[8:], uint32(42))
	conn.Write(req)

	resp := make([]byte, 8)
	conn.Read(resp)
	fmt.Printf("%v", resp)
}
