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

	for _, line := range data[1:] {
		next := make([]int, cols)

		for col, n := range beams {
			if n == 0 {
				continue
			}

			if line[col] == '^' {
				if col-1 >= 0 {
					next[col-1] += n
				} else {
					count += n
				}

				if col+1 < cols {
					next[col+1] += n
				} else {
					count += n
				}
			} else {
				next[col] += n
			}
		}

		beams = next
	}

	for _, n := range beams {
		count += n
	}

	fmt.Println(count)
}
