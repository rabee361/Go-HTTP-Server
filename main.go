package main

import (
	"fmt"
	"net"
	"os"
	"io"
	"bufio"
	"strings"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)


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
		go processRequest(conn)
	}

}

 
func processRequest(conn net.Conn) { 
	reader := bufio.NewReader(conn)
	first_line , err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading", err.Error())
	} else {
		method := extractMethod(first_line)
		path := extractPath(first_line)
		version := extractVersion(first_line)
		fmt.Println("Method: ", method)
		fmt.Println("path", path)
		fmt.Println("HTTP version: ", version)
		sendResponse(conn, path)
	}
} 

func extractMethod(msg string) string {
	return strings.Split(msg, " ")[0]
}

func extractPath(msg string) string {
	return strings.Split(msg, " ")[1]
}

func extractVersion(msg string) string {
	return strings.Split(msg, " ")[2]
}

func sendResponse(conn net.Conn, path string) {

	if path == "/hello" {
		response := "HTTP/1.1 200 OK\r\n Content-Type: text/html\r\n\r\n <h1> " + render("assets/hello.html") + " </h1>"
		conn.Write([]byte(response))
		
	} else {
		response := "HTTP/1.1 200 OK\r\n Content-Type: text/html\r\n\r\n <h1> " + render("assets/world.html") + " </h1>"
		conn.Write([]byte(response))
	}
	conn.Close() 
}


func render(file_path string) string {
    file, err := os.Open(file_path)
    if err != nil {
        fmt.Println("Error opening file:", err)
    }
    defer file.Close() // Ensure file is closed
    
    // Read all content
    content, err := io.ReadAll(file)
    if err != nil {
        fmt.Println("Error reading file:", err)
    }
    
    text := string(content)
	return text
}
