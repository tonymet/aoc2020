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
const (
	STATE_START = iota
)

type tracker struct {
	path []*node
	seen graph
}

var gTrackers []*tracker

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
	state int
}

func (t tracker) copy() tracker {
	var newTracker tracker
	newTracker.path = make([]*node, len(t.path))
	newTracker.seen.index = make(tIndex)
	for i := range t.path {
		newTracker.path[i] = t.path[i]
	}
	for k := range t.seen.index {
		newTracker.seen.index[k] = t.seen.index[k]
	}
	return newTracker
}

func (t tracker) String() string {
	names := make([]string, len(t.path))
	for i := range t.path {
		names[i] = t.path[i].name
	}
	return strings.Join(names, "->")
}

func (t *tracker) appendGlobal() {
	gTrackers = append(gTrackers, t)
}

func (t *tracker) init() {
	t.path = make([]*node, 0)
	t.seen.index = make(tIndex)
}

func (t *tracker) track(n *node) {
	t.path = append(t.path, n)
	t.seen.index[n.name] = n
}
func (t *tracker) seenNode(n *node) bool {
	if _, ok := t.seen.index[n.name]; ok {
		return true
	}
	return false
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

func (g *graph) traverse(curNode *node, t *tracker) {
	// traverse from start
	if curNode.name == "end" {
		t.track(curNode)
		return
	}
	if t.seenNode(curNode) {
		return
	}
	t.track(curNode)
	for i, nextNode := range curNode.next {
		if i > 0 {
			// copy and use new
			tmp := t.copy()
			t = &tmp
			t.appendGlobal()
		}
		g.traverse(nextNode, t)
	}
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
			break
		}
		tokens := strings.Split(curLine, "-")
		curEdge.a, curEdge.b = tokens[0], tokens[1]
		fmt.Printf("Edge: %+v\n", curEdge)
		g.addEdge(curEdge)
		// index the graph
		// lookup
	}
	gTrackers = make([]*tracker, 0)
	var curTracker tracker
	curTracker.init()
	curTracker.appendGlobal()
	g.traverse(g.index["start"], &curTracker)
	fmt.Printf("\ngTracker: %s\n", gTrackers)
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
