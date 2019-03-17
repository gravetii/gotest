package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	bufsize = 1024 * 512
)

func main() {
	f, _ := os.Open("NewFile")
	//fmt.Println(err)
	read(f)
}

func write(file *os.File) int {
	buf := bufio.NewWriter(file)
	b, _ := buf.WriteString("Writing data to buffer!\n")
	buf.Flush()
	file.Close()
	return b
}

func read(f *os.File) {
	fmt.Println("Reading file now...")
	p := make([]byte, bufsize)
	reader := bufio.NewReader(f)
	reader.Read(p)
	str := string(p)
	fmt.Println(str)
}

func append(file *os.File) {
	file.WriteString("Trying to append!\n")
	file.Close()
}
