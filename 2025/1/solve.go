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

	pos := 50
	countPart1 := 0
	countPart2 := 0
	var prevPos int

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0:1]
		dist, _ := strconv.Atoi(line[1:])
		prevPos = pos
		if dir == "R" {
			divideResult := (pos + dist) / 100
			divideRest := (pos + dist) % 100
			if divideResult > 0 {
				// It's how many times it went through 0 or ended on 0 when turning to right
				countPart2 += divideResult
			}
			pos = divideRest
		} else {
			divideResult := (pos - dist) / -100
			divideRest := (pos - dist) % 100
			if divideResult > 0 {
				// It how many times it went through 0 when turning to left
				countPart2 += divideResult
			}
			// Additional case when it did not go through 0 but just ended on 0 when turning to left
			if divideRest == 0 {
				countPart2++
			}

			pos = divideRest
			if pos < 0 {
				// Case when it went through the 0 just once when turning to left
				if prevPos != 0 {
					countPart2++
				}
				pos = 100 + pos
			}
		}

		if pos == 0 {
			countPart1++
		}
	}

	fmt.Printf("Count: %d\n", countPart1)
	fmt.Printf("Count part2: %d\n", countPart2)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
