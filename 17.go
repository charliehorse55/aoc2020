package main

import "fmt"

func adventDay17A(path string) {
	strs := readLines(path)

	oldData := make([]bool, 25*25*25)

	for i, str := range strs {
		for j,char := range str {
			if char == '#' {
				oldData[12*(25*25)+((i+8)*25)+j+8] = true
			}
		}
	}

	newData := make([]bool, 25*25*25)

	for r := 0; r < 6; r++ {
		for i := 1; i < 24; i++ {
			for j := 1; j < 24; j++ {
				for k := 1; k < 24; k++ {
					sum := 0
					for x := i - 1; x <= (i + 1); x++ {
						for y := j - 1; y <= (j + 1); y++ {
							for z := k - 1; z <= (k + 1); z++ {
								if x == i && y == j && z == k {
									continue
								}
								if oldData[x*(25*25)+y*25+z] {
									sum++
								}
							}
						}
					}
					newVal := oldData[i*(25*25)+j*25+k]
					currVal := oldData[i*(25*25)+j*25+k]
					if currVal {
						if sum != 2 && sum != 3 {
							newVal = false
						}
					} else {
						if sum == 3 {
							newVal = true
						}
					}
					newData[i*(25*25)+j*25+k] = newVal
				}
			}
		}
		newData, oldData = oldData, newData

	}
	sum := 0
	for i := 1; i < 24; i++ {
		for j := 1; j < 24; j++ {
			for k := 1; k < 24; k++ {
				if oldData[i*(25*25)+j*25+k] {
					sum++
				}
			}
		}
	}
	fmt.Printf("%d\n", sum)

}


func adventDay17B(path string) {
	strs := readLines(path)

	oldData := make([]bool, 25*25*25*25)

	for i, str := range strs {
		for j,char := range str {
			if char == '#' {
				oldData[12*(25*25*25) + 12*(25*25)+((i+8)*25)+j+8] = true
			}
		}
	}

	newData := make([]bool, 25*25*25*25)

	for r := 0; r < 6; r++ {
		for i := 1; i < 24; i++ {
			for j := 1; j < 24; j++ {
				for k := 1; k < 24; k++ {
					for l := 1; l < 24; l++ {
						sum := 0
						for x := i - 1; x <= (i + 1); x++ {
							for y := j - 1; y <= (j + 1); y++ {
								for z := k - 1; z <= (k + 1); z++ {
									for w := l - 1; w <= (l + 1); w++ {
										if x == i && y == j && z == k && w == l {
											continue
										}
										if oldData[w*(25*25*25) + x*(25*25)+y*25+z] {
											sum++
										}
									}
								}
							}
						}
						newVal := oldData[l*(25*25*25) + i*(25*25)+j*25+k]
						currVal := oldData[l*(25*25*25) + i*(25*25)+j*25+k]
						if currVal {
							if sum != 2 && sum != 3 {
								newVal = false
							}
						} else {
							if sum == 3 {
								newVal = true
							}
						}
						newData[l*(25*25*25) + i*(25*25)+j*25+k] = newVal
					}
				}
			}
		}
		newData, oldData = oldData, newData

	}
	sum := 0
	for i := 1; i < 24; i++ {
		for j := 1; j < 24; j++ {
			for k := 1; k < 24; k++ {
				for l := 1; l < 24; l++ {
					if oldData[l*(25*25*25) + i*(25*25)+j*25+k] {
						sum++
					}
				}
			}
		}
	}
	fmt.Printf("%d\n", sum)

}

