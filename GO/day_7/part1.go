package day7

import (
	"bytes"
	"fmt"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadByteLines(path[f])
	count := 0

	beams := make(map[int]struct{})
	beams[bytes.Index(data[0], []byte("S"))] = struct{}{}

	for _, line := range data[1:] {
		for i, el := range line {
			if el == '^' {
				if _, ok := beams[i]; ok {
					delete(beams, i)
					beams[i-1] = struct{}{}
					beams[i+1] = struct{}{}
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
