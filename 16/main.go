package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	byteData, err := ioutil.ReadFile("inp")
	if err != nil {
		panic(err)
	}
	r := bytes.NewReader(byteData)
	input, err := readLines(r)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%v", input)

	fmt.Printf("%v\n", firstPart(input))
	// fmt.Printf("%v\n", secondPart(input))
}

type Move struct {
	Location string
	Time     int
	Flow     int
}

type PressureNode struct {
	Location      string
	TotalPressure int
	Time          int
	Visited       []string
}

func firstPart(input map[string]Valve) int {
	moves := make(map[string][]Move, 0)
	// for each valve, calculate the time it takes to move to all other valves which could be opened
	for k, v := range input {
		if v.Flowrate > 0 || k == "AA" {
			for k2, v2 := range input {
				if v2.Flowrate > 0 && k2 != k {
					time := shortestDistance(input, k, k2)
					moves[k] = append(moves[k], Move{Location: k2, Time: time + 1, Flow: v2.Flowrate})
				}
			}
		}
	}

	// start on AA
	// visited := make(map[string]interface{}, 0)
	max := 0
	c := 0
	queue := make([]PressureNode, 0)
	queue = append(queue, PressureNode{Location: "AA", TotalPressure: 0, Time: 0, Visited: []string{"AA"}})

	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]

		m := moves[pop.Location]

		for _, v := range m {
			if contains(v.Location, pop.Visited) {
				continue
			}
			c++
			if (pop.Time + v.Time) > 30 {
				continue
			}
			if !contains(v.Location, pop.Visited) {
				//newVisited := make([]string, 0)
				//for _, s := range pop.Visited {
				//	newVisited = append(newVisited, s)
				//}
				newVisited := pop.Visited
				newVisited = append(newVisited, v.Location)
				//newVisited = append(newVisited, v.Location)
				timeSpent := pop.Time + v.Time
				timeLeft := 30 - timeSpent
				newPressure := v.Flow * timeLeft
				totalPressure := newPressure + pop.TotalPressure
				// fmt.Printf("%v - %v\n", newVisited, totalPressure)
				//pres := pop.TotalPressure + ((30 - (pop.Time + v.Time)) * v.Flow)
				if totalPressure > max {
					max = totalPressure
				}
				queue = append(queue, PressureNode{Location: v.Location, TotalPressure: totalPressure, Time: timeSpent, Visited: newVisited})
			}

		}
	}

	// fmt.Println(c)

	return max
}

func contains(loc string, visited []string) bool {
	for _, v := range visited {
		if v == loc {
			return true
		}
	}

	return false
}

type Node struct {
	Valve    string
	Distance int
}

func shortestDistance(input map[string]Valve, from, to string) int {
	shortest := 50
	visited := make([]string, 0)

	queue := make([]Node, 0)
	queue = append(queue, Node{Valve: from, Distance: 0})

	for len(queue) > 0 {
		pop := queue[0]
		queue = queue[1:]

		if pop.Valve == to {
			if pop.Distance < shortest {
				shortest = pop.Distance
			}
		}

		visited = append(visited, pop.Valve)

		v := input[pop.Valve]
		for k := range v.Valves {
			if !contains(k, visited) {
				queue = append(queue, Node{Valve: k, Distance: pop.Distance + 1})
			}
		}
	}

	return shortest
}

func secondPart(input []string) int {
	return 0
}

type Valve struct {
	Flowrate int
	Valves   map[string]interface{}
}

func readLines(r io.Reader) (map[string]Valve, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	result := make(map[string]Valve, 0)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.ReplaceAll(line, "Valve ", "")
		s = strings.ReplaceAll(s, " has flow rate=", ", ")
		s = strings.ReplaceAll(s, "; tunnels lead to valves ", ", ")
		s = strings.ReplaceAll(s, "; tunnel leads to valve ", ", ")
		split := strings.Split(s, ", ")
		v := make(map[string]interface{}, 0)
		flowrate, _ := strconv.Atoi(split[1])
		for i := 2; i < len(split); i++ {
			v[split[i]] = nil
		}
		result[split[0]] = Valve{
			Flowrate: flowrate,
			Valves:   v,
		}
	}
	return result, scanner.Err()
}
