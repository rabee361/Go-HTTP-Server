package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	address , err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, address)
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println(">")
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("EOF")
			return
		}
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		conn.Write([]byte(line))
	}
}