package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	timeStart := time.Now()
	inputFile, _ := os.Open("input")

	var groups [][]string
	var allComputers []string
	compConnections := make(map[string][]string)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		players := strings.Split(line, "-")
		player1 := players[0]
		player2 := players[1]
		compConnections[player1] = append(compConnections[player1], player2)
		compConnections[player2] = append(compConnections[player2], player1)
		if !slices.Contains(allComputers, player1) {
			allComputers = append(allComputers, player1)
		}
		if !slices.Contains(allComputers, player2) {
			allComputers = append(allComputers, player2)
		}
	}

	for k, v := range compConnections {
		for _, compToCheck := range v {
			for _, secondCompToCheck := range compConnections[compToCheck] {
				if slices.Contains(compConnections[secondCompToCheck], k) {
					newGroup := []string{k, compToCheck, secondCompToCheck}
					isNewGroup := true
					for _, group := range groups {
						if slices.Contains(group, k) && slices.Contains(group, compToCheck) && slices.Contains(group, secondCompToCheck) {
							isNewGroup = false
						}
					}
					if isNewGroup {
						groups = append(groups, newGroup)
					}
				}
			}
		}
	}

	part1Sum := 0
	for _, group := range groups {
		for _, comp := range group {
			if string(comp[0]) == "t" {
				part1Sum++
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
