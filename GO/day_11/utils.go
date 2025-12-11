package day11

import "strings"

func createMap(s []string) map[string]map[string]struct{} {
	adj := make(map[string]map[string]struct{})

	for _, line := range s {
		k, v := parseLine(line)
		adj[k] = v
	}

	return adj
}

func parseLine(s string) (string, map[string]struct{}) {
	adj := make(map[string]struct{})
	parts := strings.Split(s, ":")

	for _, el := range strings.Fields(parts[1]) {
		adj[el] = struct{}{}
	}

	return parts[0], adj
}
