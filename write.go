package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/csv"
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

func writeBufio() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	writeConsole("hello\n")
	buffer.WriteString("example\n")
	buffer.Flush()
}

func writeBufio2() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.WriteString("example\n")
	buffer.Flush()
}

func writeBufio3() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("2start ")
	writeBufio2()
	buffer.WriteString("2end\n")
	buffer.Flush()
}

func writeGzip() {
	file, err := os.Create("gzip.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "gzip.txt"
	io.WriteString(writer, "gzip.Writer example\n")
	writer.Close()
}

func writeCsv() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{"Nagoya", "052"})
	writer.Write([]string{"Tokyo", "03"})
	writer.Write([]string{"Kyoto", "075"})
	writer.Flush()

}

func writeCsvStdout() {
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"Nagoya", "052"})
	writer.Write([]string{"Tokyo", "03"})
	writer.Write([]string{"Kyoto", "075"})
	writer.Flush()

}

func writeCsvMulti() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(file, os.Stdout)
	writer := csv.NewWriter(multiWriter)
	writer.Write([]string{"Nagoya", "052"})
	writer.Write([]string{"Tokyo", "03"})
	writer.Write([]string{"Kyoto", "075"})
	writer.Flush()

}
