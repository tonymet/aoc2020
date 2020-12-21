package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ticketType []int

type highLow struct {
	low, high int
}

type ticketFile struct {
	defs     map[string][]highLow
	myTicket []int
	tickets  [][]int
}

func scanFile() (ticketFile, error) {
	var tf ticketFile
	tf.defs = make(map[string][]highLow)
	tf.tickets = make([][]int, 0)

	modes := []string{"def", "header", "myticket", "header", "tickets"}
	m := 0
	//for m, i, err := 0, 0, error(nil); err != io.EOF; i++ {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//_, err = fmt.Scanln(&line)
		if line == "" {
			m++
			continue
		}
		switch modes[m] {
		case "def":
			tf.readDef(line)
		case "myticket":
			tf.myTicket = tf.readTicket(line)
		case "tickets":
			tf.tickets = append(tf.tickets, tf.readTicket(line))
		case "header":
			m++
			continue
		}
	}
	return tf, nil
}
func (tf *ticketFile) readTicket(line string) []int {
	r := make([]int, 0)
	nums := strings.Split(line, ",")
	for _, n := range nums {
		i, _ := strconv.ParseInt(n, 10, 32)
		r = append(r, int(i))
	}
	return r
}

func (tf *ticketFile) part1() {
	c := 0
	for _, t := range tf.tickets {
		invalid := tf.invalidValues(t)
		if len(invalid) == 0 {
			continue
		}
		if len(invalid) > 1 {
			panic("long")
		}
		c += invalid[0]
		fmt.Printf("%+v\n", invalid)
	}
	fmt.Printf("sum: %d", c)
}

func (tf *ticketFile) readDef(line string) {
	chunks := strings.Split(line, ": ")
	if len(chunks) < 2 {
		panic("chunks < 2: " + line)
	}
	def := chunks[0]
	ranges := strings.Split(chunks[1], " or ")
	for _, r := range ranges {
		var hl highLow
		highLowStrings := strings.Split(r, "-")
		l, _ := strconv.ParseInt(highLowStrings[0], 10, 32)
		h, _ := strconv.ParseInt(highLowStrings[1], 10, 32)
		hl.low, hl.high = int(l), int(h)
		tf.defs[def] = append(tf.defs[def], hl)
	}
	return
}

func (tf *ticketFile) invalidValues(t ticketType) ticketType {
	r := make(ticketType, 0)
outer:
	for _, v := range t {
		for _, highLowSlice := range tf.defs {
			if (v >= highLowSlice[0].low && v <= highLowSlice[0].high) ||
				(v >= highLowSlice[1].low && v <= highLowSlice[1].high) {
				continue outer
			}
		}
		r = append(r, v)
	}
	return r
}

func main() {
	tf, _ := scanFile()
	tf.part1()
	//fmt.Printf("%+v \n", tf)
}
