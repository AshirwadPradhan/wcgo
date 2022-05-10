package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()
	fmt.Println(wc(os.Stdin, *lines, *bytes))
}

func wc(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else if countBytes && countLines {
		fmt.Printf("Cannot specify both flags at once!!!")
		os.Exit(1)
	}

	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}
