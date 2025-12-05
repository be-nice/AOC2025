package day5

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	start int
	end   int
}

func checkIDs(s string, ranges []interval) int {
	counter := 0
	for _, el := range strings.Fields(s) {
		num, err := strconv.Atoi(el)
		if err != nil {
			log.Fatal("Invalid value to convert to int")
		}

		for _, inter := range ranges {
			if num < inter.start {
				break
			}

			if num >= inter.start && num <= inter.end {
				counter++
				break
			}
		}
	}

	return counter
}

func buildRanges(s string) []interval {
	var ranges []interval

	for _, line := range strings.Fields(s) {

		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}

		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Fatal("Invalid value to convert")
		}

		ranges = append(ranges, interval{start: start, end: end})
	}

	return mergeRanges(ranges)
}

func mergeRanges(intervals []interval) []interval {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	merged := []interval{intervals[0]}

	for _, curr := range intervals[1:] {
		last := &merged[len(merged)-1]

		if curr.start <= last.end+1 {
			if curr.end > last.end {
				last.end = curr.end
			}
		} else {
			merged = append(merged, curr)
		}
	}

	return merged
}
