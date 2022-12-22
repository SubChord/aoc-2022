package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

type robot struct {
	t        string
	ore      int
	clay     int
	obsidian int
}

type blueprint struct {
	id     int
	robots []robot
}

func robotFromString(s string) robot {
	// Each ore robot costs 4 ore.
	// Each obsidian robot costs 2 ore and 20 clay.

	re := regexp.MustCompile(`Each (\w+) robot costs (.*)`)
	matches := re.FindStringSubmatch(s)

	robot := robot{
		t: matches[1],
	}

	parts := strings.Split(matches[2], " and ")
	for _, part := range parts {
		v := strings.Split(part, " ")
		switch strings.Trim(v[1], ".") {
		case "ore":
			robot.ore = toInt(v[0])
		case "clay":
			robot.clay = toInt(v[0])
		case "obsidian":
			robot.obsidian = toInt(v[0])
		}
	}

	return robot
}

func blueprintFromString(s string) blueprint {
	//Blueprint 20: Each ore robot costs 2 ore. Eac...
	re := regexp.MustCompile(`Blueprint (\d+): (.*)`)
	matches := re.FindStringSubmatch(s)

	bp := blueprint{
		id: toInt(matches[1]),
	}

	robotString := matches[2]
	split := strings.Split(robotString, ". ")
	for _, s2 := range split {
		bp.robots = append(bp.robots, robotFromString(s2))
	}

	return bp
}

func toInt(s string) int {
	atoi, _ := strconv.Atoi(s)
	return atoi
}

func main() {
	lines, err := readLines("inp")
	if err != nil {
		panic(err)
	}

	blueprints := []blueprint{}
	for _, line := range lines {
		blueprints = append(blueprints, blueprintFromString(line))
	}

	part1(blueprints)
	//part2(readings)
}

func part1(blueprints []blueprint) {
	c := 0
	for _, b := range blueprints {
		c += b.id * simulate(b)
	}
	fmt.Println("part 1: " + strconv.Itoa(c))
}

func simulate(b blueprint) int {
	robots := map[string]int{
		"ore": 1,
	}

	resources := map[string]int{}
	timeLeft := 24

	for t := 0; t < timeLeft; t++ {

	}

}
