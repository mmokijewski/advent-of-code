package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func blink(stones []string) []string {

	var newStones []string

	for _, stone := range stones {
		stoneLength := len(stone)
		if stone == "0" {
			newStones = append(newStones, "1")
		} else if stoneLength%2 == 0 {
			firstHalf := ""
			secondHalf := ""
			for i, char := range stone {
				if i < stoneLength/2 {
					firstHalf = fmt.Sprintf("%s%s", firstHalf, string(char))
				} else {
					secondHalf = fmt.Sprintf("%s%s", secondHalf, string(char))
				}
			}
			left, _ := strconv.Atoi(firstHalf)
			right, _ := strconv.Atoi(secondHalf)
			newStones = append(newStones, strconv.Itoa(left))
			newStones = append(newStones, strconv.Itoa(right))
		} else {
			stoneNum, _ := strconv.Atoi(stone)
			newStones = append(newStones, strconv.Itoa(stoneNum*2024))
		}
	}
	return newStones
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	var stones []string

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		for _, stone := range strings.Split(line, " ") {
			stones = append(stones, stone)
		}
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Printf("Part 1: %d\n", len(stones))
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
