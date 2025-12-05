package day2

import "aoc2025/types"

var Funcs = types.DayStruct{
	Parts: map[string]func(string){
		"a": Part1,
		"b": Part2,
	},
}

var path = map[string]string{
	"a":  "day_2/data/part1.txt",
	"b":  "day_2/data/part2.txt",
	"at": "day_2/data/test_part1.txt",
	"bt": "day_2/data/test_part2.txt",
}
