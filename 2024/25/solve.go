package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func countHashes(array []rune) int {
	result := 0
	for _, sign := range array {
		if sign == '#' {
			result++
		}
	}
	return result
}

func main() {
	timeStart := time.Now()
	inputFile, _ := os.Open("input")

	var keys [][]int
	var locks [][]int

	scanner := bufio.NewScanner(inputFile)
	i := 0
	temp := make([][]rune, 7)
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			temp = make([][]rune, 5)
			for j := range temp {
				temp[j] = make([]rune, 7)
			}
		}
		if line == "" {
			i = 0
			continue
		}
		for j, sign := range line {
			temp[j][i] = sign
		}

		if i == 6 {
			if temp[0][0] == '#' && temp[1][0] == '#' && temp[2][0] == '#' && temp[3][0] == '#' && temp[4][0] == '#' {
				locks = append(locks, []int{countHashes(temp[0]), countHashes(temp[1]), countHashes(temp[2]), countHashes(temp[3]), countHashes(temp[4])})
			} else {
				keys = append(keys, []int{countHashes(temp[0]), countHashes(temp[1]), countHashes(temp[2]), countHashes(temp[3]), countHashes(temp[4])})
			}
		}
		i++
	}

	part1Result := 0
	for _, key := range keys {
		for _, lock := range locks {
			if key[0]+lock[0] < 8 && key[1]+lock[1] < 8 && key[2]+lock[2] < 8 && key[3]+lock[3] < 8 && key[4]+lock[4] < 8 {
				part1Result++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1Result)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
