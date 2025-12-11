// Package day6 implements functions for the day6 in
// AOC 2025. Each day should be in it's own package.
package day6

import (
	"fmt"
	"strconv"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadByteLines(path[f])
	res := 0

	parseInput(data)

	for i := range len(values[0]) {
		colRes := 0
		for j := range len(values) {
			n, _ := strconv.Atoi(string(values[j][i]))

			switch operations[i] {
			case '+':
				colRes += n
			case '*':
				if colRes == 0 {
					colRes = 1
				}
				colRes *= n
			}
		}

		res += colRes
	}

	fmt.Printf("Day 6 | Part 1: %d\n", res)
}
