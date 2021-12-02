package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// hash of ingredient to list of recipes
// hash of allergen to list of recipe
// rank each in order
// delete
// for each allergen > 1 recipe
// find ingredient in all recipe and remove it
//
type stringIntMap map[string][]int
type intStringMap map[int][]string
type itemList []string

type indexType struct {
	ingToLine stringIntMap
	lineToIng intStringMap
	algToLine stringIntMap
	lineToAlg intStringMap
	algRank   []string
	ingRank   []string
}

func (idx *indexType) init() {
	idx.ingToLine, idx.algToLine = make(stringIntMap), make(stringIntMap)
	idx.lineToIng, idx.lineToAlg = make(intStringMap), make(intStringMap)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (idx indexType) matchLines(algLineList []int, ing string) bool {
	lineMap := make(map[int]bool)
	for _, v := range idx.ingToLine[ing] {
		lineMap[v] = true
	}
	for _, l := range algLineList {
		if _, ok := lineMap[l]; !ok {
			return false
		}
	}
	return true
}

func (idx *indexType) matchAndRemove() {
	allCandidates := make(map[string]bool)
	for _, a := range idx.algRank {
		lines := idx.algToLine[a]
		//fmt.Printf("alg: %s, lines: %d\n", a, len(lines))
		candidateIng := make(map[string]int)
		// add all ingredient, eliminate missing
		for _, l := range lines {
			// get ing
			if ing, ok := idx.lineToIng[l]; ok {
				for _, i := range ing {
					//fmt.Printf("alg %s, line %d, ing %s\n", a, l, i)
					candidateIng[i]++
				}
			} else {
				panic("bad line : " + fmt.Sprint(l))
			}
		}
		//fmt.Printf("alg: %s, candidateIng: %+v\n", a, candidateIng)
		for k, v := range candidateIng {
			if v != len(lines) {
				delete(candidateIng, k)
			} else {
				allCandidates[k] = true
			}
		}
		fmt.Printf("alg: %s , canidates: %+v\n", a, candidateIng)
	}
	// deleteAllCandidates
	cleanIng := make(map[string]int)
	fmt.Printf("len ingToLine : %d\n", len(idx.ingToLine))
	fmt.Printf("allCandidates: %+v\n", allCandidates)
	for k, v := range idx.ingToLine {
		cleanIng[k] = len(v)
	}
	for k := range allCandidates {
		delete(cleanIng, k)
	}
	fmt.Printf("len cleanIng : %d\n", len(cleanIng))
	sumClean := 0
	for _, v := range cleanIng {
		sumClean += v
	}
	fmt.Printf("sumClean : %d\n", sumClean)

	// count instances of cleanIng

}

/*
func (l stringIntMap) equals(r stringIntMap) bool {
	for k, v := range l {
		if rval, ok := r[k]; ok {
			// compare slices
			if !reflect.DeepEqual(v, rval) {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
*/

// order alg by rank
// for each recipe find ing with same lines
// if lines match, remove both ing and alg
// when out of alg, remaining ing should be the correct list

type sortByRank struct {
	itemRank *[]string
	itemMap  stringIntMap
}

func (s sortByRank) Len() int {
	return len(*s.itemRank)
}
func (s sortByRank) Swap(i, j int) {
	(*s.itemRank)[i], (*s.itemRank)[j] = (*s.itemRank)[j], (*s.itemRank)[i]
}
func (s sortByRank) Less(i, j int) bool {
	return len(s.itemMap[(*s.itemRank)[i]]) < len(s.itemMap[(*s.itemRank)[j]])
}

/*
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
*/

func keys(m stringIntMap) []string {
	r := make([]string, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func scanFile() {
	scanner := bufio.NewScanner(os.Stdin)
	var idx indexType
	idx.init()
	l := 0
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " (contains ")
		// ingredients
		ingriedients := strings.Split(parts[0][0:len(parts[0])], " ")
		allergens := strings.Split(parts[1][0:len(parts[1])-1], ", ")
		//fmt.Printf("ing: %+v, alg: %+v\n", ingriedients, allergens)
		for _, i := range ingriedients {
			idx.ingToLine[i] = append(idx.ingToLine[i], l)
			idx.lineToIng[l] = append(idx.lineToIng[l], i)
		}
		for _, a := range allergens {
			idx.algToLine[a] = append(idx.algToLine[a], l)
			idx.lineToAlg[l] = append(idx.lineToAlg[l], a)
		}
		l++
	}
	idx.ingRank = keys(idx.ingToLine)
	idx.algRank = keys(idx.algToLine)
	// sort both

	sort.Sort(sortByRank{&idx.ingRank, idx.ingToLine})
	sort.Sort(sortByRank{&idx.algRank, idx.algToLine})

	// go by alg rank, find ing with matching list and then remove
	//fmt.Printf("idx: %+v\n", idx)
	idx.matchAndRemove()
}

func main() {
	scanFile()
}
