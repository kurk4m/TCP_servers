package main

import (
	"fmt"
	"log"
	"os"
)

type ioRead struct{}

func (d *ioRead) Read(b []byte) (int, error) {
	fmt.Println("in> ")
	return os.Stdin.Read(b)
}

type ioWrite struct{}

func (d *ioWrite) Write(b []byte) (int, error) {
	fmt.Println("out> ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader ioRead
		writer ioWrite
	)

	input := make([]byte, 4096)
	r, err := reader.Read(input)

	if err != nil {
		log.Fatalln("Unnable to read data")
	}

	fmt.Printf("Read %d bytes \n", r)

	w, err := writer.Write(input)

	if err != nil {
		log.Fatalln("Unnable to write data")
	}

	fmt.Printf("Wrote %d bytes", w)
}
