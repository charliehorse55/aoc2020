package main

import (
	"fmt"
	"sort"
)

func adventDay5A(path string) {
	strs := readLines(path)
	maxSeatID := 0
	for _,seat := range strs {
		rowMax := 127
		rowMin := 0
		for i := 0; i < 7; i++ {
			halfway := ((rowMax-rowMin)/2) + rowMin
			if seat[i] == 'F' {
				rowMax = halfway
			} else {
				rowMin = halfway+1
			}
		}
		colMin := 0
		colMax := 7
		for j := 7; j < 10; j++ {
			halfway := ((colMax-colMin)/2) + colMin

			if seat[j] == 'L' {
				colMax = halfway
			} else {
				colMin = halfway+1
			}
		}
		seatID := 8*rowMax + colMax
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}
	fmt.Printf("max seat id: %d\n", maxSeatID)
}

func adventDay5B(path string) {
	strs := readLines(path)
	seatIDs := make([]int, len(strs))
	for k,seat := range strs {
		rowMax := 127
		rowMin := 0
		for i := 0; i < 7; i++ {
			halfway := ((rowMax-rowMin)/2) + rowMin
			if seat[i] == 'F' {
				rowMax = halfway
			} else {
				rowMin = halfway+1
			}
		}
		colMin := 0
		colMax := 7
		for j := 7; j < 10; j++ {
			halfway := ((colMax-colMin)/2) + colMin

			if seat[j] == 'L' {
				colMax = halfway
			} else {
				colMin = halfway+1
			}
		}
		seatID := 8*rowMax + colMax
		seatIDs[k] = seatID
	}
	sort.Ints(seatIDs)

	for i := 0; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1] == seatIDs[i]+2 {
			fmt.Printf("found missing seat: %d\n", seatIDs[i]+1)
		}
	}

}

