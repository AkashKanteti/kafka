package main

import (
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

	req := []byte{0, 0, 0, 35, 0, 18, 0, 4, 110, 87, 76, 139, 0, 9, 107, 97, 102, 107, 97, 45, 99, 108, 105, 0, 10, 107, 97, 102, 107, 97, 45, 99, 108, 105, 4, 48, 46, 49, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, '\n'}

	// binary.BigEndian.PutUint32(req[0:4], uint32(42))
	// binary.BigEndian.PutUint16(req[4:6], uint16(4))
	// binary.BigEndian.PutUint16(req[6:8], uint16(4))
	// binary.BigEndian.PutUint32(req[8:], uint32(42))
	conn.Write(req)

	resp := make([]byte, 128)
	conn.Read(resp)
	fmt.Printf("%v", resp)
}
