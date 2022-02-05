package main

import (
	"fmt"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

/**
 * Command to execute: `go run main.go myfile.txt`
 */
func main() {
	// Use os.Args to read arguments, which is a slice of type string.
	// Open file: check out docs on the `Open` function in the `os` package.
	fileName := os.Args[1]
	fmt.Println("File name to read:", fileName)

	// Read file
	contents, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Output contents
	lw := logWriter{}
	lw.Write(contents)
}
