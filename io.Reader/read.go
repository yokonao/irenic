package main

import (
	"archive/zip"
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	readBigEndian()
	readPNG("lena.png") // GOは大文字小文字を区別しない？
	embedText("IRENIC PROGRAMMING++", "lena.png", "lenaSecret.png")
	readPNG("lenaSecret.png")
	copyFile()
	genRandomFile()
	zipWriteFile()
	zipWriteString("There is always light behind the clouds.")
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

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("length: %d, chunk: %v\n", length, string(buffer))
	if string(buffer) == "tEXt" {
		textBuffer := make([]byte, length)
		chunk.Read(textBuffer)
		// printlnだと表示されるタイミングがずれる
		fmt.Printf("%s\n", string(textBuffer))
	}
}

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	file.Seek(8, 0)
	var offset int64 = 8
	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
		offset, _ = file.Seek(int64(length)+8, 1)
	}
	return chunks
}

func readPNG(filename string) {
	println(filename)
	file, err := os.Open("../resource/" + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}

func textChunk(text string) bytes.Buffer {
	byteData := []byte(text)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	io.WriteString(&buffer, "tEXt")
	buffer.Write(byteData)
	crc := crc32.NewIEEE() // nanikore
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())
	return buffer
}
func embedText(text string, oldFileName string, newFileName string) {
	file, err := os.Open("../resource/" + oldFileName)
	if err != nil {
		panic(err)
	}
	chunks := readChunks(file)
	defer file.Close()
	newFile, err := os.Create("../resource/" + newFileName)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	io.Copy(newFile, chunks[0])
	buffer := textChunk(text)
	io.Copy(newFile, &buffer)
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}
}

func copyFile() {
	file, err := os.Open("../resource/old.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	io.Copy(newFile, file)
}

func genRandomFile() {
	randFile, err := os.Create("random.txt")
	if err != nil {
		panic(err)
	}
	b := make([]byte, 10)
	_, err = rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	randFile.Write(b)
}

func zipWriteFile() {
	file, err := os.Open("../resource/old.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	body, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	outFile, err := os.Create("old.zip")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("old.txt")
	if err != nil {
		panic(err)
	}
	_, err = writer.Write(body)
	if err != nil {
		panic(err)
	}
}

func zipWriteString(text string) {
	byteData := []byte(text)
	fmt.Println(string(byteData))
	outFile, err := os.Create("quote.zip")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("quote.txt")
	if err != nil {
		panic(err)
	}
	_, err = writer.Write(byteData)
	if err != nil {
		panic(err)
	}

}
