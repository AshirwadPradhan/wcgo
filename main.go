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
	flag.Parse()
	fmt.Println(wc(os.Stdin, *lines))
}

func wc(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}
