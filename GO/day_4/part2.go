package day4

import (
	"fmt"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadByteLines(path[f])

	for y := range len(data) {
		for x := range len(data[0]) {
			if data[y][x] == '@' {
				countEdge(data, pos{y, x}, false)
			}
		}
	}

	count := len(toDelete)

	for len(toDelete) > 0 {
		curr := toDelete[len(toDelete)-1]
		toDelete = toDelete[:len(toDelete)-1]

		for _, d := range dirs {
			offset := pos{curr.y + d.y, curr.x + d.x}

			if _, ok := countMap[offset]; !ok {
				continue
			}

			countMap[offset]--
			if countMap[offset] < 4 {
				count++
				toDelete = append(toDelete, offset)
				delete(countMap, offset)
			}
		}
	}

	fmt.Println(count)
}
