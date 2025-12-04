// Package day4 implements functions for the day4 in
// AOC 2025. Each day should be in it's own package.
package day4

import (
	"fmt"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadByteLines(path[f])
	count := 0

	for y := range len(data) {
		for x := range len(data[0]) {
			if data[y][x] == '@' {
				if countEdge(data, pos{y, x}, false) {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
