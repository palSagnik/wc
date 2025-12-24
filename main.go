package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Size bool `short:"c" long:"size" description:"print the size of file in bytes"`
	Line bool `short:"l" long:"lines" description:"print the number of lines"`
}

func main() {
	args, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	if len(args) == 0 {
		fmt.Printf("Err: Filename is required")
	}
	filepath := args[0]

	// SIZE LOGIC
	// bytes
	if opts.Size {
		info, err := os.Stat(filepath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d ", info.Size())
	}

	// lines
	if opts.Line {
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var lines int64
		for scanner.Scan() {
			_ = scanner.Text()
			lines++
		}

		fmt.Printf("%d ", lines)
	}

	fmt.Printf("%s\n", filepath)
}