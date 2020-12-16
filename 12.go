package main

import (
	"fmt"
	"math"
	"strconv"
)

func modLikePython(d, m int) int {
	var res int = d % m
	if ((res < 0 && m > 0) || (res > 0 && m < 0)) {
		return res + m
	}
	return res
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func adventDay12A(path string) {
	strs := readLines(path)

	direction := 0

	east := 0
	north := 0

	for _,str := range strs {
		amount,err := strconv.Atoi(str[1:])
		if err != nil {
			panic(err)
		}
		switch str[0] {
		case 'F':
			switch direction {
			case 0:
				east += amount
			case 1:
				north += amount
			case 2:
				east -= amount
			case 3:
				north -= amount
			}
		case 'E':
			east += amount
		case 'N':
			north += amount
		case 'W':
			east -= amount
		case 'S':
			north -= amount
		case 'L':
			direction = modLikePython(direction + amount/90, 4)
		case 'R':
			direction = modLikePython(direction - amount/90, 4)
		}
	}
	fmt.Printf("magnitude is %d\n", intAbs(east) + intAbs(north))
}

func rotateCoordinates(x, y, degrees int) (int, int) {
	rads := float64(degrees)/360 * 2 * math.Pi
	Xi := float64(x)
	Yi := float64(y)
	Xf := math.Cos(rads)*Xi - math.Sin(rads)*Yi
	Yf := math.Sin(rads)*Xi + math.Cos(rads)*Yi
	return int(math.Round(Xf)), int(math.Round(Yf))
}

func adventDay12B(path string) {
	strs := readLines(path)

	wayPointNorth := 1
	wayPointEast := 10

	east := 0
	north := 0

	for _,str := range strs {
		amount,err := strconv.Atoi(str[1:])
		if err != nil {
			panic(err)
		}
		switch str[0] {
		case 'F':
			east += amount*wayPointEast
			north += amount*wayPointNorth
		case 'E':
			wayPointEast += amount
		case 'N':
			wayPointNorth += amount
		case 'W':
			wayPointEast -= amount
		case 'S':
			wayPointNorth -= amount
		case 'L':
			wayPointNorth, wayPointEast = rotateCoordinates(wayPointNorth, wayPointEast, -amount)
		case 'R':
			wayPointNorth, wayPointEast = rotateCoordinates(wayPointNorth, wayPointEast, amount)
		}
	}
	fmt.Printf("magnitude is %d\n", intAbs(east) + intAbs(north))
}

