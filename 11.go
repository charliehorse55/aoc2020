package main

import "fmt"

func copyDat(old [][]rune) [][]rune {
	new := make([][]rune, len(old))
	for i,row := range old {
		new[i] = make([]rune, len(row))
		copy(new[i], row)
	}
	return new
}

func adventDay11A(path string) {
	strs := readLines(path)

	height := len(strs)
	width := len(strs[0])
	dat := make([][]rune, height+2)
	dat[0] = make([]rune, width+2)
	dat[height+1] = make([]rune, width+2)
	for i,str := range strs {
		dat[i+1] = make([]rune, width+2)
		for j := 0; j < len(str); j++ {
			dat[i+1][j+1] = rune(str[j])
		}
	}

	//printLife(dat)

	nextDat := copyDat(dat)

	iterations := 0
	numChanged := 1
	for numChanged > 0 {
		numChanged = 0
		for i,row := range dat {
			if i == 0 || i == (height+1) {
				continue
			}
			for j,val := range row {
				if j == 0 || j == (width+1) {
					continue
				}
				totalOccupiedNearby := 0
				//count
				for x := -1; x < 2; x++ {
					for y := -1; y < 2; y++ {
						if x == 0 && y == 0 {
							continue
						}
						if dat[i+x][j+y] == '#' {
							totalOccupiedNearby++
						}
					}
				}

				if totalOccupiedNearby == 0 && val == 'L' {
					nextDat[i][j] = '#'
					numChanged++
				} else if totalOccupiedNearby >= 4 && val == '#' {
					nextDat[i][j] = 'L'
					numChanged++
				} else {
					nextDat[i][j] = dat[i][j]
				}
			}
		}
		dat = nextDat
		nextDat = copyDat(dat)
		//printLife(dat)
		iterations++
	}

	occupied := 0
	for _,row := range dat {
		for _,val := range row {
			if val == '#' {
				occupied++
			}
		}
	}

	fmt.Printf("%d seats occupied\n", occupied)
}

func marchUntilSeat(dat [][]rune, startX, startY, xDelta, yDelta int) bool {
	i := startY
	j := startX
	for {
		i += yDelta
		j += xDelta

		if i > 0 && i < len(dat)-1 && j > 0 && j < len(dat[0])-1 {
			if dat[i][j] == '#' {
				return true
			} else if dat[i][j] == 'L' {
				return false
			}
		} else {
			break
		}
	}
	return false
}

func adventDay11B(path string) {
	strs := readLines(path)

	height := len(strs)
	width := len(strs[0])
	dat := make([][]rune, height+2)
	dat[0] = make([]rune, width+2)
	dat[height+1] = make([]rune, width+2)
	for i,str := range strs {
		dat[i+1] = make([]rune, width+2)
		for j := 0; j < len(str); j++ {
			dat[i+1][j+1] = rune(str[j])
		}
	}

	//printLife(dat)

	nextDat := copyDat(dat)

	iterations := 0
	numChanged := 1
	for numChanged > 0 {
		numChanged = 0
		for i,row := range dat {
			if i == 0 || i == (height+1) {
				continue
			}
			for j,val := range row {
				if j == 0 || j == (width+1) {
					continue
				}
				totalOccupiedNearby := 0
				//count
				for x := -1; x < 2; x++ {
					for y := -1; y < 2; y++ {
						if x == 0 && y == 0 {
							continue
						}
						if marchUntilSeat(dat, j, i, y, x) {
							totalOccupiedNearby++
						}
					}
				}

				if totalOccupiedNearby == 0 && val == 'L' {
					nextDat[i][j] = '#'
					numChanged++
				} else if totalOccupiedNearby >= 5 && val == '#' {
					nextDat[i][j] = 'L'
					numChanged++
				} else {
					nextDat[i][j] = dat[i][j]
				}
			}
		}
		dat = nextDat
		nextDat = copyDat(dat)
		//printLife(dat)
		iterations++
	}

	occupied := 0
	for _,row := range dat {
		for _,val := range row {
			if val == '#' {
				occupied++
			}
		}
	}

	fmt.Printf("%d seats occupied\n", occupied)
}

