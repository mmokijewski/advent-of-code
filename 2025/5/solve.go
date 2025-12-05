package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	part2sum := 0
	sort.Slice(freshIdRanges, func(i, j int) bool {
		return freshIdRanges[i].start < freshIdRanges[j].start
	})
	var currentEnd int
	for i, singleIdRange := range freshIdRanges {
		if i == 0 {
			part2sum += singleIdRange.end - singleIdRange.start + 1
			currentEnd = singleIdRange.end
		} else {
			if singleIdRange.start <= currentEnd {
				if singleIdRange.end <= currentEnd {
					continue
				} else {
					part2sum += singleIdRange.end - currentEnd
					currentEnd = singleIdRange.end
				}
			} else {
				part2sum += singleIdRange.end - singleIdRange.start + 1
				currentEnd = singleIdRange.end
			}
		}
	}

	fmt.Printf("Part1 : %d\n", part1sum)
	fmt.Printf("Part2 : %d\n", part2sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
