package main

import (
	"bufio"
	"os"
)

// read all lines from file
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("inp")
	if err != nil {
		panic(err)
	}

	part1(lines[0])
	part2(lines[0])
}

func part1(s string) {
	size := 4
	for i := size; i < len(s)-1; i++ {
		sub := s[i-size : i]
		m := map[string]bool{}
		for _, c := range sub {
			m[string(c)] = true
		}

		if len(m) == size {
			println(i)
			return
		}
	}
}

func part2(s string) {
	size := 14
	for i := size; i < len(s)-1; i++ {
		sub := s[i-size : i]
		m := map[string]bool{}
		for _, c := range sub {
			m[string(c)] = true
		}

		if len(m) == size {
			println(i)
			return
		}
	}
}
