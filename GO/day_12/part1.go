package day12

import (
	"fmt"
	"strings"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadString(path[f])
	splitData := strings.Split(data, "\n\n")

	shapeMap := parseShapes(splitData[:len(splitData)-1])
	count := 0

	for _, line := range strings.Split(strings.TrimSpace(splitData[len(splitData)-1]), "\n") {
		if isValid(line, shapeMap) {
			count++
		}
	}

	fmt.Printf("Day 12 | Part 1-Final: %d\n", count)
}

func parseShapes(s []string) map[int]int {
	idxMap := make(map[int]int)

	for _, shape := range s {
		idx, rest, _ := strings.Cut(shape, ":")
		size := 0
		for _, line := range strings.Fields(rest) {
			for _, ch := range line {
				if ch == '#' {
					size++
				}
			}
		}

		idxMap[utils.FastAtoi(idx)] = size
	}

	return idxMap
}

func isValid(s string, idxMap map[int]int) bool {
	size, rest, _ := strings.Cut(s, ":")
	parts := strings.Split(size, "x")
	lim := utils.FastAtoi(string(parts[0])) * utils.FastAtoi(string(parts[1]))
	curr := 0

	for i, val := range strings.Fields(rest) {
		curr += idxMap[i] * utils.FastAtoi(val)
	}

	return curr <= lim
}
