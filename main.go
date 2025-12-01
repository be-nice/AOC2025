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

	fmt.Printf("Running-> Day: %d | Part: %s | Test: %t\n", day, part, strings.Contains(os.Args[2], "t"))
	start := time.Now()
	DaySequence[day-1].Parts[part](os.Args[2])

	fmt.Println(time.Since(start))
}
