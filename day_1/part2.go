package day1

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadString(path[f])

	count := 0
	pointer := 50

	for _, el := range strings.Fields(data) {
		dir := el[0]
		dist, err := strconv.Atoi(el[1:])
		if err != nil {
			log.Fatalf("invalid line: %s", el)
		}

		count += dist / 100
		dist %= 100

		switch dir {
		case 'L':
			if pointer != 0 && dist >= pointer {
				count++
			}
			pointer = (pointer - dist + 100) % 100
		case 'R':
			if pointer != 0 && pointer+dist >= 100 {
				count++
			}
			pointer = (pointer + dist) % 100
		}
	}

	fmt.Println(count)
}
