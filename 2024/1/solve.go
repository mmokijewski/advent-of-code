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

func countOccurrences(slice []int, numToCount int) int {
	count := 0
	for i := range slice {
		if slice[i] == numToCount {
			count++
		}
	}
	return count
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	totalDistance := 0
	part2sum := 0

	var leftColumn = make([]int, 5)
	var rightColumn = make([]int, 5)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "   ")
		leftNum, err := strconv.Atoi(items[0])
		checkError(err)
		rightNum, err := strconv.Atoi(items[1])
		checkError(err)
		leftColumn = append(leftColumn, leftNum)
		rightColumn = append(rightColumn, rightNum)
	}

	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	for i := range leftColumn {
		leftNum := leftColumn[i]
		rightNum := rightColumn[i]

		if leftNum > rightNum {
			totalDistance += leftNum - rightNum
		} else {
			totalDistance += rightNum - leftNum
		}

		part2sum += leftNum * countOccurrences(rightColumn, leftNum)
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
	fmt.Printf("Part 2 sum: %d\n", part2sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
