// Package day8ents functions for the day8 in
// AOC 2025. Each day should be in it's own package.
package day8

import "aoc2025/types"

var Funcs = types.DayStruct{
	Parts: map[string]func(string){
		"a": Part1,
		"b": Part2,
	},
}

var path = map[string]string{
	"a":  "day_8/data/part1.txt",
	"b":  "day_8/data/part2.txt",
	"at": "day_8/data/test_part1.txt",
	"bt": "day_8/data/test_part2.txt",
}
