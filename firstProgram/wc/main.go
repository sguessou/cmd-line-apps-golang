package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	words := flag.Bool("w", false, "Count words")
	countBytes := flag.Bool("b", false, "Count number of bytes")

	flag.Parse()

	fmt.Println(count(os.Stdin, *words, *countBytes))
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
