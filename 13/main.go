package main

import (
	"bufio"
	"encoding/json"
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

type node struct {
	list   *[]node
	v      *int
	marker *bool
}

func (n node) isNumber() bool {
	return n.v != nil
}

func (n node) isList() bool {
	return n.list != nil
}

func (n node) allNumbers() bool {
	if n.isNumber() {
		return true
	}

	for _, node := range *n.list {
		if !node.isNumber() {
			return false
		}
	}

	return true
}

func main() {
	lines, err := readLines("inp")
	if err != nil {
		panic(err)
	}

	part1(lines)
	part2(lines)
}

func lineToNode(line string) node {
	// [1,[2,[3,[4,[5,6,7]]]],8,9]

	line = strings.TrimPrefix(line, "[")
	line = strings.TrimSuffix(line, "]")
	if line == "" {
		return node{
			list: &[]node{},
		}
	}

	atoi, err := strconv.Atoi(line)
	if err == nil {
		return node{
			v: &atoi,
		}
	}

	parts := strings.Split(line, ",")
	buff := ""
	nOpen := 0

	list := []node{}
	for _, part := range parts {
		nOpen += strings.Count(part, "[")
		nOpen -= strings.Count(part, "]")

		if nOpen == 0 {
			list = append(list, lineToNode(buff+part))
			buff = ""
		}

		if nOpen > 0 {
			buff += part + ","
		}
	}

	return node{
		list: &list,
	}
}

func lineToNode2(line string) node {
	var n interface{}
	err := json.Unmarshal([]byte(line), &n)
	if err != nil {
		panic(err)
	}

	return interfaceToNode(n)
}

func interfaceToNode(n interface{}) node {
	switch v := n.(type) {
	case []interface{}:
		var list []node
		for _, n := range v {
			list = append(list, interfaceToNode(n))
		}

		return node{
			list: &list,
		}
	case float64:
		x := int(v)
		return node{
			v: &x,
		}
	default:
		panic("Should not happen")
	}
}

func nodeCompare(a, b node) int {
	if a.isNumber() && b.isNumber() {
		v1 := *a.v
		v2 := *b.v

		if v1 == v2 {
			return 0
		}

		if v1 < v2 {
			return 1
		}

		return -1
	}

	if a.isList() && b.isList() {
		minLen := len(*a.list)
		if len(*b.list) < minLen {
			minLen = len(*b.list)
		}

		for i := 0; i < minLen; i++ {
			n1 := (*a.list)[i]
			n2 := (*b.list)[i]

			r := nodeCompare(n1, n2)
			if r != 0 {
				return r
			}
		}

		if len(*a.list) == len(*b.list) {
			return 0
		}

		if len(*a.list) < len(*b.list) {
			return 1
		}

		return -1
	}

	// If we get here, we know a and b are not the same type

	var r int
	if a.isNumber() && b.isList() {
		r = nodeCompare(node{list: &[]node{a}}, b)
	} else if a.isList() && b.isNumber() {
		r = nodeCompare(a, node{list: &[]node{b}})
	} else {
		panic("Should not happen")
	}

	return r
}

func part1(lines []string) {
	indices := []int{}
	j := 1
	for i := 0; i < len(lines); i += 3 {
		a := lineToNode2(lines[i])
		b := lineToNode2(lines[i+1])

		if nodeCompare(a, b) > 0 {
			indices = append(indices, j)
		}

		j++
	}

	sum := 0
	for _, idx := range indices {
		sum += idx
	}

	fmt.Printf("Part 1: %d\n", sum)
}

func part2(lines []string) {
	linesWithoutBlanks := []string{}

	for _, line := range lines {
		if line != "" {
			linesWithoutBlanks = append(linesWithoutBlanks, line)
		}
	}

	nodes := []node{}
	for _, line := range linesWithoutBlanks {
		nodes = append(nodes, lineToNode2(line))
	}

	n2 := lineToNode2("[[2]]")
	n2.marker = new(bool)
	n6 := lineToNode2("[[6]]")
	n6.marker = new(bool)

	nodes = append(nodes, n2)
	nodes = append(nodes, n6)

	sort.Slice(nodes, func(i, j int) bool {
		return nodeCompare(nodes[i], nodes[j]) > 0
	})

	indices := []int{}
	j := 1
	for i := 0; i < len(nodes); i++ {
		n := nodes[i]
		if n.marker != nil {
			indices = append(indices, j)
		}
		j++
	}

	multiply := indices[0] * indices[1]
	fmt.Printf("Part 2: %d\n", multiply)
}
