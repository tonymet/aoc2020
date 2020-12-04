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
	count int
}

func (r record) isValid() bool {
	c := int64(strings.Count(r.password, r.query))
	return c >= r.low && c <= r.high
}

func readRec() (record, error) {
	var cur record
	var query string
	_, err := fmt.Scanf("%d-%d %s %s\n", &cur.low, &cur.high, &query, &cur.password)
	if err != nil {
		return record{}, err
	}
	cur.query = string(query[0])
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
	}

}

func main() {
	var c counter
	c.readAll()
	fmt.Printf("Good: %d\n", c.count)
}
