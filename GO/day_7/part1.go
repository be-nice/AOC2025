package day7

import (
	"bytes"
	"fmt"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadByteLines(path[f])
	count := 0

	beams := make([]bool, len(data[0]))
	beams[bytes.Index(data[0], []byte("S"))] = true

	next := make([]bool, len(beams))

	for _, line := range data[1:] {
		for i := range next {
			next[i] = false
		}

		for col, ok := range beams {
			if !ok {
				continue
			}

			switch line[col] {
			case '^':
				count++
				next[col+1] = true
				next[col-1] = true
			default:
				next[col] = true
			}
		}

		beams, next = next, beams
	}

	fmt.Printf("Day 7 | Part 1: %d\n", count)
}
