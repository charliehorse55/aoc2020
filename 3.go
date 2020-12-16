package main

import "fmt"

func adventDay3A(path string) {
	strs := readLines(path)
	width := len(strs[0])

	isTree := make([]bool, len(strs)*width)
	for i,str := range strs {
		for j,c := range str {
			isTree[i*width + j] = c == '#'
		}
	}
	treesHit := checkSlope(isTree, width, len(strs), 3, 1)
	fmt.Printf("hit %d trees\n", treesHit)
}

func adventDay3B(path string) {
	strs := readLines(path)
	width := len(strs[0])

	isTree := make([]bool, len(strs)*width)
	for i,str := range strs {
		for j,c := range str {
			isTree[i*width + j] = c == '#'
		}
	}
	a := checkSlope(isTree, width, len(strs), 1, 1)
	b := checkSlope(isTree, width, len(strs), 3, 1)
	c := checkSlope(isTree, width, len(strs), 5, 1)
	d := checkSlope(isTree, width, len(strs), 7, 1)
	e := checkSlope(isTree, width, len(strs), 1, 2)
	fmt.Printf("hit %d trees\n", a*b*c*d*e)
}
