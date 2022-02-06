package main

import (
	"fmt"
	"io"
	"os"
)

/**
 * Command to execute: `go run main.go myfile.txt`
 */
func main() {
	// Use os.Args to read arguments, which is a slice of type string.
	// Open file: check out docs on the `Open` function in the `os` package.
	fileName := os.Args[1]
	fmt.Println("File name to read:", fileName)

	// Read file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
