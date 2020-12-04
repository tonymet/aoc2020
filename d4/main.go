package main

import (
	"fmt"
	"io"
)

var fields []string = []string{"byr",
	"byr", "iyr", "eyr", "hgt", "hcl",
	"ecl", "pid", /*"cid",*/
}

/*,
byr,(Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)
*/

type record struct {
	key, value string
}

type passport struct {
	fields map[string]string
}

func (p passport) ok() bool {
	for _, f := range fields {
		_, ok := p.fields[f]
		if !ok {
			return false
		}
	}
	return true
}

func main() {
	var (
		c int
		r record
		p passport
	)
	p.fields = make(map[string]string)
	for {
		_, err := fmt.Scanf("%3s:%s", &r.key, &r.value)
		p.fields[r.key] = r.value
		if err != nil {
			// new record do check
			//fmt.Printf("%+v", err)
			if p.ok() {
				c++
			}

			fmt.Printf("OK: %+v ", p.ok())
			fmt.Printf("p: %+v \n", p)
			if err == io.EOF {
				break
			}
			p.fields = make(map[string]string)
		}
		//fmt.Printf("%+v", r)
	}
	fmt.Printf("Good: %d \n", c)
}
