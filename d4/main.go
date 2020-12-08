package main

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
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

func checkYear(test string, min, max int64) bool {
	v, err := strconv.ParseInt(test, 10, 64)
	if err != nil {
		return false
	}
	return v >= min && v <= max
}

func checkHeight(text string) bool {
	re := regexp.MustCompile("(?m)(\\d+)(cm|in)")
	match := re.FindStringSubmatch(text)
	if match == nil {
		return false
	}
	fmt.Printf("%+v\n", match)
	h, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		panic(err)
	}
	switch match[2] {
	case "cm":
		return h >= 150 && h <= 193
	case "in":
		return h >= 59 && h <= 76
	default:
		panic("not cm or in")
	}
}

func checkHair(text string) bool {
	pattern := "^#[0-9a-f]{6}$"
	match, _ := regexp.MatchString(pattern, text)
	return match
}

func checkEye(text string) bool {
	eyes := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, v := range eyes {
		if v == text {
			return true
		}
	}
	return false
}

func checkPid(text string) bool {
	pattern := "^[0-9]{9}$"
	match, _ := regexp.MatchString(pattern, text)
	return match
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

func (p passport) ok2() bool {
	/*
		fmt.Printf("%t %t %t %t %t %t %t", checkYear(p.fields["byr"], 1920, 2002),
			checkYear(p.fields["iyr"], 2010, 2020),
			checkYear(p.fields["eyr"], 2020, 2030),
			checkHeight(p.fields["hgt"]),
			checkHair(p.fields["hcl"]),
			checkEye(p.fields["ecl"]),
			checkPid(p.fields["pid"]))
	*/
	return checkYear(p.fields["byr"], 1920, 2002) &&
		checkYear(p.fields["iyr"], 2010, 2020) &&
		checkYear(p.fields["eyr"], 2020, 2030) &&
		checkHeight(p.fields["hgt"]) &&
		checkHair(p.fields["hcl"]) &&
		checkEye(p.fields["ecl"]) &&
		checkPid(p.fields["pid"])
}

func main() {
	var (
		c, c2 int
		r     record
		p     passport
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
			if p.ok2() {
				c2++
			}

			fmt.Printf("OK: %+v ", p.ok())
			fmt.Printf("OK2: %+v ", p.ok2())
			fmt.Printf("p: %+v \n", p)
			if err == io.EOF {
				break
			}
			p.fields = make(map[string]string)
		}
		//fmt.Printf("%+v", r)
	}
	fmt.Printf("Good: %d \n, Good2: %d\n", c, c2)
}
