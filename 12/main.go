package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	grid := [][]int{}
	for i, line := range lines {
		grid = append(grid, []int{})
		for _, c := range line {
			if c == 'S' {
				grid[i] = append(grid[i], 0)
				continue
			}
			if c == 'E' {
				grid[i] = append(grid[i], 27)
				continue
			}

			grid[i] = append(grid[i], int(c)-int('a')+1)
		}
	}

	part1(grid)
	part2(grid)
}

func part1(grid [][]int) {
	startX, startY := 0, 0
	stopX, stopY := 0, 0
	for row, gridLine := range grid {
		for col, cell := range gridLine {
			if cell == 0 {
				startX, startY = col, row
			}
			if cell == 27 {
				stopX, stopY = col, row
			}
		}
	}

	flood := floodFill(grid, startY, startX)
	fmt.Printf("Part 1: %d\n", *flood[strconv.Itoa(stopY)+","+strconv.Itoa(stopX)])
}

func part2(grid [][]int) {
	paths := []int{}
	stopX, stopY := 0, 0
	for row, gridLine := range grid {
		for col, cell := range gridLine {
			if cell == 27 {
				stopX, stopY = col, row
			}
		}
	}

	for row, gridLine := range grid {
		for col, cell := range gridLine {
			if cell == 1 {
				flood := floodFill(grid, row, col)
				if flood[strconv.Itoa(stopY)+","+strconv.Itoa(stopX)] != nil {
					paths = append(paths, *flood[strconv.Itoa(stopY)+","+strconv.Itoa(stopX)])
				}
			}
		}
	}

	sort.Ints(paths)
	fmt.Printf("Part 2: %d\n", paths[0])
}

func floodFill(grid [][]int, startY int, startX int) map[string]*int {
	flood := make(map[string]*int)

	key := func(y, x int) string {
		return strconv.Itoa(y) + "," + strconv.Itoa(x)
	}

	// flood fill
	uniqueQ := []string{key(startY, startX)}
	n0 := 0
	flood[key(startY, startX)] = &n0

	for len(uniqueQ) > 0 {
		pop := uniqueQ[0]
		uniqueQ = uniqueQ[1:]

		split := strings.Split(pop, ",")
		row, _ := strconv.Atoi(split[0])
		col, _ := strconv.Atoi(split[1])
		v := grid[row][col]

		nextFloodValue := *flood[pop] + 1

		// up
		if row > 0 && flood[key(row-1, col)] == nil {
			if grid[row-1][col]-v <= 1 {
				uniqueQ = append(uniqueQ, key(row-1, col))
				flood[key(row-1, col)] = &nextFloodValue
			}
		}

		// down
		if row < len(grid)-1 && flood[key(row+1, col)] == nil {
			if grid[row+1][col]-v <= 1 {
				uniqueQ = append(uniqueQ, key(row+1, col))
				flood[key(row+1, col)] = &nextFloodValue
			}
		}

		// left
		if col > 0 && flood[key(row, col-1)] == nil {
			if grid[row][col-1]-v <= 1 {
				uniqueQ = append(uniqueQ, key(row, col-1))
				flood[key(row, col-1)] = &nextFloodValue
			}
		}

		// right
		if col < len(grid[0])-1 && flood[key(row, col+1)] == nil {
			if grid[row][col+1]-v <= 1 {
				uniqueQ = append(uniqueQ, key(row, col+1))
				flood[key(row, col+1)] = &nextFloodValue
			}
		}
	}

	return flood
}
