package day2

import (
	"fmt"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadString(path[f])

	ranges, maxEnd, maxDigits := parseInput(data)
	sum := generateRepeatedIDs(ranges, maxEnd, maxDigits, 2, maxDigits)

	fmt.Printf("Day 2 | Part 2: %d\n", sum)
}
