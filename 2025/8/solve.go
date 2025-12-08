package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type junctionBox struct {
	x, y, z int
}

type connection struct {
	boxIndex1, boxIndex2 int
	distance             float64
}

func main() {
	start := time.Now()
	inputFile, err := os.Open("input")
	checkError(err)

	pairsToConnect := 1000

	var part1 int
	var part2 int
	var junctionBoxes []junctionBox
	var allConnections []connection
	var groups [][]int

	// Read junction boxes
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		box := scanner.Text()
		split := strings.Split(box, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		junctionBoxes = append(junctionBoxes, junctionBox{x, y, z})
	}

	// Go through the boxes and count distances
	for i, box := range junctionBoxes {
		for j := i - 1; j >= 0; j-- {
			boxToMeasure := junctionBoxes[j]
			distance := math.Sqrt(math.Pow(float64(box.x-boxToMeasure.x), 2) + math.Pow(float64(box.y-boxToMeasure.y), 2) + math.Pow(float64(box.z-boxToMeasure.z), 2))
			allConnections = append(allConnections, connection{j, i, distance})
		}
	}

	// Sort possible connections by distance
	sort.Slice(allConnections, func(i int, j int) bool {
		return allConnections[i].distance < allConnections[j].distance
	})

	for _, pair := range allConnections {
		// Check if some of the box already exists in some group
		boxExistsInGroup := false
		for groupIndex, group := range groups {
			if slices.Contains(group, pair.boxIndex1) {
				boxExistsInGroup = true
				if !slices.Contains(group, pair.boxIndex2) {
					groupsConnected := false
					for i := range groups {
						if i == groupIndex {
							continue
						}
						if slices.Contains(groups[i], pair.boxIndex2) {
							groups[groupIndex] = append(groups[groupIndex], groups[i]...)
							groups[i] = []int{}
							groupsConnected = true
							break
						}
					}
					if !groupsConnected {
						groups[groupIndex] = append(groups[groupIndex], pair.boxIndex2)
					}
					break
				}
			} else if slices.Contains(group, pair.boxIndex2) {
				boxExistsInGroup = true
				groupsConnected := false
				for i := range groups {
					if i == groupIndex {
						continue
					}
					if slices.Contains(groups[i], pair.boxIndex1) {
						groups[groupIndex] = append(groups[groupIndex], groups[i]...)
						groups[i] = []int{}
						groupsConnected = true
						break
					}
				}
				if !groupsConnected {
					groups[groupIndex] = append(groups[groupIndex], pair.boxIndex1)
				}
				break
			}
		}
		if !boxExistsInGroup {
			groups = append(groups, []int{pair.boxIndex1, pair.boxIndex2})
		}
		pairsToConnect--

		// Sort groups by length
		sort.Slice(groups, func(i int, j int) bool {
			return len(groups[i]) > len(groups[j])
		})

		// Catch condition for part 1
		if pairsToConnect == 0 {
			part1 = len(groups[0]) * len(groups[1]) * len(groups[2])
		}

		// Connect until there is one group containing all boxes
		if len(groups[0]) == len(junctionBoxes) {
			part2 = junctionBoxes[pair.boxIndex1].x * junctionBoxes[pair.boxIndex2].x
			break
		}
	}

	fmt.Printf("Part1 : %d\n", part1)
	fmt.Printf("Part2 : %d\n", part2)
	fmt.Printf("Total time elapsed: %dms\n", time.Since(start).Milliseconds())
}
