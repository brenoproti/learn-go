package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	const fileName string = "file-manipulation/test.txt"

	// Create file
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	// size, err := f.WriteString("Hello, World!\n")
	size, err := f.Write([]byte("Hello, World!\n"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("The size of the file is: %d bytes \n", size)
	f.Close()

	// Read file
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// Read file with buffer
	file2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 5)

	for {
		n, _ := reader.Read(buffer)

		if n == 0 {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	// Delete file
	err = os.Remove(fileName)
	if err != nil {
		panic(err)
	}

}
