// Package utils is used to read input from txt files
// in the most useful format for the task
package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadBytes(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to open files, on path %s", path)
	}

	return data
}

func ReadByteLines(path string) [][]byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file, on path %s", path)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	res := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Bytes()
		buf := make([]byte, len(line))
		copy(buf, line)
		res = append(res, buf)
	}

	return res
}

func ReadString(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to open file, on path %s", path)
	}

	return string(data)
}

func ReadStringLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file, on path %s", path)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	res := make([]string, 0)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res
}

func ReadIntLines(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file, on path %s", path)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	res := make([]int, 0)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		res = append(res, num)
	}

	return res
}
