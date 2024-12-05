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

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	part1Sum := 0
	var rules []string
	var updates []string

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rules = append(rules, line)
		} else if line == "" {
			continue
		} else {
			updates = append(updates, line)
		}
	}

	rulesCount := len(rules)

	for i := range updates {

		isCorrect := true
		j := 0

		for isCorrect && j < rulesCount {
			leftNum := strings.Split(rules[j], "|")[0]
			rightNum := strings.Split(rules[j], "|")[1]
			reg := regexp.MustCompile(rightNum + ",.*" + leftNum)
			result := reg.FindString(updates[i])
			if len(result) != 0 {
				isCorrect = false
			}
			j++
		}

		if isCorrect {
			numbers := strings.Split(updates[i], ",")
			middleNum, _ := strconv.Atoi(numbers[len(numbers)/2])
			part1Sum += middleNum
		}
	}

	fmt.Printf("Part 1 sum: %d\n", part1Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
