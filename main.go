package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	buffer := make([]byte, 8);
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal("Error reading file:", err)
		}
		fmt.Printf("read: %s\n" ,string(buffer[:n]))
	}
}