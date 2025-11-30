package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

type Request struct {

}

func main() {
	
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	fmt.Print("Server runnin on port: " , SERVER_PORT)
	if err != nil {
		fmt.Println("Error dialing", err.Error()) // if connection failed
		os.Exit(1)
	}	

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error()) // if accept failed
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(conn)
	}

}

 
func processClient(conn net.Conn) { 
	reader := bufio.NewReader(conn)
	msg , err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}
	request_slice := strings.Split(msg, " / ")
	fmt.Println("Recieved", string(msg))
	fmt.Println("Method: ", request_slice[0])
	fmt.Println("HTTP version: ", request_slice)
	sendResponse(conn)
	// conn.Close()
} 

func sendResponse(conn net.Conn) {
	response := "HTTP/1.1 200 OK\r\n Content-Type: text/html\r\n\r\n <h1> Hello World </h1>"
	conn.Write([]byte(response))
	conn.Close()
}
