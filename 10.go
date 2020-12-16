package main

import (
	"fmt"
	"sort"
)

func adventDay10A(path string) {
	strs := readLines(path)
	jolts := toInts(strs)
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	deltas := make(map[int]int)

	currentJolt := 0
	for _,adapter := range jolts {
		delta := adapter - currentJolt
		deltas[delta]++
		currentJolt = adapter
	}
	fmt.Printf("%d\n", deltas[1]*deltas[3])
}

func adventDay10B(path string) {
	strs := readLines(path)
	jolts := toInts(strs)
	jolts = append(jolts, 0)
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	ways := make([]int, len(jolts))

	for i,adapter := range jolts {
		thisWays := 0
		if i >= 3 && adapter - jolts[i-3] <= 3 {
			thisWays += ways[i-3]
		}
		if i >= 2 && adapter - jolts[i-2] <= 3 {
			thisWays += ways[i-2]
		}
		if i >= 1 && adapter - jolts[i-1] <= 3 {
			thisWays += ways[i-1]
		}
		if i == 0 {
			thisWays = 1
		}
		ways[i] = thisWays
	}
	//fmt.Printf("%v\n", jolts)
	//fmt.Printf("%v\n", ways)
	fmt.Printf("%d ways\n", ways[len(ways)-1])
}

