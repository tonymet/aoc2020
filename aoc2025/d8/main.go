package main

import (
	"container/heap"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"sort"
)

// storage
// opt a hashtable a,b to len
// opt b

var (
	filetypes = map[string]fileparam{
		"sample.txt": {20, 10, 3},
		"input.txt":  {1000, 1000, 3},
	}
	activeParam fileparam
)

type fileparam struct {
	records, top, productLimit int
}

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
	seen     map[pair]bool
	circuits map[int][]coord
	// lookup circuit and increment
	cLookup  map[coord]int
	allEdges edges
}

func dist(p1, p2 coord) float64 {
	dx, dy, dz := p2.x-p1.x, p2.y-p1.y,
		p2.z-p1.z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

type edges []edge

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

func prepTracker(in io.Reader, t *tracker) {
	var allCoords coords
	allCoords = make(coords, 0, activeParam.records)
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
		//fmt.Printf("%d,%d,%d \n", cur.x, cur.y, cur.z)
	}
	t.allEdges = make(edges, 0, 1e6)
	t.seen = make(map[pair]bool)
	heap.Init(&t.allEdges)
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
				heap.Push(&t.allEdges, e)
				t.seen[e.pair] = true
			}
		}
	}
	fmt.Printf("allCoords: %d\t allE: %d\n", len(allCoords), len(t.allEdges))
	t.cLookup = make(map[coord]int)
	t.circuits = make(map[int][]coord)
	for cid, v := range allCoords {
		t.cLookup[v] = cid
		t.circuits[cid] = append(t.circuits[cid], v)
	}
}

func solve(in io.Reader) {
	var (
		t      tracker
		last   edge
		i      = 0
		cond   func() bool
		maxLen = 1
	)
	prepTracker(in, &t)
	switch part {
	case 1:
		cond = func() bool {
			return i < activeParam.top
		}
	case 2:
		cond = func() bool {
			return maxLen <= activeParam.records-1
		}
	default:
		panic("no part")
	}
	for i = 0; cond(); i++ {
		cur := heap.Pop(&t.allEdges).(edge)
		last = cur
		vl, okl := t.cLookup[cur.l]
		vr, okr := t.cLookup[cur.r]
		if !okl || !okr {
			panic("oob")
		}
		if t.cLookup[cur.l] == t.cLookup[cur.r] {
			continue
		}
		for _, v := range t.circuits[vr] {
			t.cLookup[v] = vl
		}
		t.circuits[vl] = append(t.circuits[vl], t.circuits[vr]...)
		t.circuits[vr] = make([]coord, 0)
		if len(t.circuits[vl]) > maxLen {
			maxLen = len(t.circuits[vl])
		}
	}
	switch part {
	case 1:
		fmt.Printf("part1 prod: %d\n", part1Prod(&t))
	case 2:
		fmt.Printf("part2 prod: %d\n", last.l.x*last.r.x)
	}
}

func part1Prod(t *tracker) int64 {
	circSlice := make([]int, len(t.circuits))
	for k, v := range t.circuits {
		circSlice[k] = len(v)
	}
	sort.Slice(circSlice, func(i, j int) bool {
		return circSlice[i] > circSlice[j]
	})
	part1Prod := int64(1)
	for i := 0; i < activeParam.productLimit; i++ {
		part1Prod *= int64(circSlice[i])
	}
	return part1Prod
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
	activeParam = filetypes[path.Base(file)]
	solve(os.Stdin)
}
