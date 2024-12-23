package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

type ChangeSequence struct {
	current    int
	last       int
	secondLast int
	thirdLast  int
}

type SequenceValue struct {
	priceIndexes []int
	bananas      int
}

type Price struct {
	current        int
	changeSequence ChangeSequence
}

func main() {
	timeStart := time.Now()
	inputFile, _ := os.Open("input")

	var prices []Price
	scanner := bufio.NewScanner(inputFile)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		price, _ := strconv.Atoi(line)
		initialSequence := ChangeSequence{10, 10, 10, 10}
		prices = append(prices, Price{price, initialSequence})
		i++
	}

	sequences := make(map[ChangeSequence]SequenceValue)

	for range 2000 {
		for i, price := range prices {
			newPrice := price.current
			newPrice = ((newPrice * 64) ^ newPrice) % 16777216
			newPrice = ((newPrice / 32) ^ newPrice) % 16777216
			newPrice = ((newPrice * 2048) ^ newPrice) % 16777216
			modifiedPrice := newPrice % 10
			change := modifiedPrice - (price.current % 10)

			prices[i].current = newPrice

			prices[i].changeSequence.thirdLast = prices[i].changeSequence.secondLast
			prices[i].changeSequence.secondLast = prices[i].changeSequence.last
			prices[i].changeSequence.last = prices[i].changeSequence.current
			prices[i].changeSequence.current = change

			if !slices.Contains(sequences[prices[i].changeSequence].priceIndexes, i) && prices[i].changeSequence.thirdLast != 10 {
				var newIndexes []int
				newIndexes = append(newIndexes, sequences[prices[i].changeSequence].priceIndexes...)
				newIndexes = append(newIndexes, i)
				newBananas := sequences[prices[i].changeSequence].bananas + modifiedPrice
				sequences[prices[i].changeSequence] = SequenceValue{newIndexes, newBananas}
			}
		}
	}

	part1Sum := 0
	for _, price := range prices {
		part1Sum += price.current
	}

	part2Sum := 0
	for _, sequence := range sequences {
		if sequence.bananas > part2Sum {
			part2Sum = sequence.bananas
		}
	}

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Part 2: %d\n", part2Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
