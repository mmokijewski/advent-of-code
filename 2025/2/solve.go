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

func checkId(idString string, repeats int) bool {
	idLength := len(idString)
	for i := 0; i < idLength/repeats; i++ {
		digits := strings.Split(idString, "")

		for j := range repeats {
			if digits[i] != digits[i+(j*idLength/repeats)] {
				return false
			}
		}
		if i == (idLength/repeats)-1 {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	part1sum := 0
	part2sum := 0

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		for _, idRange := range strings.Split(line, ",") {
			startId, _ := strconv.Atoi(strings.Split(idRange, "-")[0])
			endId, _ := strconv.Atoi(strings.Split(idRange, "-")[1])
			for i := startId; i <= endId; i++ {
				idString := strconv.Itoa(i)
				idLength := len(idString)
				for repeats := 2; repeats <= idLength; repeats++ {
					if idLength%repeats == 0 {
						if checkId(idString, repeats) {
							if repeats == 2 {
								part1sum += i
								part2sum += i
							} else {
								part2sum += i
							}
							break
						}
					}
				}
			}
		}
	}

	fmt.Printf("Part 1 sum: %d\n", part1sum)
	fmt.Printf("Part 2 sum: %d\n", part2sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
