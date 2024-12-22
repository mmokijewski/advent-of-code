package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	timeStart := time.Now()
	inputFile, _ := os.Open("input")

	var prices []int
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		price, _ := strconv.Atoi(line)
		prices = append(prices, price)
	}

	for range 2000 {
		for i, price := range prices {
			price = ((price * 64) ^ price) % 16777216
			price = ((price / 32) ^ price) % 16777216
			price = ((price * 2048) ^ price) % 16777216
			prices[i] = price
		}
	}

	part1Sum := 0
	for _, price := range prices {
		part1Sum += price
	}

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
