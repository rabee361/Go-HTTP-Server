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
	fmt.Println("Server runnin on port: " , SERVER_PORT)
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		os.Exit(1)
	}	

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			os.Exit(1)
		}
		go processRequest(conn)
	}

}

 
func processRequest(conn net.Conn) { 
	reader := bufio.NewReader(conn)
	first_line , err := reader.ReadString('\n')
	if err != nil {
		return
	} else {
		method := extractMethod(first_line)
		path := extractPath(first_line)
		version := extractVersion(first_line)
		fmt.Print(method, " ")
		fmt.Print(path, " ")
		fmt.Print(version, " ")
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
		status := HelloPage(conn)
		fmt.Println(status)
	} else if path == "/world"{
		status := WorldPage(conn)
		fmt.Println(status)
	} else {
		status := Handle404(conn)
		fmt.Println(status)
	}
	conn.Close() 
}

func HelloPage(conn net.Conn) int {
	response := "HTTP/1.1 200 OK\r\n Content-Type: text/html\r\n\r\n <h1> " + render("assets/hello.html")
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing", err.Error()) // if write failed
		os.Exit(1)
	}
	return 200
}

func WorldPage(conn net.Conn) int {
	response := "HTTP/1.1 200 OK\r\n Content-Type: text/html\r\n\r\n <h1> " + render("assets/world.html")
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing", err.Error()) // if write failed
		os.Exit(1)
	}
	return 200
}

func Handle404(conn net.Conn) int {
	response := "HTTP/1.1 404 OK\r\n Content-Type: text/html"
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing", err.Error()) // if write failed
		os.Exit(1)
	}
	return 404
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
