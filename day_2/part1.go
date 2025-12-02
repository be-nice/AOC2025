// Package day2 implements functions for the day2 in
// AOC 2025. Each day should be in it's own package.
package day2

import (
	"fmt"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadString(path[f])

	ranges, maxEnd, maxDigits := parseInput(data)
	sum := generateRepeatedIDs(ranges, maxEnd, maxDigits, 2, 2)

	fmt.Println(sum)
}
