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

func findNextBiggest(bank string, index int, maxIndex int) (string, int) {
	maxNum := 0
	maxNumIndex := 0
	for i := index; i <= maxIndex; i++ {
		num := int(bank[i] - '0')
		if num > maxNum {
			maxNum = num
			maxNumIndex = i
		}
	}
	maxNumString := strconv.Itoa(maxNum)
	return maxNumString, maxNumIndex
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	maxLength := 12

	totalSum := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		bank := scanner.Text()
		maxNumString := ""
		previousIndex := 0

		firstMaxNum, previousIndex := findNextBiggest(bank, 0, len(bank)-maxLength)
		maxNumString = maxNumString + firstMaxNum

		for i := maxLength - 1; i > 0; i-- {
			maxNum, index := findNextBiggest(bank, previousIndex+1, len(bank)-i)
			previousIndex = index
			maxNumString = maxNumString + maxNum
		}
		maxNum, _ := strconv.Atoi(maxNumString)

		totalSum += maxNum
	}

	fmt.Printf("Total sum for max length '%d': %d\n", maxLength, totalSum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
