package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var bagIndex struct {
	parentChild map[string][]string
	childParent map[string][]string
}

type Bag struct {
	color    string
	contains []Bag
}

type Content struct {
	bag   Bag
	count int
}

func parseContents(contentsLine string) []Content {
	contents := make([]Content, 0)
	var re = regexp.MustCompile(`(?i)([0-9]+) ([a-zA-Z ]+) bag`)
	match := re.FindAllStringSubmatch(contentsLine, -1)
	for _, m := range match {
		var c Content
		color := m[2]
		count, _ := strconv.ParseInt(m[1], 10, 32)
		c.count = int(count)
		c.bag.color = color
		contents = append(contents, c)
	}
	return contents
}
func scanLine() {
	bagIndex.parentChild = make(map[string][]string)
	bagIndex.childParent = make(map[string][]string)
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
		for _, v := range carray {
			bagIndex.parentChild[bagType] = append(bagIndex.parentChild[bagType], v.bag.color)
			bagIndex.childParent[v.bag.color] = append(bagIndex.childParent[v.bag.color], bagType)
		}
	}
	//fmt.Printf("bagIndex: %+v\n", bagIndex)
	//fmt.Printf("len bagIndex: %d\n", len(bagIndex.parentChild))
	fmt.Printf("childParent shinyGold bagIndex: %+v\n", bagIndex.childParent["shiny gold"])
}

var seen map[string]int

func countOwnership(q []string) {
	for _, v := range q {
		seen[v]++
		fmt.Printf("ownerCount +=%d\n", len(bagIndex.childParent[v]))
		fmt.Printf("countOwnership %+v\n", bagIndex.childParent[v])
		countOwnership(bagIndex.childParent[v])
	}
}

func main() {
	seen = make(map[string]int)
	scanLine()
	countOwnership([]string{"shiny gold"})
	fmt.Printf("ownerCount: %d\n", len(seen)-1)

}
