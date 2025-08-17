package main

import (
	"fmt"
	"log"
	"io"
	"strings"
	"bytes"
	"net"
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
    listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Error listening to 42069: ", err)
	}
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err)
		}

		for line := range getLinesChannel(conn) {
			fmt.Printf("read: %s\n", line)
		}
	}
}