package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func updateMask(maskStr string) (orMask, andMask uint) {
	maskBits := strings.Trim(maskStr, " ")
	orMask = uint(0)
	andMask = math.MaxUint64
	for i := 0; i < 36; i++ {
		switch maskBits[i] {
		case 'X':
		case '1':
			orMask |= 1 << (35-i)
		case '0':
			andMask ^= 1 << (35-i)
		}
	}
	return orMask, andMask
}

func adventDay14A(path string) {
	strs := readLines(path)

	orMask, andMask := uint(0), uint(0)

	memory := make(map[uint]uint)
	for _,str := range strs {
		if str[0:4] == "mask" {
			orMask, andMask = updateMask(strings.Split(str, "=")[1])
		} else {
			addr,err := strconv.Atoi(strings.Split(str[4:], "]")[0])
			if err != nil {
				panic(err)
			}
			val,err := strconv.ParseUint(strings.Trim(strings.Split(str, "=")[1], " "), 10, 64)
			if err != nil {
				panic(err)
			}
			memory[uint(addr)] = (uint(val) & andMask) | orMask
		}
	}
	sum := uint(0)
	for _,val := range memory {
		sum += val
	}
	fmt.Printf("%d\n", sum)
}


func addressesForMask(mask string, addr uint) []uint {
	nX := 0
	for i, char := range mask {
		if char == '1' {
			addr |= 1 << (35-i)
		}
		if char == 'X' {
			nX++
		}
	}
	addrs := make([]uint, 1 << nX)
	for i := range addrs {
		addrs[i] = addr
		xj := 0
		for k, char := range mask {
			if char == 'X' {
				toSet := (i >> (nX-(xj+1))) & 0x1
				w := 35-k
				addrs[i] = (addrs[i] & ^uint(uint(1) << uint(w))) | (uint(toSet) << uint(w))
				xj++
			}
		}
	}
	return addrs
}

func adventDay14B(path string) {
	strs := readLines(path)

	currentMask := ""

	memory := make(map[uint]uint)
	for _,str := range strs {
		if str[0:4] == "mask" {
			currentMask = strings.Trim(strings.Split(str, "=")[1], " ")
		} else {
			rawAddr,err := strconv.Atoi(strings.Split(str[4:], "]")[0])
			if err != nil {
				panic(err)
			}
			val,err := strconv.ParseUint(strings.Trim(strings.Split(str, "=")[1], " "), 10, 64)
			if err != nil {
				panic(err)
			}

			addrs := addressesForMask(currentMask, uint(rawAddr))
			for _,addr := range addrs {
				memory[addr] = uint(val)
			}
		}
	}
	sum := uint(0)
	for _,val := range memory {
		sum += val
	}
	fmt.Printf("%d\n", sum)
}

