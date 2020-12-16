package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func adventDay13A(path string) {
	strs := readLines(path)

	arrivedAt,err := strconv.Atoi(strs[0])
	if err != nil {
		panic(err)
	}

	busses := strings.Split(strs[1], ",")
	var busIDs []int
	for _,bus := range busses {
		busID,err := strconv.Atoi(bus)
		if err != nil {
			continue
		}
		busIDs = append(busIDs, busID)
	}

	best := math.MaxInt64
	bestID := 0
	for _,busID := range busIDs {
		departTime := ((arrivedAt + (busID -1))/busID) * busID
		if departTime < best {
			best = departTime
			bestID = busID
		}
	}
	fmt.Printf("%d\n", (best - arrivedAt) * bestID)

}

type bus struct {
	period int
	offset int
}

func firstTime(busses []bus) int {
	i := 1
	period := busses[0].period
	pos := 0
	for {
		pos += period
		if (pos + busses[i].offset) % busses[i].period == 0 {
			period *= busses[i].period
			i++
		}
		if i >= len(busses) {
			break
		}
	}
	return pos
}

func adventDay13B(path string) {
	strs := readLines(path)
	rawBusses := strings.Split(strs[1], ",")
	var busses []bus
	for i,rawBus := range rawBusses {
		busID,err := strconv.Atoi(rawBus)
		if err != nil {
			continue
		}
		busses = append(busses, bus{period:busID, offset: i})
	}
	fmt.Printf("%d\n", firstTime(busses))
}

