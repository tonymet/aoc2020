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
	_ "sort"
)

// storage
// opt a hashtable a,b to len
// opt b

var (
	filetypes = map[string]fileparam{
		"sample.txt": fileparam{20, 10, 3},
		"input.txt":  fileparam{1000, 1000, 3},
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
	circuits []int64
	circ2    map[int][]coord
	// lookup circuit and increment
	cLookup map[coord]int
	oEdges  []edge
}

// type node struct {
// 	coord
// 	neighbors []coord
// }

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
	//fmt.Printf("%+v\n", allCoords)
	var (
		allEdges edges
		t        tracker
	)
	allEdges = make(edges, 0, 1e6)
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
	// set up initial circuits
	// show the top edges
	fmt.Printf("allCoords: %d\t allE: %d\n", len(allCoords), len(allEdges))
	cid := 0
	t.cLookup = make(map[coord]int)
	t.circuits = make([]int64, activeParam.records)
	t.circ2 = make(map[int][]coord)
	for _, v := range allCoords {
		t.cLookup[v] = cid
		t.circuits[cid]++
		t.circ2[cid] = append(t.circ2[cid], v)
		cid++
	}
	t.oEdges = make([]edge, 0, len(allEdges))
	for i := 0; i < activeParam.top; i++ {
		cur := heap.Pop(&allEdges).(edge)
		t.oEdges = append(t.oEdges, cur)
		vl, okl := t.cLookup[cur.l]
		vr, okr := t.cLookup[cur.r]
		if !okl || !okr {
			panic("oob")
		}
		if t.cLookup[cur.l] == t.cLookup[cur.r] {
			continue
		}
		if t.circuits[vl] > t.circuits[vr] {
			t.circuits[vl] += t.circuits[vr]
			t.circuits[vr] = 0
			t.cLookup[cur.r] = vl
			for _, v := range t.circ2[vr] {
				t.cLookup[v] = vl
			}
			t.circ2[vl] = append(t.circ2[vl], t.circ2[vr]...)
			t.circ2[vr] = make([]coord, 0)

		} else {
			t.circuits[vr] += t.circuits[vl]
			t.circuits[vl] = 0
			t.cLookup[cur.l] = vr
			for _, v := range t.circ2[vl] {
				t.cLookup[v] = vr
			}
			t.circ2[vr] = append(t.circ2[vr], t.circ2[vl]...)
			t.circ2[vl] = make([]coord, 0)
		}
	}
	_ = t.oEdges
	// sort
	sl := t.circuits
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] > sl[j]
	})
	circSlice := make([]int, len(t.circ2))
	for k, v := range t.circ2 {
		circSlice[k] = len(v)
	}
	sort.Slice(circSlice, func(i, j int) bool {
		return circSlice[i] > circSlice[j]
	})

	//fmt.Printf("ordered: %+x\n", sl)
	prod := int64(1)
	for i := 0; i < activeParam.productLimit; i++ {
		prod *= sl[i]
	}
	prod2 := int64(1)
	for i := 0; i < activeParam.productLimit; i++ {
		prod2 *= int64(circSlice[i])
	}
	fmt.Printf("prod: %d\n", prod)
	fmt.Printf("prod2: %d\n", prod)
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
	basename := path.Base(file)
	activeParam = filetypes[basename]
	switch part {
	case 2:
		part2(os.Stdin)
	default:
		part1(os.Stdin)
	}
}
