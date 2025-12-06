package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var operations []string
	var results []int

	scanner := bufio.NewScanner(inputFile)
	reg := regexp.MustCompile("[+*]")
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSuffix(line, "  ")
		if strings.HasPrefix(line, "*") || strings.HasPrefix(line, "+") {
			split := reg.FindAllString(line, -1)
			operations = append(operations, split...)
		}
	}

	inputFile, _ = os.Open("input")
	scanner = bufio.NewScanner(inputFile)
	reg = regexp.MustCompile("\\d+")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "*") || strings.HasPrefix(line, "+") {
			continue
		}

		split := reg.FindAllString(line, -1)
		for i, numString := range split {
			num, _ := strconv.Atoi(numString)
			if len(results) < len(operations) {
				results = append(results, num)
			} else {
				if operations[i] == "+" {
					results[i] = results[i] + num
				} else if operations[i] == "*" {
					results[i] = results[i] * num
				}
			}
		}
	}

	part1sum := 0
	for _, result := range results {
		part1sum += result
	}

	fmt.Printf("Part1 : %d\n", part1sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
