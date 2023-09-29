package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Writing data to the file\n"))
	// size, err := f.WriteString("Hello, world!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", size)

	f.Close()

	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(file))

	// Data streaming
	archive, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(archive)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}

