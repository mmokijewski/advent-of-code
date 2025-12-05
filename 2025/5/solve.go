package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type idRange struct {
	start, end int
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	part1sum := 0
	var freshIdRanges []idRange

	scanner := bufio.NewScanner(inputFile)
	ingredientsStarted := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			ingredientsStarted = true
			continue
		}
		if !ingredientsStarted {
			split := strings.Split(line, "-")
			startId, _ := strconv.Atoi(split[0])
			endId, _ := strconv.Atoi(split[1])
			freshIdRanges = append(freshIdRanges, idRange{startId, endId})
		} else {
			ingredientId, _ := strconv.Atoi(line)
			for _, singleIdRange := range freshIdRanges {
				if ingredientId >= singleIdRange.start && ingredientId <= singleIdRange.end {
					part1sum++
					break
				}
			}
		}
	}

	fmt.Printf("Part1 : %d\n", part1sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
