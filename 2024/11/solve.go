package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func intLen(i int) int {
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func splitIntToTwo(input int) (int, int) {
	splitValue := int(math.Pow10(intLen(input) / 2))
	right := input % splitValue
	left := (input - right) / splitValue

	return left, right
}

func blink(stones map[int]int, times int) map[int]int {
	newStones := make(map[int]int)

	if times == 0 {
		return stones
	}

	for stone, count := range stones {
		if stone == 0 {
			newStones[1] += count
		} else if intLen(stone)%2 == 0 {
			left, right := splitIntToTwo(stone)
			newStones[left] += count
			newStones[right] += count
		} else {
			newStones[stone*2024] += count
		}
	}
	return blink(newStones, times-1)
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	stones := make(map[int]int)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		for _, stone := range strings.Split(line, " ") {
			stoneNum, _ := strconv.Atoi(stone)
			stones[stoneNum]++
		}
	}

	sum := 0
	stones = blink(stones, 75)

	for _, count := range stones {
		sum += count
	}

	fmt.Printf("Part 2: %d\n", sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
