// Package day3 implements functions for the day3 in
// AOC 2025. Each day should be in it's own package.
package day3

import (
	"fmt"
	"strconv"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadByteLines(path[f])

	sum := 0

	for _, line := range data {
		curr := make([]byte, 2)

		for i, b := range line {
			if b > curr[0] && i < len(line)-1 {
				curr[0] = b
				curr[1] = line[i+1]
			} else if b > curr[1] {
				curr[1] = b
			}
		}

		num, _ := strconv.Atoi(string(curr))
		sum += num
	}

	fmt.Println(sum)
}
