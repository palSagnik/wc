package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Size string `short:"c" long:"size" description:"print the size of file in bytes"`
	Line string `short:"l" long:"lines" description:"print the number of lines"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	filepath := opts.Size

	// SIZE LOGIC
	// bytes
	info, err := os.Stat(filepath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	byteSize := info.Size()

	// lines
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	scanner := bufio.NewScanner(file)
	var lines int64
	for scanner.Scan() {
		_ = scanner.Text()
		lines++
	}

	//fmt.Printf("%d %s\n", byteSize, filepath)
}