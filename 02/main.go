package main

import (
	"bufio"
	"os"
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

	rounds := [][]rune{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		round := []rune{}
		for _, s := range split {
			round = append(round, rune(s[0]))
		}
		rounds = append(rounds, round)
	}

	part1(rounds)
	part2(rounds)
}

func part1(rounds [][]rune) {
	score := 0
	for _, round := range rounds {
		playScore := int(round[1]) - int('W')
		combatScore := int(round[1]-23) - int(round[0])
		switch combatScore {
		case 0:
			score += 3 + playScore
		case 1, -2:
			score += 6 + playScore
		default:
			score += playScore
		}
	}

	println(score)
}

func part2(rounds [][]rune) {
	// X = lose, Y = draw, Z = win
	score := 0
	for _, round := range rounds {
		playScore := int(round[0]) - int('A') + 1
		switch round[1] {
		case 'X':
			v := (playScore - 1) % 3
			if v == 0 {
				v = 3
			}
			score += v
		case 'Y':
			score += 3 + playScore
		case 'Z':
			v := (playScore + 1) % 3
			if v == 0 {
				v = 3
			}
			score += 6 + v
		}
	}

	println(score)
}
