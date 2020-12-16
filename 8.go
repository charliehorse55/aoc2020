package main

import (
	"fmt"
	"strconv"
	"strings"
)

func adventDay8A(path string) {
	strs := readLines(path)

	visited := make(map[int]bool)
	pos := 0
	accum := 0
	for {
		_,ok := visited[pos]
		if ok {
			fmt.Printf("loop, value is %d\n", accum)
			break
		}
		a := strings.Fields(strs[pos])
		val,err := strconv.ParseInt(a[1], 10, 64)
		if err != nil {
			panic(err)
		}
		switch a[0]  {
		case "nop":
		case "acc":
			accum += int(val)
		case "jmp":
			pos += int(val)-1
		}
		visited[pos] = true
		pos++
	}



}

func runProgram(ins []string) bool {
	visited := make(map[int]bool)
	pos := 0
	accum := 0
	for {
		if pos >= len(ins) {
			fmt.Printf("value is %d\n", accum)
			return true
		}
		_,ok := visited[pos]
		if ok {
			return false
		}
		a := strings.Fields(ins[pos])
		val,err := strconv.ParseInt(a[1], 10, 64)
		if err != nil {
			panic(err)
		}
		switch a[0]  {
		case "nop":
		case "acc":
			accum += int(val)
		case "jmp":
			pos += int(val)-1
		}
		visited[pos] = true
		pos++
	}
}

func adventDay8B(path string) {
	strs := readLines(path)

	for i,str := range strs {
		orig := str
		a := strings.Fields(str)
		if a[0] == "nop" {
			strs[i] = "jmp" + orig[3:]
			if runProgram(strs) {
				return
			}
			strs[i] = orig
		} else if a[0] == "jmp" {
			strs[i] = "nop" + orig[3:]
			if runProgram(strs) {
				return
			}
			strs[i] = orig
		} else {
			continue
		}
	}

}

