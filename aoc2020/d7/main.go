package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var bagIndex struct {
	parentChild         map[string][]string
	childParent         map[string][]string
	parentChildContents map[string][]Content
}

type Bag struct {
	color    string
	contains []Bag
}

type Content struct {
	color string
	count uint64
}

func parseContents(contentsLine string) []Content {
	contents := make([]Content, 0)
	var re = regexp.MustCompile(`(?i)([0-9]+) ([a-zA-Z ]+) bag`)
	match := re.FindAllStringSubmatch(contentsLine, -1)
	for _, m := range match {
		var c Content
		color := m[2]
		count, _ := strconv.ParseInt(m[1], 10, 32)
		c.count = uint64(count)
		c.color = color
		contents = append(contents, c)
	}
	return contents
}
func scanLine() {
	bagIndex.parentChild = make(map[string][]string)
	bagIndex.childParent = make(map[string][]string)
	bagIndex.parentChildContents = make(map[string][]Content)
	var (
		bagType, contents string
	)
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile("([a-zA-Z ]+) bags contain ([0-9a-zA-Z, ]+)")
	for {
		text, err := reader.ReadString('\n')
		match := re.FindStringSubmatch(text)
		if err != nil {
			fmt.Println(err)
			break
		}
		bagType = match[1]
		contents = match[2]
		//fmt.Printf("bagType: %s, contents: %s\n", bagType, contents)
		carray := parseContents(contents)
		//fmt.Printf("parseContents: %+v\n", carray)
		bagIndex.parentChildContents[bagType] = carray
		for _, v := range carray {
			bagIndex.parentChild[bagType] = append(bagIndex.parentChild[bagType], v.color)
			bagIndex.childParent[v.color] = append(bagIndex.childParent[v.color], bagType)
		}
	}
	//fmt.Printf("bagIndex: %+v\n", bagIndex)
	//fmt.Printf("len bagIndex: %d\n", len(bagIndex.parentChild))
	//fmt.Printf("parentChildContents: %d\n", bagIndex.parentChildContents)
	fmt.Printf("childParent mirroredGold bagIndex: %+v\n", bagIndex.parentChildContents["mirrored gold"])
}

var seen map[string]int

func countOwnership(q []string) {
	for _, v := range q {
		seen[v]++
		//fmt.Printf("ownerCount +=%d\n", len(bagIndex.childParent[v]))
		//fmt.Printf("countOwnership %+v\n", bagIndex.childParent[v])
		countOwnership(bagIndex.childParent[v])
	}
}

func countChildBags(cur Content) uint64 {
	// for each query
	// multiply by factor
	// add to sum
	// add childdren as query and call recursively
	var curSum uint64
	fmt.Printf("q: %+v\n", cur)
	//fmt.Printf("ownerCount +=%d\n", len(bagIndex.childParent[v]))
	//fmt.Printf("total: %d\n", *total)
	children := bagIndex.parentChildContents[cur.color]
	if len(children) == 0 {
		curSum = uint64(cur.count)
	} else {
		var childSum uint64
		for _, child := range children {
			childSum += countChildBags(child)
		}
		curSum = cur.count*childSum + cur.count
	}
	fmt.Printf("cur count: %d ; cur Sum : %d\n", cur.count, curSum)
	return curSum
}

func main() {
	seen = make(map[string]int)
	scanLine()
	countOwnership([]string{"shiny gold"})
	fmt.Printf("ownerCount: %d\n", len(seen)-1)
	childCount := countChildBags(Content{"shiny gold", 1})
	fmt.Printf("countChildBags: %d\n", childCount-1)
}
