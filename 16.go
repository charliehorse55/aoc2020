package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start,end int
}

func adventDay16A(path string) {
	strs := readLines(path)

	rules := make(map[string][2]Range)
	pos := 0
	for {
		if strings.HasPrefix(strs[pos],"your ticket:") {
			pos++
			break
		}
		components := strings.Split(strs[pos], ":")
		name := components[0]
		ranges := strings.Fields(components[1])
		ranges[1] = ranges[2]
		var ruleRanges [2]Range
		for i := 0; i < 2; i++ {
			x := strings.Split(ranges[i], "-")
			start,err := strconv.Atoi(x[0])
			if err != nil { panic(err)}
			end, err := strconv.Atoi(x[1])
			if err != nil { panic(err)}
			ruleRanges[i] = Range{start: start,end:end}
		}
		rules[name] = ruleRanges
		pos++
	}

	//fmt.Printf("%v\n", rules)

	sum := 0

	pos += 2
	//fmt.Printf("strs[pos-1]:%s\n", strs[pos-1])
	for pos < len(strs) {
		vals := strings.Split(strs[pos], ",")
		ints := toInts(vals)

	out:
		for _,val := range ints {
			for _,ranges := range rules {
				for i := 0; i < 2; i++ {
					if val <= ranges[i].end && val >= ranges[i].start {
						continue out
					}
				}
			}
			//fmt.Printf("invalid number %d\n", val)
			sum += val
		}
		pos++
	}
	fmt.Printf("%d\n", sum)
}

type Poskeys struct {
	index int
	keys []string
}

type PosKeyArray []Poskeys

func (p PosKeyArray)Len() int {
	return len(p)
}

func (p PosKeyArray)Less(i,j int) bool {
	return len(p[i].keys) < len(p[j].keys)
}

func (p PosKeyArray)Swap(i,j int) {
	p[i], p[j] = p[j], p[i]
}

func adventDay16B(path string) {
	strs := readLines(path)

	rules := make(map[string][2]Range)
	pos := 0
	for {
		if strings.HasPrefix(strs[pos],"your ticket:") {
			pos++
			break
		}
		components := strings.Split(strs[pos], ":")
		name := components[0]
		ranges := strings.Fields(components[1])
		ranges[1] = ranges[2]
		var ruleRanges [2]Range
		for i := 0; i < 2; i++ {
			x := strings.Split(ranges[i], "-")
			start,err := strconv.Atoi(x[0])
			if err != nil { panic(err)}
			end, err := strconv.Atoi(x[1])
			if err != nil { panic(err)}
			ruleRanges[i] = Range{start: start,end:end}
		}
		rules[name] = ruleRanges
		pos++
	}

	//fmt.Printf("%v\n", rules)
	vals := strings.Split(strs[pos], ",")
	myTicket := toInts(vals)

	fmt.Printf("myTicket: %v\n", myTicket)

	pos += 2

	var otherTickets [][]int

	//fmt.Printf("strs[pos-1]:%s\n", strs[pos-1])
	for pos < len(strs) {
		vals := strings.Split(strs[pos], ",")
		ints := toInts(vals)

		valid := true
	out:
		for _,val := range ints {
			for _,ranges := range rules {
				for i := 0; i < 2; i++ {
					if val <= ranges[i].end && val >= ranges[i].start {
						continue out
					}
				}
			}
			//fmt.Printf("invalid number %d\n", val)
			valid = false
		}
		if valid {
			otherTickets = append(otherTickets, ints)
		}
		pos++
	}

	nfields := len(otherTickets[0])
	validPossible := make([]Poskeys, nfields)
	for i := 0; i < nfields; i++ {
		possible := make(map[string]bool)
		for key := range rules {
			possible[key] = true
		}
		for _,ticket := range otherTickets {
			val := ticket[i]
		ruleLoop:
			for key := range possible {
				ranges := rules[key]
				found := false
				for j := 0; j < 2; j++ {
					if val <= ranges[j].end && val >= ranges[j].start {
						found = true
					}
				}
				if !found {
					delete(possible, key)
					continue ruleLoop
				}
			}
		}
		posStrs := make([]string, len(possible))
		k := 0
		for key := range possible {
			posStrs[k] = key
			k++
		}
		validPossible[i] = Poskeys{
			index: i,
			keys:  posStrs,
		}
	}
	sort.Sort(PosKeyArray(validPossible))

	available := make(map[string]bool)
	for key := range rules {
		available[key] = true
	}

	mult := 1
	for _,posKey := range validPossible {
		for _,key := range posKey.keys {
			_,ok := available[key]
			if ok {
				fmt.Printf("%d -> %s\n", posKey.index, key)
				if strings.HasPrefix(key, "departure") {
					mult *= myTicket[posKey.index]
				}
				delete(available, key)
				break
			}
		}
	}
	fmt.Printf("depature multiplied: %d\n", mult)
}

