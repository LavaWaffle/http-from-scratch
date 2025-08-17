package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"strings"
	"bytes"
)

func main() {
    file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	buffer := make([]byte, 8);
	var currentLine strings.Builder
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Error reading file:", err)
		}

		data := buffer[:n]
		for {
			i := bytes.IndexByte(data, '\n')
			if i == -1 {
				currentLine.Write(data)
				break
			}
			currentLine.Write(data[:i])

			fmt.Printf("read: %s\n", currentLine.String())
			currentLine.Reset()
			data = data[i+1:]
		}
	}

	if currentLine.Len() > 0 {
		fmt.Printf("read: %s\n", currentLine.String())
	}
}