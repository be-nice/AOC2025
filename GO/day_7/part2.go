package day7

import (
	"bytes"
	"fmt"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadByteLines(path[f])
	cols := len(data[0])

	count := 0

	beams := make([]int, cols)
	beams[bytes.Index(data[0], []byte("S"))] = 1

	next := make([]int, cols)

	for _, line := range data[1:] {
		for i := range next {
			next[i] = 0
		}

		for col, n := range beams {
			if n == 0 {
				continue
			}

			switch line[col] {
			case '^':
				next[col+1] += n
				next[col-1] += n
			default:

				next[col] += n
			}
		}

		beams, next = next, beams
	}

	for _, n := range beams {
		count += n
	}

	fmt.Printf("Day 7 | Part 2: %d\n", count)
}
