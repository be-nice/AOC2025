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

	for _, line := range data[1:] {
		next := make([]bool, len(beams))
		for col, ok := range beams {
			if !ok {
				continue
			}

			if line[col] == '^' {
				if col-1 >= 0 {
					next[col-1] = true
				}

				if col+1 < len(line) {
					next[col+1] = true
				}

				count++
			} else {
				next[col] = true
			}
		}

		beams = next
	}

	fmt.Println(count)
}
