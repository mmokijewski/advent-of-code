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

func checkSingeRule(update string, rule string) bool {
	leftNum := strings.Split(rule, "|")[0]
	rightNum := strings.Split(rule, "|")[1]
	reg := regexp.MustCompile(rightNum + ",.*" + leftNum)
	result := reg.FindString(update)
	if len(result) != 0 {
		return false
	} else {
		return true
	}
}

func checkIfCorrect(update string, rules []string) bool {
	for _, rule := range rules {
		if !checkSingeRule(update, rule) {
			return false
		}
	}
	return true
}

func replaceNumbers(update string, rule string) string {
	leftNum := strings.Split(rule, "|")[0]
	rightNum := strings.Split(rule, "|")[1]
	update = strings.Replace(update, leftNum, rightNum, 1)
	update = strings.Replace(update, rightNum, leftNum, 1)
	return update
}

func getMiddleNum(update string) int {
	numbers := strings.Split(update, ",")
	middleNum, _ := strconv.Atoi(numbers[len(numbers)/2])
	return middleNum
}

func main() {
	start := time.Now()
	inputFile, _ := os.Open("input")

	part1Sum := 0
	part2Sum := 0
	var rules []string

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rules = append(rules, line)
		} else if line == "" {
			continue
		} else {
			isCorrect := checkIfCorrect(line, rules)
			if isCorrect {
				part1Sum += getMiddleNum(line)
			} else {
				for !isCorrect {
					for _, rule := range rules {
						if !checkSingeRule(line, rule) {
							line = replaceNumbers(line, rule)
						}
					}
					isCorrect = checkIfCorrect(line, rules)
				}
				part2Sum += getMiddleNum(line)
			}
		}
	}

	fmt.Printf("Part 1 sum: %d\n", part1Sum)
	fmt.Printf("Part 2 sum: %d\n", part2Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
