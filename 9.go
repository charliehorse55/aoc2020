package main

import (
	"fmt"
	"sort"
)

func adventDay9A(path string) {
	strs := readLines(path)
	ints := toInts(strs)

	fifo := make([]int, 25)

	pos := 0

	//fill it in
	for i := 0; i < len(fifo); i++ {
		fifo[pos % len(fifo)] = ints[pos]
		pos++
	}

	for pos < len(ints) {
		toCheck := ints[pos]
		foundSol := false
	out:
		for j := 0; j < len(fifo); j++ {
			for k := j+1; k < len(fifo); k++ {
				a := fifo[(pos+j) % len(fifo)]
				b := fifo[(pos+k) % len(fifo)]
				if a + b == toCheck {
					foundSol = true
					break out
				}
			}
		}
		if !foundSol {
			fmt.Printf("%d can't be constructed with the previous %d entries\n", toCheck, len(fifo))
		}
		fifo[pos % len(fifo)] = ints[pos]

		pos++
	}

}

func adventDay9B(path string) {
	strs := readLines(path)
	ints := toInts(strs)

	fifo := make([]int, 25)

	pos := 0

	//fill it in
	for i := 0; i < len(fifo); i++ {
		fifo[pos % len(fifo)] = ints[pos]
		pos++
	}

	invalidVal := 0

	for pos < len(ints) {
		toCheck := ints[pos]
		foundSol := false
	out:
		for j := 0; j < len(fifo); j++ {
			for k := j+1; k < len(fifo); k++ {
				a := fifo[(pos+j) % len(fifo)]
				b := fifo[(pos+k) % len(fifo)]
				if a + b == toCheck {
					foundSol = true
					break out
				}
			}
		}
		if !foundSol {
			fmt.Printf("%d can't be constructed with the previous %d entries\n", toCheck, len(fifo))
			invalidVal = toCheck
		}
		fifo[pos % len(fifo)] = ints[pos]

		pos++
	}

	start := 0
	end := 0
	sum := 0

	for sum != invalidVal {
		if sum < invalidVal {
			sum += ints[end]
			end++
		}
		if sum > invalidVal {
			sum -= ints[start]
			start++
		}
	}
	sort.Ints(ints[start:end])
	fmt.Printf("range %d to %d, sum is %d\n", start, end, ints[start] + ints[end-1])

}

