package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

type Pair struct {
	A int
	B int
}

func main() {
	lines, err := readLines("inp")
	if err != nil {
		panic(err)
	}

	// lines to int ranges
	ranges := [][]Pair{}
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		couple := []Pair{}
		for _, pair := range pairs {
			split := strings.Split(pair, "-")
			a, _ := strconv.Atoi(split[0])
			b, _ := strconv.Atoi(split[1])
			couple = append(couple, Pair{a, b})
		}
		ranges = append(ranges, couple)
	}

	part1(ranges)
	part2(ranges)
}

func part1(ranges [][]Pair) {
	count := 0
	for _, rangePair := range ranges {
		p1 := rangePair[0]
		p2 := rangePair[1]

		// make sure p1 is the largest rage
		if p1.B-p1.A < p2.B-p2.A {
			p1, p2 = p2, p1
		}

		// check if p2 is a subset of p1
		if p1.A <= p2.A && p2.B <= p1.B {
			count += 1
		}
	}

	println(count)
}

func part2(ranges [][]Pair) {
	count := 0
	for _, rangePair := range ranges {
		p1 := rangePair[0]
		p2 := rangePair[1]

		m := map[int]bool{}
		for i := p1.A; i <= p1.B; i++ {
			m[i] = true
		}

		for i := p2.A; i <= p2.B; i++ {
			if m[i] {
				count += 1
				break
			}
		}
	}

	println(count)
}
