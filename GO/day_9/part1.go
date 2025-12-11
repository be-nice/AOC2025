package day9

import (
	"fmt"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadStringLines(path[f])
	maxRect := -1

	for i := range len(data) - 1 {
		for j := i + 1; j < len(data); j++ {
			ax, ay := parseNum(data[i])
			bx, by := parseNum(data[j])

			area := calcArea(ax, ay, bx, by)

			maxRect = max(maxRect, area)
		}
	}

	fmt.Printf("Day 9 | Part 1: %d\n", maxRect)
}
