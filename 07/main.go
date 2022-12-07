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

	part1Tree := makeTree(lines)
	part1Result := part1(part1Tree)
	println(part1Result)

	part2Tree := makeTree(lines)
	part2Result := part2(part2Tree)
	println(part2Result)
}

func part1(n *node) int64 {
	sum := int64(0)
	for _, child := range n.children {
		if child.isDir {
			totalSize := child.totalSize()
			if totalSize < 100000 {
				sum += totalSize
			}
		}

		sum += part1(&child)
	}

	return sum
}

func part2(s *node) int64 {
	totalSystem := int64(70000000)
	required := int64(30000000)

	unused := totalSystem - s.totalSize()
	toFree := required - unused

	// find all directories with a size greater than toFree
	sizes := part2Walk(s, toFree)

	// sort sizes
	for i := 0; i < len(sizes); i++ {
		for j := 0; j < len(sizes); j++ {
			if sizes[i] < sizes[j] {
				sizes[i], sizes[j] = sizes[j], sizes[i]
			}
		}
	}

	return sizes[0]
}

func part2Walk(n *node, free int64) []int64 {
	sizes := []int64{}
	if n.isDir {
		for _, child := range n.children {
			if child.isDir && child.totalSize() > free {
				sizes = append(sizes, child.totalSize())
			}

			sizes = append(sizes, part2Walk(&child, free)...)
		}
	}

	return sizes
}

type node struct {
	name     string
	children map[string]node
	parent   *node
	isDir    bool
	size     int64
}

func (n *node) totalSize() int64 {
	if !n.isDir {
		return n.size
	}

	sum := int64(0)
	for _, child := range n.children {
		sum += child.totalSize()
	}

	return sum
}

func makeTree(s []string) *node {
	var rootNode node
	var currentNode *node
	for i := 0; i < len(s); i++ {
		line := s[i]

		// if line starts with $, then parse command
		if line[0] == '$' {
			parts := strings.Split(line, " ")
			switch parts[1] {
			case "cd":
				if currentNode == nil {
					rootNode = node{
						name:     parts[2],
						children: make(map[string]node),
						parent:   nil,
						isDir:    true,
					}
					currentNode = &rootNode
					continue
				}

				if parts[2] == ".." {
					if currentNode.parent != nil {
						currentNode = currentNode.parent
					}
					continue
				}

				if parts[2] != "" {
					// check if dir exists in currentNode children
					if _, ok := currentNode.children[parts[2]]; ok {
						n := currentNode.children[parts[2]]
						currentNode = &n
						continue
					}
				}
			case "ls":
				lsLines := []string{}
				for j := i + 1; j < len(s); j++ {
					if s[j][0] == '$' {
						i = j - 1
						break
					}
					lsLines = append(lsLines, s[j])
				}

				for _, lsLine := range lsLines {
					parts := strings.Split(lsLine, " ")
					if parts[0] == "dir" {
						// dir
						currentNode.children[parts[1]] = node{
							name:     parts[1],
							children: make(map[string]node),
							parent:   currentNode,
							isDir:    true,
						}
					} else {
						// file
						atoi, _ := strconv.Atoi(parts[0])
						currentNode.children[parts[1]] = node{
							name:   parts[1],
							parent: currentNode,
							isDir:  false,
							size:   int64(atoi),
						}
					}
				}
			}
		}
	}

	return &rootNode
}
