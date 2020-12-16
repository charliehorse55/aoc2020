package main

import (
	"fmt"
	"strings"
)

func adventDay15Helper(path string, nrounds int32) {
	strs := readLines(path)
	ns := strings.Split(strs[0], ",")
	n := toInts(ns)

	round := int32(1)
	spokenHistory := make([]int32, nrounds)
	for _,val := range n {
		spokenHistory[int32(val)] = round
		round++
	}

	toSpeak := int32(0)
	for round < nrounds {
		last := spokenHistory[toSpeak]
		spokenHistory[toSpeak] = round
		if last == 0 {
			toSpeak = 0
		} else {
			toSpeak = round - last
		}
		round++
	}
	fmt.Printf("%d\n", toSpeak)
}

func adventDay15A(path string) {
	adventDay15Helper(path, 2020)
}

func adventDay15B(path string) {
	adventDay15Helper(path, 30000000)
}

