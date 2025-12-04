package day4

import (
	"fmt"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadByteLines(path[f])
	count := 0

	for {
		res := controller(data)
		if res == 0 {
			break
		}
		count += res
	}

	fmt.Println(count)
}

func controller(data [][]byte) int {
	count := 0

	for y := range len(data) {
		for x := range len(data[0]) {
			if data[y][x] == '@' {
				if countEdge(data, pos{y, x}, true) {
					count++
				}
			}
		}
	}

	return count
}
