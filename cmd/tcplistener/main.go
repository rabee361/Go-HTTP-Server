package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)


func main() {
	listener, err := startServer()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server Started on port localhost:42069...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()
	lines := getLinesChannel(conn)
	for line := range lines {
		fmt.Println(line)
	}
}


func startServer() (net.Listener, error) {
	listener, err := net.Listen("tcp", "localhost:42069")
	return listener, err
}

func getLinesChannel(conn net.Conn) <-chan string {
	reader := bufio.NewReader(conn)
	ch := make(chan string)

	go func() {
		for {
			b, err := reader.ReadBytes('\n')
			if err == io.EOF {
				fmt.Println("End of file")
				close(ch)
				break
			}
			ch <- string(b)
		}
	}()
	return ch
}
