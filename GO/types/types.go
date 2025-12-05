// Package types implements custom types for AOC 2025
package types

type DayStruct struct {
	Parts map[string]func(string)
}
