package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
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
	elves := map[int]int{}

	elfIdx := 0
	for _, line := range lines {
		if line == "" {
			elfIdx++
			continue
		}

		v, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		elves[elfIdx] += v
	}

	// print max elf
	max := 0
	maxIdx := 0
	for idx, v := range elves {
		if v > max {
			max = v
			maxIdx = idx
		}
	}

	println(maxIdx, max)
}

func part2(lines []string) {
	elves := map[int]int{}

	elfIdx := 0
	for _, line := range lines {
		if line == "" {
			elfIdx++
			continue
		}

		v, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		elves[elfIdx] += v
	}

	values := []int{}
	for _, v := range elves {
		values = append(values, v)
	}

	sort.Ints(values)

	// print sum max 3
	sum := 0
	for i := len(values) - 3; i <= len(values)-1; i++ {
		sum += values[i]
	}

	println(sum)
}
