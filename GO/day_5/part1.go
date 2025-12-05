// Package day5 implements functions for the day5 in
// AOC 2025. Each day should be in it's own package.
package day5

import (
	"fmt"
	"strings"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadString(path[f])

	parts := strings.Split(data, "\n\n")

	ranges := buildRanges(parts[0])
	count := checkIDs(parts[1], ranges)

	fmt.Println(count)
}
