package main

import (
	"bufio"
	"os"
	"regexp"
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

type monkey struct {
	id      int
	items   []int
	op      func(x int) int
	testDiv int
	t       int
	f       int
}

func parseMonkey(lines []string) monkey {
	// Monkey 0:
	//  Starting items: 79, 98
	//  Operation: new = old * 19
	//  Test: divisible by 23
	//    If true: throw to monkey 2
	//    If false: throw to monkey 3

	matches := regexp.MustCompile(`^Monkey (\d+):$`).FindStringSubmatch(lines[0])
	id, _ := strconv.Atoi(matches[1])

	matches = regexp.MustCompile(`Starting items: (.*)$`).FindStringSubmatch(lines[1])
	items := []int{}
	for _, item := range strings.Split(matches[1], ", ") {
		i, _ := strconv.Atoi(item)
		items = append(items, i)
	}

	var op func(x int) int
	matches = regexp.MustCompile(`Operation: new = old (.) ([0-9a-z]+)$`).FindStringSubmatch(lines[2])
	if matches[2] == "old" {
		if matches[1] == "*" {
			op = func(x int) int {
				return x * x
			}
		}
		if matches[1] == "+" {
			op = func(x int) int {
				return x + x
			}
		}
	} else {
		atoi, _ := strconv.Atoi(matches[2])
		if matches[1] == "*" {
			op = func(x int) int {
				return x * atoi
			}
		}

		if matches[1] == "+" {
			op = func(x int) int {
				return x + atoi
			}
		}
	}

	matches = regexp.MustCompile(`Test: divisible by (\d+)$`).FindStringSubmatch(lines[3])
	testDiv, _ := strconv.Atoi(matches[1])

	matches = regexp.MustCompile(`If true: throw to monkey (\d+)$`).FindStringSubmatch(lines[4])
	t, _ := strconv.Atoi(matches[1])

	matches = regexp.MustCompile(`If false: throw to monkey (\d+)$`).FindStringSubmatch(lines[5])
	f, _ := strconv.Atoi(matches[1])

	return monkey{id, items, op, testDiv, t, f}
}

func main() {
	lines, err := readLines("inp")
	if err != nil {
		panic(err)
	}

	monkeys := []monkey{}
	for i := 0; i < len(lines); i += 7 {
		monkeys = append(monkeys, parseMonkey(lines[i:i+7]))
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].id < monkeys[j].id
	})
	part1(monkeys)

	monkeys = []monkey{}
	for i := 0; i < len(lines); i += 7 {
		monkeys = append(monkeys, parseMonkey(lines[i:i+7]))
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].id < monkeys[j].id
	})
	part2(monkeys)
}

func part1(monkeys []monkey) {
	inspections := map[int]int{}
	for i := 0; i < 20; i++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				newV := monkey.op(item)
				newV = newV / 3

				if newV%monkey.testDiv == 0 {
					monkeys[monkey.t].items = append(monkeys[monkey.t].items, newV)
				} else {
					monkeys[monkey.f].items = append(monkeys[monkey.f].items, newV)
				}
				inspections[i]++
			}
			monkeys[i].items = []int{}
		}
	}

	inspectionValues := []int{}
	for _, v := range inspections {
		inspectionValues = append(inspectionValues, v)
	}

	sort.Ints(inspectionValues)
	println(inspectionValues[len(inspectionValues)-1] * inspectionValues[len(inspectionValues)-2])
}

func part2(monkeys []monkey) {
	mod := 1
	for _, m := range monkeys {
		mod *= m.testDiv
	}

	inspections := map[int]int{}
	for i := 0; i < 10000; i++ {
		for i, monkey := range monkeys {
			for _, item := range monkey.items {
				newV := monkey.op(item)
				newV %= mod

				if newV%monkey.testDiv == 0 {
					monkeys[monkey.t].items = append(monkeys[monkey.t].items, newV)
				} else {
					monkeys[monkey.f].items = append(monkeys[monkey.f].items, newV)
				}
				inspections[i]++
			}
			monkeys[i].items = []int{}
		}
	}

	inspectionValues := []int{}
	for _, v := range inspections {
		inspectionValues = append(inspectionValues, v)
	}

	sort.Ints(inspectionValues)
	println(inspectionValues[len(inspectionValues)-1] * inspectionValues[len(inspectionValues)-2])
}
