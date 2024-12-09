package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func part1(line string) int {
	var numLine []int
	sum := 0
	for i, sign := range strings.Split(line, "") {
		currentNum, _ := strconv.Atoi(sign)
		space := i%2 == 1
		for j := 0; j < currentNum; j++ {
			if space {
				numLine = append(numLine, -1)
			} else {
				numLine = append(numLine, i/2)
			}
		}
	}

	arrayLen := len(numLine)
	numLineReverted := make([]int, arrayLen)
	copy(numLineReverted, numLine)
	slices.Reverse(numLineReverted)
	for i, num := range numLineReverted {
		firstSpace := slices.Index(numLine, -1)
		if firstSpace == -1 || firstSpace > arrayLen-i {
			break
		}
		if num == -1 {
			continue
		} else {
			numLine[firstSpace] = num
			numLine = slices.Delete(numLine, arrayLen-1-i, arrayLen-i)
		}
	}

	for i, num := range numLine {
		if num != -1 {
			sum += i * num
		}
	}

	return sum
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	var line string

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line = scanner.Text()
	}

	part1Sum := part1(line)

	fmt.Printf("Part 1 sum: %d\n", part1Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
