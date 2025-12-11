package day11

import "strings"

func createAdjMap(s []string) map[string][]string {
	adj := make(map[string][]string, len(s))

	for _, line := range s {
		parts := strings.SplitN(line, ":", 2)
		adj[parts[0]] = strings.Fields(parts[1])
	}

	return adj
}
