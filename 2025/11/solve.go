package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	part1 := 0

	devices := make(map[string][]string)
	var sequences [][]string
	sequences = append(sequences, []string{"you"})

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		name := line[0:3]
		nextDeviceNames := strings.Split(line[5:], " ")
		devices[name] = nextDeviceNames
	}

	for len(sequences) > 0 {
		var tmpSequences [][]string
		for _, seq := range sequences {
			seqLength := len(seq)
			last := seq[seqLength-1]
			lastNextDevices := devices[last]
			for _, devName := range lastNextDevices {
				if devName == "out" {
					part1++
					continue
				}
				if slices.Contains(seq, devName) {
					continue
				}
				var newSeq []string
				newSeq = append(newSeq, seq...)
				newSeq = append(newSeq, devName)
				tmpSequences = append(tmpSequences, newSeq)
			}
		}
		sequences = tmpSequences
	}

	fmt.Printf("Part1 : %d\n", part1)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
