package main

import (
	"fmt"
	"os"
	"strconv"
)

type tCohort = []int

func part1Solution() (twoCount int) { return }

const DAYS = 80

var max_days int = DAYS

func sumMap(m map[int]int) (count int) {
	for _, v := range m {
		count += v
	}
	return
}
func parseAndSetup() {
	var (
		population      tCohort = make(tCohort, 0)
		days                    = make(map[int]int)
		populationCount int
	)

	for {
		var cur int
		n, err := fmt.Scanf("%d", &cur)
		if n != 1 || err != nil {
			break
		}
		population = append(population, cur)
		days[cur+1]++
	}
	populationCount += len(population)
	fmt.Printf("pop: %+v\n", population)
	fmt.Printf("days: %+v\n", days)

	// play the population

	for d := 0; d <= max_days; d++ {
		// reset next birthday
		days[d+7] += days[d]
		// set birthday for new fish
		days[d+9] += days[d]
		// count all new fish
		populationCount += days[d]
		days[d] = 0
	}
	fmt.Printf("pop: %+v\n", population)
	fmt.Printf("days: %+v\n", days)
	fmt.Printf("populationCount: %+v\n", populationCount)
	fmt.Printf("sumMap: %+v\n", sumMap(days))
}

func part1() {
	parseAndSetup()
}
func part2() {}

func main() {
	if stdin := os.Getenv("STDIN"); len(stdin) != 0 {
		stdinFile, err := os.Open(stdin)
		if err != nil {
			panic(err)
		}
		os.Stdin = stdinFile

	}
	if days_env := os.Getenv("DAYS"); len(days_env) != 0 {
		tmp, _ := strconv.ParseInt(days_env, 10, 32)
		max_days = int(tmp)
	}

	switch os.Getenv("PART") {
	case "2":
		part2()
	default:
		part1()
	}
}
