package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(wc(os.Stdin))
}

func wc(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	
	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}
