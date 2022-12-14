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

func main() {
	lines, err := readLines("inp")
	if err != nil {
		panic(err)
	}

	part1(createGrid(lines))
	part2(createGrid(lines))
}

func createGrid(lines []string) [][]int {
	// find max Y in input
	maxY := 0
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		for _, part := range parts {
			split := strings.Split(part, ",")
			y := toInt(split[1])
			maxY = max(maxY, y)
		}
	}

	grid := [][]int{}
	for i := 0; i < maxY+2; i++ {
		grid = append(grid, make([]int, 1000))
	}
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		var prevX, prevY *int
		for _, part := range parts {
			split := strings.Split(part, ",")
			x, y := toInt(split[0]), toInt(split[1])
			if prevX == nil && prevY == nil {
				prevX, prevY = &x, &y
			}

			// draw line from prevX, prevY to x, y
			if *prevX == x {
				// vertical line
				for i := min(*prevY, y); i <= max(*prevY, y); i++ {
					grid[i][x]++
				}
				prevX, prevY = &x, &y
				continue
			}

			if *prevY == y {
				// horizontal line
				for i := min(*prevX, x); i <= max(*prevX, x); i++ {
					grid[y][i]++
				}
				prevX, prevY = &x, &y
				continue
			}

			panic("invalid line")
		}
	}
	return grid
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func part1(grid [][]int) {
	sourceX, sourceY := 500, 0
	sandX, sandY := sourceX, sourceY
	for {
		// try to drop sand by 1 Y

		// if we hit something we cant move down, we need to move diagonally left or right
		if grid[sandY+1][sandX] > 0 {
			// try to move left
			if grid[sandY+1][sandX-1] == 0 {
				sandX--
				sandY++
				continue
			}
			// try to move right
			if grid[sandY+1][sandX+1] == 0 {
				sandX++
				sandY++
				continue
			}

			// we cant move, add sand to grid and proceed
			grid[sandY][sandX] = 10
			sandX, sandY = sourceX, sourceY
		} else {

			sandY++
			if sandY == len(grid)-1 {
				break
			}
		}
	}

	// count all sand
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == 10 {
				count++
			}
		}
	}

	println(count)
}

func part2(grid [][]int) {
	sourceX, sourceY := 500, 0
	sandX, sandY := sourceX, sourceY
	for {
		// try to drop sand by 1
		// if we hit something we cant move down, we need to move diagonally left or right
		if grid[sandY+1][sandX] > 0 {
			// try to move left
			if grid[sandY+1][sandX-1] == 0 {
				sandX--
				sandY++

				// if we hit bottom we need to add sand to grid and proceed
				if sandY == len(grid)-1 {
					grid[sandY][sandX] = 10
					sandX, sandY = sourceX, sourceY
				}
				continue
			}
			// try to move right
			if grid[sandY+1][sandX+1] == 0 {
				sandX++
				sandY++

				// if we hit bottom we need to add sand to grid and proceed
				if sandY == len(grid)-1 {
					grid[sandY][sandX] = 10
					sandX, sandY = sourceX, sourceY
				}
				continue
			}

			// we cant move, add sand to grid and proceed
			grid[sandY][sandX] = 10
			sandX, sandY = sourceX, sourceY

			// if sand equals source break
			if sandX == sourceX && sandY == sourceY {
				break
			}
		} else {
			sandY++

			// if we hit bottom we need to add sand to grid and proceed
			if sandY == len(grid)-1 {
				grid[sandY][sandX] = 10
				sandX, sandY = sourceX, sourceY
			}
		}
	}

	// count all sand
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == 10 {
				count++
			}
		}
	}

	println(count)

}
