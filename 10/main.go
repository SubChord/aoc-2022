package main

import (
	"bufio"
	"fmt"
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

func main() {
	lines, err := readLines("inp")
	if err != nil {
		panic(err)
	}

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	x := 1
	cycle := 0

	sum := 0

	checkCycle := func(c, x int) int {
		cycles := []int{20, 60, 100, 140, 180, 220}
		// check if c in cycles
		for _, v := range cycles {
			if c == v {
				fmt.Printf("x: %d, cycle: %d :: %v\n", x, c, c*x)
				return c * x
			}
		}
		return 0
	}

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		if parts[0] == "noop" {
			cycle++
			sum += checkCycle(cycle, x)
			continue
		}

		if parts[0] == "addx" {
			cycle++
			sum += checkCycle(cycle, x)
			cycle++
			sum += checkCycle(cycle, x)
			atoi, _ := strconv.Atoi(parts[1])
			x += atoi
		}
	}

	fmt.Printf("the sum is %v\n", sum)
	fmt.Println()
}

func part2(lines []string) {
	cycle := 0
	crt := []string{}
	for i := 0; i < 10; i++ {
		crt = append(crt, "")
	}
	x := 1

	drawCrt := func(cycle, x int, crt []string) {
		row := cycle / 40
		col := cycle % 40

		if x-1 <= col && x+1 >= col {
			crt[row] += "[]"
		} else {
			crt[row] += "  "
		}
	}

	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")
		if parts[0] == "noop" {
			drawCrt(cycle, x, crt)
			cycle++
			continue
		}

		if parts[0] == "addx" {
			drawCrt(cycle, x, crt)
			cycle++
			drawCrt(cycle, x, crt)
			cycle++
			atoi, _ := strconv.Atoi(parts[1])
			x += atoi
		}
	}

	for _, r := range crt {
		fmt.Printf("%v\n", r)
	}
}
