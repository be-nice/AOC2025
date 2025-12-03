package day1

import "aoc2025/types"

var Funcs = types.DayStruct{
	Parts: map[string]func(string){
		"a": Part1,
		"b": Part2,
	},
}

var path = map[string]string{
	"a":  "day_1/data/part1.txt",
	"b":  "day_1/data/part2.txt",
	"at": "day_1/data/test_part1.txt",
	"bt": "day_1/data/test_part2.txt",
}
