package main

import "fmt"

func adventDay1A(path string) {
	strs := readLines(path)
	ints := toInts(strs)
	for i := range ints {
		for j := i+1; j < len(ints); j++ {
			if ints[i] + ints[j] == 2020 {
				fmt.Printf("%d\n", ints[i]*ints[j])
			}
		}
	}
}

func adventDay1B(path string) {
	strs := readLines(path)
	ints := toInts(strs)
	for i := range ints {
		for j := i+1; j < len(ints); j++ {
			for k := range ints {
				if k == i || k == j {
					continue
				}
				if ints[i] + ints[j] + ints[k] == 2020 {
					fmt.Printf("%d\n", ints[i]*ints[j]*ints[k])
				}
			}
		}
	}
}
