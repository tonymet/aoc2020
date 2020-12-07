package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func parseContents(contentsLine string) []string {
	contents := make([]string, 0)
	var re = regexp.MustCompile(`(?i)([0-9a-zA-Z ]+) bag`)
	match := re.FindAllStringSubmatch(contentsLine, -1)
	for _, m := range match {
		contents = append(contents, m[1])
	}
	return contents
}
func scanLine() {
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
		fmt.Printf("bagType: %s, contents: %s\n", bagType, contents)
		carray := parseContents(contents)
		fmt.Printf("parseContents: %+v\n", carray)
	}

}

func main() {
	scanLine()
	fmt.Printf("hello\n")

}
