// Package day1 implements functions for the dayx in
// AOC. Each day should be in it's own package.
package day1

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadString(path[f])

	count := 0
	pointer := 50

	for _, el := range strings.Fields(data) {
		dir := el[0]
		dist, err := strconv.Atoi(el[1:])
		if err != nil {
			log.Fatalf("%d%s", dist, el[1:])
		}

		switch dir {
		case 'L':
			pointer = (pointer - dist + 100) % 100
		case 'R':
			pointer = (pointer + dist) % 100
		}

		if pointer == 0 {
			count++
		}
	}

	fmt.Println(count)
}
