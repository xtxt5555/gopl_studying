package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fileNames := os.Args[1:]
	counts := make(map[string]int)
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exec1.4: %v \n", err)
			continue
		}
		countLines(file, counts)
		file.Close()
	}
	for k, v := range counts {
		if (v > 1) {
			fmt.Printf("line: %s, \t numbers: %d\n", k, v)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}