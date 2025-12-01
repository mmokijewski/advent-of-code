package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	pos := 50
	count := 0

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0:1]
		dist, _ := strconv.Atoi(line[1:])
		if dir == "R" {
			pos = (pos + dist) % 100
		} else {
			pos = (pos - dist) % 100
			if pos < 0 {
				pos = 100 + pos
			}
		}

		if pos == 0 {
			count++
		}
	}
	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
