package main

import (
	"fmt"
	"strings"
)

func adventDay6A(path string) {
	strs := readPassports(path)

	count := 0
	for _,str := range strs {
		questions := make(map[rune]bool)
		stuff := strings.Fields(str)
		for _,str2 := range stuff {
			for _,char := range str2 {
				questions[char] = true
			}
		}
		count += len(questions)
	}
	fmt.Printf("count %d\n", count)
}

func adventDay6B(path string) {
	strs := readPassports(path)

	count := 0
	for _,str := range strs {
		var questions map[rune]bool
		done1 := false
		stuff := strings.Fields(str)
		for _,str2 := range stuff {
			questions2 := make(map[rune]bool)
			for _,char := range str2 {
				questions2[char] = true
			}
			if !done1 {
				questions = questions2
				done1 = true
			} else {
				for key := range questions {
					ok, _ := questions2[key]
					if !ok {
						delete(questions, key)
					}
				}
			}
		}
		count += len(questions)
	}
	fmt.Printf("count %d\n", count)
}

