// Package day7 implements functions for the day7 in
// AOC 2025. Each day should be in it's own package.
package day7

import "aoc2025/types"

var Funcs = types.DayStruct{
	Parts: map[string]func(string){
		"a": Part1,
		"b": Part2,
	},
}

var path = map[string]string{
	"a":  "day_7/data/part1.txt",
	"b":  "day_7/data/part2.txt",
	"at": "day_7/data/test_part1.txt",
	"bt": "day_7/data/test_part2.txt",
}
