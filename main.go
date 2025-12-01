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
		HelloPage(conn)
	} else {
		WorldPage(conn)
	}
	conn.Close() 
}

func HelloPage(conn net.Conn) {
	 render(conn, "assets/hello.html")
}

func WorldPage(conn net.Conn) {
	jsonResponse(conn)
}

func extractHTML(file_path string) (string, error) {
    file, err := os.Open(file_path)
    if err != nil {
        fmt.Println("Error opening file:", err)
    }
    defer file.Close()
    
    content, err := io.ReadAll(file)
    if err != nil {
		fmt.Println("Error reading file:", err)
	}
    
    text := string(content)
	return text, err
}

func render(conn net.Conn, file_path string) {
	html_text, err := extractHTML(file_path)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	response := "HTTP/1.1 200 OK\r\n Content-Type: text/html\r\n\r\n " + html_text
	conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing", err.Error())
		os.Exit(1)
	}
}

func jsonResponse(conn net.Conn) {
	response := "HTTP/1.1 200 OK\r\n Content-Type: application/json\r\n\r\n { 'text': hello }"
	conn.Write([]byte(response))

}
