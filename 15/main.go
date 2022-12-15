package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
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

type reading struct {
	sensorX, sensorY int
	beaconX, beaconY int
	manhattan        int
}

func main() {
	lines, err := readLines("inp")
	if err != nil {
		panic(err)
	}

	readings := make([]reading, len(lines))
	for i, line := range lines {
		// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
		re := regexp.MustCompile(`Sensor at x=([-]?\d+), y=([-]?\d+): closest beacon is at x=([-]?\d+), y=([-]?\d+)`)
		matches := re.FindStringSubmatch(line)
		readings[i] = reading{
			sensorX: toInt(matches[1]),
			sensorY: toInt(matches[2]),
			beaconX: toInt(matches[3]),
			beaconY: toInt(matches[4]),
		}

		readings[i].manhattan = manhattanDistance(readings[i].sensorX, readings[i].sensorY, readings[i].beaconX, readings[i].beaconY)
	}

	part1(readings)
	part2(readings)
}

func part1(readings []reading) {
	minX, maxX := math.MaxInt, math.MinInt
	maxManhattan := math.MinInt
	for _, reading := range readings {
		if reading.sensorX < minX {
			minX = reading.sensorX
		}
		if reading.sensorX > maxX {
			maxX = reading.sensorX
		}

		manhattan := manhattanDistance(reading.sensorX, reading.sensorY, reading.beaconX, reading.beaconY)
		if manhattan > maxManhattan {
			maxManhattan = manhattan
		}
	}

	c := 0
	y := 2000000
	for x := minX - maxManhattan; x < maxX+maxManhattan; x++ {
		for _, r := range readings {
			// check if x,y is beacon
			if r.beaconX == x && r.beaconY == y {
				continue
			}

			manhattan := manhattanDistance(x, y, r.sensorX, r.sensorY)
			if manhattan <= r.manhattan {
				c++
				break
			}
		}
	}

	fmt.Printf("part 1: %v\n", c)
}

func part2(readings []reading) {
	for _, r := range readings {
		topX, topY := r.sensorX, r.sensorY-r.manhattan-1
		rightX, rightY := r.sensorX+r.manhattan+1, r.sensorY

		found, x, y := checkEdge(topX, topY, rightX, rightY, readings)
		if found {
			fmt.Printf("Part 2: %d\n", x*4000000+y)
			fmt.Printf("Found at %d,%d\n", x, y)
		}
	}
}

func checkEdge(fromX, fromY int, toX, toY int, readings []reading) (bool, int, int) {
	// for every point between from and to check if manhattan distance is less than or equal to any reading
	// if so, return false
	// if not, return true and its coordinates

	minV := 0
	maxV := 4000000

	leftX, leftY := fromX, fromY
	rightX, rightY := toX, toY
	if toX < fromX {
		leftX = toX
		leftY = toY
		rightX = fromX
		rightY = fromY
	}

	if leftY < rightY {
		// slope is positive
		y := max(leftY, minV)
		for x := max(leftX, minV); x <= min(rightX, maxV); x++ {
			found := false
			for _, r := range readings {
				manhattan := manhattanDistance(x, y, r.sensorX, r.sensorY)
				if manhattan <= r.manhattan {
					found = true
				}
			}
			if !found {
				return true, x, y
			}
			y++
			if y > min(rightY, maxV) {
				break
			}
		}
	} else {
		// slope is negative
		y := min(leftY, maxV)
		for x := max(leftX, minV); x <= min(rightX, maxV); x++ {
			found := false
			for _, r := range readings {
				manhattan := manhattanDistance(x, y, r.sensorX, r.sensorY)
				if manhattan <= r.manhattan {
					found = true
				}
			}
			if !found {
				return true, x, y
			}
			y--
			if y >= max(rightY, minV) {
				break
			}
		}
	}

	return false, 0, 0
}

func max(x int, v int) int {
	if x > v {
		return x
	}
	return v
}

func min(x int, v int) int {
	if x < v {
		return x
	}
	return v
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func grid(readings []reading) ([][]int, int, int) {
	grid := [][]int{}
	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	maxManhattan := math.MinInt
	for _, reading := range readings {
		if reading.sensorX < minX {
			minX = reading.sensorX
		}
		if reading.sensorY < minY {
			minY = reading.sensorY
		}
		if reading.sensorX > maxX {
			maxX = reading.sensorX
		}
		if reading.sensorY > maxY {
			maxY = reading.sensorY
		}

		manhattan := manhattanDistance(reading.sensorX, reading.sensorY, reading.beaconX, reading.beaconY)
		if manhattan > maxManhattan {
			maxManhattan = manhattan
		}

		minX -= maxManhattan
		minY -= maxManhattan
		maxX += maxManhattan
		maxY += maxManhattan
	}

	for y := minY; y < maxY-minX+maxManhattan; y++ {
		grid = append(grid, make([]int, maxX-minX+maxManhattan))
	}

	for _, reading := range readings {
		manhattan := manhattanDistance(reading.sensorX, reading.sensorY, reading.beaconX, reading.beaconY)
		// correct sensorX, sensorY, beaconX and beaconY for minY and minX
		sX := reading.sensorX - minX
		sY := reading.sensorY - minY
		bX := reading.beaconX - minX
		bY := reading.beaconY - minY

		grid[sY][sX] = -1
		grid[bY][bX] = -2

		for y := sY - manhattan; y <= sY+maxManhattan; y++ {
			for x := sX - manhattan; x <= sX+maxManhattan; x++ {
				if manhattanDistance(x, y, sX, sY) <= manhattan {
					if grid[y][x] >= 0 {
						grid[y][x]++
					}
				}
			}
		}
	}

	//fmt.Printf("minY: %v ----- maxY: %v\n", minY, maxY)
	//
	//r := minY
	//for _, row := range grid {
	//	fmt.Printf("%v\t", r)
	//	for _, cell := range row {
	//		if cell > 0 {
	//			print("#")
	//		} else if cell == -1 {
	//			print("S")
	//		} else if cell == -2 {
	//			print("B")
	//		} else {
	//			print(".")
	//		}
	//	}
	//	r++
	//	println()
	//}

	return grid, minX, minY
}
