package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main() {
	fileNames := os.Args[1:]
	counts := make(map[string]int)
	for _, fileName := range fileNames {
		textBuf, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exec1.4: %v \n", err)
			continue
		}
		lines := strings.Split(string(textBuf), "\n")
		for _, line := range lines {
			counts[line]++
		}
	}
	for k, v := range counts {
		if (v > 1) {
			fmt.Printf("line: %s, \t numbers: %d\n", k, v)
		}
	}
}
