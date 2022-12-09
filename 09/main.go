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

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	seen := make(map[string]bool)

	headX, headY := 0, 0
	tailX, tailY := 0, 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		dir := parts[0]
		dist, _ := strconv.Atoi(parts[1])

		switch dir {
		case "R":
			for i := 0; i < dist; i++ {
				headX++
				if abs(headX-tailX) > 1 {
					tailX++
					tailY = headY
				}
				seen[strconv.Itoa(tailX)+","+strconv.Itoa(tailY)] = true
			}
		case "L":
			for i := 0; i < dist; i++ {
				headX--
				if abs(headX-tailX) > 1 {
					tailX--
					tailY = headY
				}
				seen[strconv.Itoa(tailX)+","+strconv.Itoa(tailY)] = true
			}
		case "U":
			for i := 0; i < dist; i++ {
				headY++
				if abs(headY-tailY) > 1 {
					tailY++
					tailX = headX
				}
				seen[strconv.Itoa(tailX)+","+strconv.Itoa(tailY)] = true
			}
		case "D":
			for i := 0; i < dist; i++ {
				headY--
				if abs(headY-tailY) > 1 {
					tailY--
					tailX = headX
				}
				seen[strconv.Itoa(tailX)+","+strconv.Itoa(tailY)] = true
			}
		}
	}

	println(len(seen))
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type knot struct {
	x, y int
}

func part2(lines []string) {
	seen := make(map[string]bool)

	prevState := [10]knot{}
	knots := [10]*knot{}
	// init knots
	for i := range knots {
		knots[i] = &knot{}
		prevState[i] = knot{}
	}

	// move head and make rest of knots follow
	for _, line := range lines {
		for i := range knots {
			prevState[i] = *knots[i]
		}

		parts := strings.Split(line, " ")
		dir := parts[0]
		dist, _ := strconv.Atoi(parts[1])

		for i := 0; i < dist; i++ {
			switch dir {
			case "R":
				knots[0].x++
			case "L":
				knots[0].x--
			case "U":
				knots[0].y++
			case "D":
				knots[0].y--
			}

			for j := 1; j < len(knots); j++ {
				currentKnot := knots[j]
				prevKnot := knots[j-1]

				if touching(currentKnot, prevKnot) {
					continue
				}

				// check if we can move horizontally
				if currentKnot.x == prevKnot.x {
					if currentKnot.y > prevKnot.y {
						currentKnot.y--
					} else {
						currentKnot.y++
					}
					continue
				}

				// check if we can move vertically
				if currentKnot.y == prevKnot.y {
					if currentKnot.x > prevKnot.x {
						currentKnot.x--
					} else {
						currentKnot.x++
					}
					continue
				}

				// move diagonally
				if currentKnot.x > prevKnot.x {
					currentKnot.x--
				} else {
					currentKnot.x++
				}

				if currentKnot.y > prevKnot.y {
					currentKnot.y--
				} else {
					currentKnot.y++
				}
			}

			seen[strconv.Itoa(knots[9].x)+","+strconv.Itoa(knots[9].y)] = true
		}
	}

	println(len(seen))
}

func printGrid(knots [10]*knot) {
	for y := 10; y > -10; y-- {
		for x := -10; x < 10; x++ {
			found := false
			for i := range knots {
				if knots[i].x == x && knots[i].y == y {
					print(i)
					found = true
					break
				}
			}
			if !found {
				print(".")
			}
		}
		println()

	}
}

func touching(currentKnot *knot, prevKnot *knot) bool {
	if currentKnot.x == prevKnot.x && currentKnot.y == prevKnot.y {
		return true
	}

	touchingHorizontally := currentKnot.x == prevKnot.x && abs(currentKnot.y-prevKnot.y) == 1
	touchingVertically := currentKnot.y == prevKnot.y && abs(currentKnot.x-prevKnot.x) == 1
	touchingDiagonally := abs(currentKnot.x-prevKnot.x) == 1 && abs(currentKnot.y-prevKnot.y) == 1
	return touchingHorizontally || touchingVertically || touchingDiagonally
}
