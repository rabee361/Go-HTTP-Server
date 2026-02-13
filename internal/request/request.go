package request

import (
	"io"
	"strings"
	"fmt"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}


func RequestFromReader(reader io.Reader) (*Request, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	request := Request{}
	request.RequestLine, err = parseRequestLine(string(bytes))
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func parseRequestLine(line string) (RequestLine, error) {
	parts := strings.Split(line, "\r\n")
	req_line := parts[0]
	req_parts := strings.Split(req_line, " ")

	method := req_parts[0]
	target := req_parts[1]
	version := req_parts[2]

	if len(req_parts) != 3 {
		return RequestLine{}, fmt.Errorf("invalid request line: %s", req_line)
	}
	if strings.ToUpper(method) != method {
		return RequestLine{}, fmt.Errorf("invalid request line: %s", req_line)
	}
	if !strings.HasPrefix(version, "HTTP/") {
		return RequestLine{}, fmt.Errorf("invalid request line: %s", req_line)
	}
	if !strings.HasPrefix(target, "/") {
		return RequestLine{}, fmt.Errorf("invalid request line: %s", req_line)
	}

	http_version, err := extractHTTPVersion(version)
	if err != nil {
		return RequestLine{}, err
	}

	return RequestLine{HttpVersion: http_version, RequestTarget: target, Method: method}, nil	
}

func extractHTTPVersion(part string) (string, error) {
	if !strings.HasPrefix(part, "HTTP/") {
		return "", fmt.Errorf("invalid http version: %s", part)
	}
	version := strings.Split(part, "/")
	if len(version) != 2 {
		return "", fmt.Errorf("invalid http version: %s", part)
	}
	return version[1], nil
}