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

	part1sum := 0
	part2sum := 0

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		bank := scanner.Text()
		firstMaxNum := 0
		secondMaxNum := 0
		firstMaxNumIndex := 0
		for index, numString := range bank {
			num := int(numString - '0')
			if num > firstMaxNum && index < len(bank)-1 {
				firstMaxNum = num
				firstMaxNumIndex = index
				secondMaxNum = 0
			}
			if num > secondMaxNum && index > firstMaxNumIndex {
				secondMaxNum = num
			}
		}
		maxNum, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstMaxNum, secondMaxNum))
		part1sum += maxNum
	}

	fmt.Printf("Part 1 sum: %d\n", part1sum)
	fmt.Printf("Part 2 sum: %d\n", part2sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
