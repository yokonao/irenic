package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	writeFile("Test for file writing\n")
	writeConsole("Test for console writing\n")
	var buf = writeBuffer("Test for buffer writing\n")
	writeConsole(buf.String())
	var builder = writeStringBuilder("Test for string builder writing\n")
	writeConsole(builder.String())
	connectNetwork()
	connectNetworkByHTTP()
}
