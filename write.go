package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func writeFile(s string) {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte(s))
	file.Close()
}

func writeConsole(s string) {
	os.Stdout.Write([]byte(s))
}

func writeBuffer(s string) bytes.Buffer {
	var buffer bytes.Buffer
	buffer.Write([]byte(s))
	return buffer
}

func writeStringBuilder(s string) strings.Builder {
	var builder strings.Builder
	builder.Write([]byte(s))
	return builder
}

func connectNetwork() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		panic(err)
	}
	// io.Writerとstringを受け取って書き込む
	// io.WriteString(conn, "GET / HTTP/1.0\r\nHost: golang.org\r\n\r\n")
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: golang.org\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}

func connectNetworkByHTTP() {
	res, err := http.Get("http://golang.org")
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, res.Body)
}

func writeJSON() {
	f, e := os.Create("test.json")
	if e != nil {
		panic(e)
	}
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "")
	encoder.Encode(map[string]string{
		"PHP":    "Laravel",
		"Ruby":   "Rails",
		"Python": "Django",
	})
}

func formatForFile() {
	file, err := os.Create("format.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(file, "digit: %d\nstring: %s\nfloat: %f\n", 135, "sample", 4.64)
}
