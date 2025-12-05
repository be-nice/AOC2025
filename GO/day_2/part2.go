package day2

import (
	"fmt"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadString(path[f])

	ranges, maxEnd, maxDigits := parseInput(data)
	sum := generateRepeatedIDs(ranges, maxEnd, maxDigits, 2, maxDigits)

	fmt.Println(sum)
}
