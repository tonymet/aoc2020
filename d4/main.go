package main

import (
	"fmt"
	"io"
)

/*

byr (Birth Year)
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
	_, ok1 := p.fields["byr"]
	_, ok2 := p.fields["iyr"]
	_, ok3 := p.fields["eyr"]
	_, ok4 := p.fields["hgt"]
	_, ok5 := p.fields["hcl"]
	_, ok6 := p.fields["ecl"]
	_, ok7 := p.fields["pid"]
	//_, ok8 := p.fields["cid"]
	return ok1 && ok2 && ok3 && ok4 && ok5 && ok6 && ok7
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
