package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"aoc2025/utils"
)

func main() {
	day, part := utils.ValidateRunArgs(os.Args[1:])

	start := time.Now()
	if part == "all" {
		fmt.Printf("Running-> Day: all | Part: all\n")
		for i := range len(DaySequence) {

			DaySequence[i].Parts["a"]("a")
			DaySequence[i].Parts["b"]("b")
		}
	} else {
		fmt.Printf("Running-> Day: %d | Part: %s | Test: %t\n", day, part, strings.Contains(os.Args[2], "t"))
		DaySequence[day].Parts[part](os.Args[2])
	}

	fmt.Println(time.Since(start))
}
