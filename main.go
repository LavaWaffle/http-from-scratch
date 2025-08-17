package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"strings"
	"bytes"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer f.Close()
		defer close(lines)

		buffer := make([]byte, 8)
		var lineBuilder strings.Builder
		for {
			n, err := f.Read(buffer)
			
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal("Error opening file", err)
			}
			
			data := buffer[:n]
			for {
				i := bytes.IndexByte(data, '\n')
				if i == -1 {
					lineBuilder.Write(data)
					break
				}
				lineBuilder.Write(data[:i])

				lines <- lineBuilder.String()

				lineBuilder.Reset()
				data = data[i+1:]
			}
		}

		if lineBuilder.Len() > 0 {
			lines <- lineBuilder.String()
		}
	}()

	return lines
}

func main() {
    file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	
	for line := range getLinesChannel(file) {
		fmt.Printf("read: %s\n", line)
	}
}