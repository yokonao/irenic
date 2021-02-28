package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	readBigEndian()
}

func readBigEndian() {
	data := []byte{0x00, 0x00, 0x27, 0x10}
	var i int32
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
	// 試しにリトルエンディアン
	binary.Read(bytes.NewReader(data), binary.LittleEndian, &i)
	fmt.Printf("data: %d\n", i)
}
