package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("NewFile")
	if err != nil {
		panic(err)
	}

	b := write(file)
	log.Printf("Finished writing to file: %d", b)

	rf, _ := os.Open("NewFile")
	read(rf)

	af, _ := os.OpenFile("NewFile", os.O_APPEND|os.O_WRONLY, 0600)
	append(af)
}

func write(file *os.File) int {
	buf := bufio.NewWriter(file)
	b, _ := buf.WriteString("Writing data to buffer!\n")
	buf.Flush()
	file.Close()
	return b
}

func read(file *os.File) {
	reader := bufio.NewReader(file)
	data, err := reader.Peek(50)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println(string(data))
}

func append(file *os.File) {
	file.WriteString("Trying to append!\n")
	file.Close()
}
