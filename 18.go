package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveSimple(foo string) int {
	f := strings.Fields(foo)
	lVal,err := strconv.Atoi(f[0])
	if err != nil { panic(err)}
	pos := 1
	for pos < len(f) {
		rVal, err := strconv.Atoi(f[pos+1])
		if err != nil { panic(err) }

		switch f[pos] {
		case "+":
			lVal += rVal
		case "*":
			lVal *= rVal
		}
		pos += 2
	}
	return lVal
}

func solveSimple2(foo string) int {
	f := strings.Fields(foo)
	lVal,err := strconv.Atoi(f[0])
	if err != nil { panic(err)}
	pos := 1
	multVal := 1
	for pos < len(f) {
		rVal, err := strconv.Atoi(f[pos+1])
		if err != nil { panic(err) }

		switch f[pos] {
		case "+":
			lVal += rVal
		case "*":
			multVal *= lVal
			lVal = rVal
		}
		pos += 2

	}
	return lVal*multVal
}


func solveEQ(x string, solver func(string) int) int {
	pos := 0
	stack := make([][]byte, 1)
	for pos < len(x) {
		switch x[pos] {
		case '(':
			stack = append(stack, nil)
		case ')':
			eq := string(stack[len(stack)-1])
			//fmt.Printf("eq: %s", eq)
			res := solver(eq)
			//fmt.Printf(" = %d\n", res)
			stack = stack[:len(stack)-1]

			stack[len(stack)-1] = append(stack[len(stack)-1], fmt.Sprintf("%d", res)...)

		//case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
			//val,_ := strconv.Atoi(string(x[pos]))
		default:
			stack[len(stack)-1] = append(stack[len(stack)-1], x[pos])

		}
		pos++
	}
	eq := string(stack[len(stack)-1])
	//fmt.Printf("eq: %s", eq)
	res := solver(eq)
	//fmt.Printf(" = %d\n", res)

	return res
}

func adventDay18A(path string) {
	strs := readLines(path)

	sum := 0
	for _,str := range strs {
		sum += solveEQ(str, solveSimple)
	}
	fmt.Printf("%d\n", sum)
}

func adventDay18B(path string) {
	strs := readLines(path)

	sum := 0
	for _,str := range strs {
		sum += solveEQ(str, solveSimple2)
	}
	fmt.Printf("%d\n", sum)
}
