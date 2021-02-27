package main

import (
	"io"
	"os"
	"bytes"
	"net"
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

func writeBuffer(s string) bytes.Buffer{
	var buffer bytes.Buffer
	buffer.Write([]byte(s))
	return buffer
}

func connectNetwork(){
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil{
		panic(err)
	}
	// io.Writerとstringを受け取って書き込む
	// io.WriteString(conn, "GET / HTTP/1.0\r\nHost: golang.org\r\n\r\n")
	conn.Write([]byte("GET / HTTP/1.0\r\nHost: golang.org\r\n\r\n"))
	io.Copy(os.Stdout, conn)
}
