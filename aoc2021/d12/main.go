package main

import (
	"fmt"
	"os"
	"strings"
)

var part = 1

const (
	TYPE_START = iota
	TYPE_MID
	TYPE_END
)

type edge struct {
	a, b string
}

type node struct {
	big      bool
	name     string
	next     []*node
	nodeType int
}

type tIndex = map[string]*node
type graph struct {
	index tIndex
}

func (g *graph) addEdge(curEdge edge) error {
	var aNode, bNode *node
	if v, ok := g.index[curEdge.a]; !ok {
		var newNode node
		newNode.name = curEdge.a
		newNode.big = big(newNode.name)
		g.index[curEdge.a] = &newNode
		aNode = &newNode
	} else {
		aNode = v
	}
	if v, ok := g.index[curEdge.b]; !ok {
		var newNode node
		newNode.name = curEdge.b
		newNode.big = big(newNode.name)
		g.index[curEdge.b] = &newNode
		bNode = &newNode
	} else {
		bNode = v
	}
	aNode.next = append(aNode.next, bNode)
	bNode.next = append(bNode.next, aNode)
	return nil
}

func big(n string) bool {
	return strings.ToUpper(n) == n
}

func parseAndSetup() {
	var g graph
	g.index = make(tIndex)

	for {
		var curEdge edge
		var curLine string
		n, err := fmt.Scanf("%s\n", &curLine)
		if n == 0 || err != nil {
			return
		}
		tokens := strings.Split(curLine, "-")
		curEdge.a, curEdge.b = tokens[0], tokens[1]
		fmt.Printf("Edge: %+v\n", curEdge)
		g.addEdge(curEdge)
		// index the graph
		// lookup
	}

}
func part1() {
	parseAndSetup()
	// sort the positions
	// binary search and take deltas
}
func part2() {
	parseAndSetup()
	// sort the positions
	// binary search and take deltas
}

func main() {
	if stdin := os.Getenv("STDIN"); len(stdin) != 0 {
		stdinFile, err := os.Open(stdin)
		if err != nil {
			panic(err)
		}
		os.Stdin = stdinFile

	}
	switch os.Getenv("PART") {
	case "2":
		part = 2
		part2()
	default:
		part1()
	}
}
