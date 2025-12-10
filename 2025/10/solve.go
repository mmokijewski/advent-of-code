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

func replaceAtIndex(in string, i int) string {
	out := []rune(in)
	if out[i] == '#' {
		out[i] = '.'
	} else {
		out[i] = '#'
	}
	return string(out)
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	var part1 int

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		machine := scanner.Text()
		var buttons [][]int
		split := strings.Fields(machine)
		desiredState := split[0]
		startState := strings.Replace(desiredState, "#", ".", -1)

		for i := 1; i < len(split)-1; i++ {
			buttonString := split[i][1 : len(split[i])-1]
			buttonIndexes := strings.Split(buttonString, ",")
			var buttonTmp []int
			for _, b := range buttonIndexes {
				num, _ := strconv.Atoi(b)
				buttonTmp = append(buttonTmp, num)
			}
			buttons = append(buttons, buttonTmp)
		}

		stateMap := make(map[string]int)
		stateMap[startState] = 0

		clicks := 0
		found := false
		for !found {
			for state, seqClicks := range stateMap {
				if seqClicks == clicks {
					for _, b := range buttons {
						newState := state
						for _, index := range b {
							newState = replaceAtIndex(newState, index+1)
						}
						_, exists := stateMap[newState]
						if exists {
							continue
						} else {
							stateMap[newState] = clicks + 1
						}
						if newState == desiredState {
							found = true
							part1 += clicks + 1
							break
						}
					}
				}
				if found {
					break
				}
			}
			clicks++
		}
	}

	fmt.Printf("Part1 : %d\n", part1)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
