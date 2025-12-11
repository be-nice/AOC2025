package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ValidateRunArgs(s []string) (int, string) {
	if len(s) == 0 {
		printUsage()
		os.Exit(0)
	}

	if len(s) == 1 {
		if strings.ToLower(s[0]) == "h" || strings.ToLower(s[0]) == "help" {
			printUsage()
			os.Exit(0)
		}

		if strings.ToLower(s[0]) == "all" {
			return -1, "all"
		}

		printUsage()
		os.Exit(0)
	}

	if len(s) != 2 {
		log.Fatalf("Expected two arguments, got <%d>\n", len(s))
	}

	return parseDayFlag(s[0]), parsePartFlag(s[1])
}

func parseDayFlag(s string) int {
	day, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Expected int as first argument")
	}
	if day < 1 || day > 12 {
		log.Fatalf("Expected first argument to be between <1-12>, got <%d>\n", day)
	}

	return day - 1
}

func parsePartFlag(s string) string {
	parts := strings.Split(strings.ToLower(s), "")
	if parts[0] != "a" && parts[0] != "b" || len(parts) > 2 {
		log.Fatalf("Expected second argument to be <'a' | 'b' | 'at' | 'bt'>, got <%s>\n", s)
	}

	if len(parts) == 2 && parts[1] != "t" {
		log.Fatalf("Expected second argument to be <'a' | 'b' | 'at' | 'bt'>, got <%s>\n", s)
	}

	return parts[0]
}

func printUsage() {
	fmt.Println("-----USAGE-----")
	fmt.Println("go run . <n>(1-12) < a | b | at | bt >")
}
