package day6

import (
	"bytes"
)

var (
	operations = []byte{}
	values     = [][][]byte{}
	rawLines   = [][][]byte{}
)

func parseInput(data [][]byte) {
	colWidth := 0
	colWidths := make([]int, 0)
	for _, el := range data[len(data)-1][1:] {
		colWidth++
		if el != ' ' {
			colWidths = append(colWidths, colWidth)
			colWidth = 0
		}
	}
	ops := bytes.Fields(data[len(data)-1])
	for _, o := range ops {
		operations = append(operations, o[0])
	}

	for _, line := range data[:len(data)-1] {
		values = append(values, bytes.Fields(line))
		rawLines = append(rawLines, splitByColWidths(line, colWidths))
	}
}

func splitByColWidths(line []byte, colWidths []int) [][]byte {
	result := make([][]byte, 0, len(colWidths))
	pos := 0

	for _, w := range colWidths {
		if w <= 0 {
			result = append(result, []byte{})
			continue
		}

		end := min(pos+w, len(line))

		seg := make([]byte, w)
		for i := range seg {
			seg[i] = ' '
		}

		if pos < len(line) {
			copy(seg, line[pos:end])
		}

		result = append(result, seg)
		pos += w
	}

	if pos < len(line) {
		extra := make([]byte, len(line)-pos)
		copy(extra, line[pos:])
		result = append(result, extra)
	}

	return result
}
