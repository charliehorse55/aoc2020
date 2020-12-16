package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2Parse(s string) (min, max int, letter rune, password string) {
	a := strings.SplitN(s, "-", 2)
	min,_ = strconv.Atoi(a[0])
	a = strings.SplitN(a[1], " ", 2)
	max,_ = strconv.Atoi(a[0])
	//fmt.Printf("%s\n", a[1])
	a = strings.SplitN(a[1], ":", 2)
	//fmt.Printf("%s\n", a[0])
	letter = []rune(a[0])[0]
	password = a[1][1:]

	return
}


func adventDay2A(path string) {
	strs := readLines(path)
	numValid := 0
	for _,str := range strs {
		min,max,letter,password := day2Parse(str)
		count := 0
		for _,char := range password {
			if char == letter {
				count++
			}
		}
		//fmt.Printf("count %d\n", count)
		if count >= min && count <= max {
			numValid++
		}

		//fmt.Printf("numValid = %d\n", numValid)

	}
	fmt.Printf("%d valid\n", numValid)
}

func adventDay2B(path string) {
	strs := readLines(path)
	numValid := 0
	for _,str := range strs {
		min,max,letter,password := day2Parse(str)
		//fmt.Printf("count %d\n", count)
		first := rune(password[min-1])
		last := rune(password[max-1])

		if (first == letter) != (last == letter) {
			numValid++
		}

		//fmt.Printf("numValid = %d\n", numValid)

	}
	fmt.Printf("%d valid\n", numValid)
}

