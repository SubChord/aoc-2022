package main

import (
	"bufio"
	"os"
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

	grid := makeGrid(lines)
	part1(grid)
	part2(grid)
}

func makeGrid(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[i]))
		for j := range grid[i] {
			atoi, _ := strconv.Atoi(string(lines[i][j]))
			grid[i][j] = atoi
		}
	}
	return grid
}

func part1(grid [][]int) {
	c := 0
	for rowIdx, row := range grid {
		for colIdx, v := range row {

			// check if any value left is greater than v
			left := append([]int{}, row[:colIdx]...)
			if isMaxValue(v, left) {
				c++
				continue
			}

			// check if any value right is greater than v
			right := append([]int{}, row[colIdx+1:]...)
			// reverse right
			for i, j := 0, len(right)-1; i < j; i, j = i+1, j-1 {
				right[i], right[j] = right[j], right[i]
			}
			if isMaxValue(v, right) {
				c++
				continue
			}

			// check if any value above is greater than v
			above := []int{}
			for i := 0; i < rowIdx; i++ {
				above = append(above, grid[i][colIdx])
			}
			if isMaxValue(v, above) {
				c++
				continue
			}

			// check if any value below is greater than v
			below := []int{}
			for i := rowIdx + 1; i < len(grid); i++ {
				below = append(below, grid[i][colIdx])
			}
			// reverse below
			for i, j := 0, len(below)-1; i < j; i, j = i+1, j-1 {
				below[i], below[j] = below[j], below[i]
			}
			if isMaxValue(v, below) {
				c++
				continue
			}
		}
	}
	println(c)
}

func isMaxValue(v int, vv []int) bool {
	for _, vvv := range vv {
		if v <= vvv {
			return false
		}
	}
	return true
}

func part2(grid [][]int) {
	maxScore := 0

	// make sure v is always a []int moving away from the tree location
	// eg. for the left part the first tree int []int is the tree directly to the left of the tree
	score := func(tree int, v []int) int {
		s := 0
		for _, vvv := range v {
			if vvv < tree {
				s++
				continue
			}
			s++
			break
		}
		return s
	}

	for rowIdx, row := range grid {
		for colIdx, tree := range row {
			left := append([]int{}, row[:colIdx]...)
			// reverse left
			for i, j := 0, len(left)-1; i < j; i, j = i+1, j-1 {
				left[i], left[j] = left[j], left[i]
			}

			leftScore := score(tree, left)

			// check if any value right is greater than tree
			right := append([]int{}, row[colIdx+1:]...)
			rightScore := score(tree, right)

			// check if any value above is greater than tree
			above := []int{}
			for i := rowIdx - 1; i >= 0; i-- {
				above = append(above, grid[i][colIdx])
			}

			aboveScore := score(tree, above)

			// check if any value below is greater than tree
			below := []int{}
			for i := rowIdx + 1; i < len(grid); i++ {
				below = append(below, grid[i][colIdx])
			}

			belowScore := score(tree, below)

			score := leftScore * rightScore * aboveScore * belowScore
			if score > maxScore {
				maxScore = score
			}
		}
	}
	println(maxScore)
}
