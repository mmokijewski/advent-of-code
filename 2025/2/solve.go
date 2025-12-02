package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		for _, idRange := range strings.Split(line, ",") {
			startId, _ := strconv.Atoi(strings.Split(idRange, "-")[0])
			endId, _ := strconv.Atoi(strings.Split(idRange, "-")[1])
			fmt.Println(startId, endId)
			for i := startId; i <= endId; i++ {
				idString := strconv.Itoa(i)
				idLength := len(idString)
				if idLength%2 != 0 {
					continue
				}
				for j := 0; j < idLength/2; j++ {
					digits := strings.Split(idString, "")
					if digits[j] != digits[j+(idLength/2)] {
						break
					}
					if j == idLength/2-1 {
						part1sum += i
					}
				}
			}
		}
	}

	fmt.Printf("Part 1 sum: %d\n", part1sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
