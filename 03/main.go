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

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		// split line in 2 halves
		h1 := line[:len(line)/2]
		h2 := line[len(line)/2:]

		m1 := map[rune]bool{}
		for _, k := range h1 {
			m1[k] = true
		}

		for _, k := range h2 {
			if m1[k] {
				sum += runeToValue(k)
				break
			}
		}
	}

	println(sum)
}

func part2(lines []string) {
	sum := 0
	for i := 0; i < len(lines); i += 3 {
		threeLines := []string{lines[i], lines[i+1], lines[i+2]}

		// map lines to occurrence of runes
		threeLinesMap := []map[rune]bool{}
		for _, line := range threeLines {
			m := map[rune]bool{}
			for _, k := range line {
				m[k] = true
			}
			threeLinesMap = append(threeLinesMap, m)
		}

		// check if any rune is in all 3 lines
		runeCount := map[rune]int{}
		for _, m := range threeLinesMap {
			for k := range m {
				runeCount[k]++
			}
		}

		for k, v := range runeCount {
			if v == 3 {
				sum += runeToValue(k)
				break
			}
		}
	}

	println(sum)
}

func runeToValue(r rune) int {
	// lowercase 1-26
	// uppercase 27-52
	if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 1)
	}
	return int(r - 'A' + 27)
}
