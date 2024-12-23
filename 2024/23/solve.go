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
	var lanParties [][]string
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

	for k, v := range compConnections {
		newParty := true
		for i, party := range lanParties {
			appendToParty := true
			for _, partyComp := range party {
				if partyComp == k {
					continue
				}
				if !slices.Contains(v, partyComp) {
					appendToParty = false
					break
				}
			}
			if appendToParty {
				lanParties[i] = append(lanParties[i], k)
				slices.Sort(lanParties[i])
				newParty = false
			}
		}
		if newParty {
			lanParties = append(lanParties, []string{k})
		}
	}

	var biggestParty []string
	for _, party := range lanParties {
		if len(party) > len(biggestParty) {
			biggestParty = party
		}
	}

	var part2Result string
	for i, comp := range biggestParty {
		part2Result += comp
		if i != len(biggestParty)-1 {
			part2Result += ","
		}
	}

	fmt.Printf("Part 1: %d\n", part1Sum)
	fmt.Printf("Part 2: %s\n", part2Result)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
