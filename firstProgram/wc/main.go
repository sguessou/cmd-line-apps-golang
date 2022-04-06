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
	countBytes := flag.Bool("b", false, "Count number of bytes")

	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *countBytes))
}

func count(r io.Reader, countWords bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	switch {
	case countWords:
		scanner.Split(bufio.ScanWords)
	case countBytes:
		scanner.Split(bufio.ScanRunes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}
