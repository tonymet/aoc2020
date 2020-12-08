package main

import (
	"fmt"
	"strings"
)

type record struct {
	low, high int64
	query     string
	password  string
}

type counter struct {
	count  int
	count2 int
}

func (r record) isValid() bool {
	c := int64(strings.Count(r.password, r.query))
	return c >= r.low && c <= r.high
}

// isValid game2
func (r record) isValid2() bool {
	if len(r.password) == 0 {
		return false
	}
	return (string(r.password[r.low-1]) == r.query) != (string(r.password[r.high-1]) == r.query)
}

func readRec() (record, error) {
	var cur record
	_, err := fmt.Scanf("%d-%d %1s: %s\n", &cur.low, &cur.high, &cur.query, &cur.password)
	if err != nil {
		return record{}, err
	}
	return cur, err
}

func (c *counter) readAll() {
	var (
		r   record
		err error
	)
	for ; err == nil; r, err = readRec() {
		if r.isValid() {
			c.count++
		}
		if r.isValid2() {
			c.count2++
		}
		fmt.Printf("%d-%d, %s password: %s, g1 :%v, g2: %v\n", r.low, r.high, r.query, r.password, r.isValid(), r.isValid2())
	}

}

func main() {
	var c counter
	c.readAll()
	fmt.Printf("Good 1: %d\n Good2: %d\n", c.count, c.count2)
}
