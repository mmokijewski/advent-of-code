package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	timeStart := time.Now()
	inputFile, _ := os.Open("input")

	towels := make(map[string]bool)
	result := 0

	scanner := bufio.NewScanner(inputFile)
	towelsLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if towelsLine {
			for _, towel := range strings.Split(line, ", ") {
				towels[towel] = true
			}
			towelsLine = false
		} else if line != "" {
			matches := make([]bool, len(line)+1)
			matches[0] = true

			for i := 1; i <= len(line); i++ {
				for j := 0; j < i; j++ {
					if matches[j] && towels[line[j:i]] {
						matches[i] = true
						break
					}
				}
			}
			if matches[len(line)] {
				result++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
