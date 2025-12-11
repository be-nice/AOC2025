package day6

import (
	"fmt"
	"strconv"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadByteLines(path[f])
	res := 0

	parseInput(data)
	rotateLeft()

	for i := range len(rawLines) {
		colRes := 0
		if operations[len(operations)-1-i] == '*' {
			colRes = 1
		}

		line := constructByIndex(rawLines[i])
		for _, el := range line {
			n, _ := strconv.Atoi(el)
			if n == 0 {
				continue
			}

			switch operations[len(operations)-1-i] {
			case '+':
				colRes += n
			case '*':
				colRes *= n
			}
		}

		res += colRes
	}

	fmt.Printf("Day 6 | Part 2: %d\n", res)
}

func rotateLeft() {
	if len(rawLines) == 0 {
		return
	}

	rows := len(rawLines)
	cols := len(rawLines[0])
	rotated := make([][][]byte, cols)

	for i := range rotated {
		rotated[i] = make([][]byte, rows)
		for j := range rotated[i] {
			rotated[i][j] = rawLines[j][cols-1-i]
		}
	}

	rawLines = rotated
}

func constructByIndex(row [][]byte) []string {
	maxLen := 0
	for _, elem := range row {
		if len(elem) > maxLen {
			maxLen = len(elem)
		}
	}

	result := make([]string, maxLen)

	for i := 0; i < maxLen; i++ {
		var buf []byte
		for _, elem := range row {
			if i < len(elem) {
				if elem[i] == ' ' {
					continue
				}
				buf = append(buf, elem[i])
			}
		}
		result[i] = string(buf)
	}

	return result
}
