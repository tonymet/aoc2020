package main

import (
	"container/heap"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	_ "sort"
)

// storage
// opt a hashtable a,b to len
// opt b

type coord struct {
	x, y, z int
}

func (a coord) Eq(b coord) bool {
	return a.x == b.x && a.y == b.y && a.z == b.z
}

type coords = []coord
type edge struct {
	pair
	dist float64
}
type pair struct {
	l, r coord
}

func (e edge) String() string {
	return fmt.Sprintf("l:{%d,%d,%d}, r:{%d,%d,%d} , dist: %f", e.l.x, e.l.y, e.l.z, e.r.x, e.r.y, e.r.z, e.dist)
}

type tracker struct {
	seen map[pair]bool
	e    *edges
}

func dist(p1, p2 coord) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	dz := p2.z - p1.z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

type edges []edge

// heap

func (g edges) Contains(c coord) (seen bool) {
	seen = false
	for _, e := range g {
		if e.l.Eq(c) {
			return true
		}
	}
	return
}

func (g edges) Len() int {
	return len(g)
}

func (g edges) Less(i, j int) bool {
	return g[i].dist < g[j].dist
}

func (e edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e *edges) Push(x any) {
	item := x.(edge)
	*e = append(*e, item)
}

func (e *edges) Pop() any {
	tmp := *e
	item := tmp[len(*e)-1]
	*e = tmp[0 : len(tmp)-1]
	return item
}

func part2(in io.Reader) {
	fmt.Printf("part2 not implemented\n")
}

func part1(in io.Reader) {
	var allCoords coords
	allCoords = make(coords, 0, 1000)
	for {
		var cur coord
		n, err := fmt.Fscanf(in, "%d,%d,%d\n", &cur.x, &cur.y, &cur.z)
		if err == io.EOF {
			break
		}
		if n < 3 || err != nil {
			panic(err)
		}
		allCoords = append(allCoords, cur)
		fmt.Printf("%d,%d,%d \n", cur.x, cur.y, cur.z)
	}
	fmt.Printf("%+v\n", allCoords)
	var (
		allEdges edges
		t        tracker
	)
	allEdges = make(edges, 0, 1000)
	t.e = &allEdges
	t.seen = make(map[pair]bool)
	heap.Init(&allEdges)
	for i := 0; i < len(allCoords); i++ {
		for j := 0; j < len(allCoords); j++ {
			if allCoords[i].Eq(allCoords[j]) {
				continue
			}
			var e edge
			if dist(coord{0, 0, 0}, allCoords[i]) > dist(coord{0, 0, 0}, allCoords[j]) {
				e.l, e.r = allCoords[i], allCoords[j]
			} else {
				e.l, e.r = allCoords[j], allCoords[i]
			}
			e.dist = dist(e.l, e.r)
			if _, ok := t.seen[e.pair]; ok {
				continue

			} else {
				heap.Push(&allEdges, e)
				t.seen[e.pair] = true
			}
		}
	}
	_ = allEdges
	// show the top edges
	fmt.Printf("allE: %+x\n", len(allEdges))
	// pop 20
	for i := 0; i < 20; i++ {
		cur := heap.Pop(&allEdges)
		fmt.Printf("%s\n", cur)
	}
}

var (
	part   int
	file   string
	silent bool
)

func init() {
	flag.IntVar(&part, "p", 1, "which exercise part?")
	flag.StringVar(&file, "f", "", "which exercise part?")
	flag.BoolVar(&silent, "s", false, "silent?")

}

func main() {
	flag.Parse()
	if file != "" {
		var err error
		if os.Stdin, err = os.Open(file); err != nil {
			panic(err)
		}
	}
	switch part {
	case 2:
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}
