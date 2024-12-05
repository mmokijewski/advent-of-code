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
	i := 0
	for i < len(rules) {
		if !checkSingeRule(update, rules[i]) {
			return false
		}
		i++
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

	for i := range updates {
		isCorrect := checkIfCorrect(updates[i], rules)
		if isCorrect {
			part1Sum += getMiddleNum(updates[i])
		} else {
			for !isCorrect {
				for j := range rules {
					if !checkSingeRule(updates[i], rules[j]) {
						updates[i] = replaceNumbers(updates[i], rules[j])
					}
				}
				isCorrect = checkIfCorrect(updates[i], rules)
			}
			part2Sum += getMiddleNum(updates[i])
		}
	}

	fmt.Printf("Part 1 sum: %d\n", part1Sum)
	fmt.Printf("Part 2 sum: %d\n", part2Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
