package main

import (
	"bufio"
	"fmt"
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

func getCrates() [][]string {

	//[S]                 [T] [Q]
	//[L]             [B] [M] [P]     [T]
	//[F]     [S]     [Z] [N] [S]     [R]
	//[Z] [R] [N]     [R] [D] [F]     [V]
	//[D] [Z] [H] [J] [W] [G] [W]     [G]
	//[B] [M] [C] [F] [H] [Z] [N] [R] [L]
	//[R] [B] [L] [C] [G] [J] [L] [Z] [C]
	//[H] [T] [Z] [S] [P] [V] [G] [M] [M]
	// 1   2   3   4   5   6   7   8   9

	stacks := [][]string{
		{"S", "L", "F", "Z", "D", "B", "R", "H"},
		{"R", "Z", "M", "B", "T"},
		{"S", "N", "H", "C", "L", "Z"},
		{"J", "F", "C", "S"},
		{"B", "Z", "R", "W", "H", "G", "P"},
		{"T", "M", "N", "D", "G", "Z", "J", "V"},
		{"Q", "P", "S", "F", "W", "N", "L", "G"},
		{"R", "Z", "M"},
		{"T", "R", "V", "G", "L", "C", "M"},
	}

	// revert the stacks
	for i, stack := range stacks {
		stacks[i] = revert(stack)
	}

	return stacks
}

type instruction struct {
	n    int
	from int
	to   int
}

func revert(stack []string) []string {
	ret := []string{}
	for i := len(stack) - 1; i >= 0; i-- {
		ret = append(ret, stack[i])
	}
	return ret
}

func main() {
	lines, err := readLines("inp")
	if err != nil {
		panic(err)
	}

	// lines to instructions
	instructions := []instruction{}
	for _, line := range lines {
		// move 2 from 2 to 4
		reg := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
		matches := reg.FindStringSubmatch(line)
		n, _ := strconv.Atoi(matches[1])
		from, _ := strconv.Atoi(matches[2])
		to, _ := strconv.Atoi(matches[3])
		instructions = append(instructions, instruction{n, from, to})
	}

	part1(instructions)
	part2(instructions)
}

func part1(instructions []instruction) {
	stacks := getCrates()
	for _, instruction := range instructions {
		from := instruction.from - 1
		to := instruction.to - 1
		n := instruction.n
		stacks[to] = append(stacks[to], revert(stacks[from][len(stacks[from])-n:])...)
		stacks[from] = stacks[from][:len(stacks[from])-n]
	}

	firstCrates := ""
	for _, stack := range stacks {
		firstCrates += stack[len(stack)-1]
	}

	fmt.Println(firstCrates)
}

func part2(instructions []instruction) {
	stacks := getCrates()
	for _, instruction := range instructions {
		from := instruction.from - 1
		to := instruction.to - 1
		n := instruction.n

		// without revert
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-n:]...)
		stacks[from] = stacks[from][:len(stacks[from])-n]
	}

	firstCrates := ""
	for _, stack := range stacks {
		firstCrates += stack[len(stack)-1]
	}

	fmt.Println(firstCrates)
}
