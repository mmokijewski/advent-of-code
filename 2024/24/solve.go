package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type WireGate struct {
	firstCable, operation, secondCable, resultCable string
}

func main() {
	timeStart := time.Now()
	inputFile, _ := os.Open("input")

	cables := make(map[string]int)
	var wireGates []WireGate

	readCables := true
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readCables = false
			continue
		}
		if readCables {
			signs := strings.Split(line, ": ")
			cableName := signs[0]
			value, _ := strconv.Atoi(signs[1])
			cables[cableName] = value
		} else {
			var firstCable, operation, secondCable, resultCable string
			fmt.Sscanf(line, "%s %s %s -> %s", &firstCable, &operation, &secondCable, &resultCable)
			wireGates = append(wireGates, WireGate{firstCable, operation, secondCable, resultCable})
		}
	}

	somethingChanged := true
	for somethingChanged {
		somethingChanged = false
		for _, wireGate := range wireGates {
			_, resultCableExist := cables[wireGate.resultCable]
			_, firstCableReady := cables[wireGate.firstCable]
			_, secondCableReady := cables[wireGate.secondCable]
			if !resultCableExist && firstCableReady && secondCableReady {
				somethingChanged = true
				if wireGate.operation == "AND" {
					cables[wireGate.resultCable] = cables[wireGate.firstCable] & cables[wireGate.secondCable]
				} else if wireGate.operation == "OR" {
					cables[wireGate.resultCable] = cables[wireGate.firstCable] | cables[wireGate.secondCable]
				} else if wireGate.operation == "XOR" {
					cables[wireGate.resultCable] = cables[wireGate.firstCable] ^ cables[wireGate.secondCable]
				}
			}
		}
	}

	var zCables []string
	var binaryResult string

	for cable := range cables {
		if string(cable[0]) == "z" {
			zCables = append(zCables, cable)
		}
	}
	slices.Sort(zCables)

	for _, cable := range zCables {
		binaryResult = strconv.Itoa(cables[cable]) + binaryResult
	}
	decimalResult, _ := strconv.ParseInt(binaryResult, 2, 64)

	fmt.Printf("Part 1: %d\n", decimalResult)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(timeStart).Milliseconds())
}
