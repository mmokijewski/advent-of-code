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

func part1(numLine []int) int {
	var filteredLine []int
	sum := 0
	for i, num := range numLine {
		space := i%2 == 1
		for j := 0; j < num; j++ {
			if space {
				filteredLine = append(filteredLine, -1)
			} else {
				filteredLine = append(filteredLine, i/2)
			}
		}
	}

	arrayLen := len(filteredLine)
	numLineReverted := make([]int, arrayLen)
	copy(numLineReverted, filteredLine)
	slices.Reverse(numLineReverted)
	for i, num := range numLineReverted {
		firstSpace := slices.Index(filteredLine, -1)
		if firstSpace == -1 || firstSpace > arrayLen-i {
			break
		}
		if num == -1 {
			continue
		} else {
			filteredLine[firstSpace] = num
			filteredLine = slices.Delete(filteredLine, arrayLen-1-i, arrayLen-i)
		}
	}

	for i, num := range filteredLine {
		if num != -1 {
			sum += i * num
		}
	}

	return sum
}

func part2(numLine []int) int {

	sum := 0
	var lineInBlocks [][]int
	var finalArray []int

	for i, num := range numLine {
		space := i%2 == 1
		var tempBlock []int
		for j := 0; j < num; j++ {
			if space {
				tempBlock = append(tempBlock, -1)
			} else {
				tempBlock = append(tempBlock, i/2)
			}
		}
		lineInBlocks = append(lineInBlocks, tempBlock)
	}

	blockArrayLen := len(lineInBlocks)
	for i := 0; i < blockArrayLen; i++ {
		block := lineInBlocks[i]
		if slices.Contains(block, -1) {
			currentBlock := block
			for {
				blockLength := len(currentBlock)
				if blockLength == 0 {
					break
				}
				findBlockToFillWith := func(E []int) bool {
					if len(E) > 0 && len(E) <= blockLength && !slices.Contains(E, -1) {
						return true
					} else {
						return false
					}
				}
				slices.Reverse(lineInBlocks)
				blockToFillWithIndex := slices.IndexFunc(lineInBlocks, findBlockToFillWith)
				slices.Reverse(lineInBlocks)
				normalIndex := blockArrayLen - 1 - blockToFillWithIndex
				if blockToFillWithIndex == -1 || normalIndex < i {
					finalArray = append(finalArray, currentBlock...)
					break
				} else {
					desiredBlock := lineInBlocks[normalIndex]
					finalArray = append(finalArray, desiredBlock...)
					var newEmptyBlock []int
					for j := 0; j < len(desiredBlock); j++ {
						newEmptyBlock = append(newEmptyBlock, -1)
					}
					lineInBlocks[normalIndex] = newEmptyBlock
					lenDiff := blockLength - len(desiredBlock)
					if lenDiff > 0 {
						var dummyBlock []int
						for j := 0; j < lenDiff; j++ {
							dummyBlock = append(dummyBlock, -1)
						}
						currentBlock = dummyBlock
					} else {
						currentBlock = []int{}
					}
				}
			}
		} else {
			finalArray = append(finalArray, block...)
		}
	}

	for i, num := range finalArray {
		if num != -1 {
			sum += i * num
		}
	}

	return sum
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	var numLine []int

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		for _, sign := range strings.Split(line, "") {
			currentNum, _ := strconv.Atoi(sign)
			numLine = append(numLine, currentNum)
		}
	}

	part1Sum := part1(numLine)
	part2Sum := part2(numLine)

	fmt.Printf("Part 1 sum: %d\n", part1Sum)
	fmt.Printf("Part 2 sum: %d\n", part2Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
