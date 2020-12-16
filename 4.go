package main

import (
	"fmt"
	"strconv"
	"strings"
)

func adventDay4A(path string) {
	passports := readPassports(path)

	fmt.Printf("%d passports\n", len(passports))
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	numValid := 0
	for _,passport := range passports {
		left := make(map[string]bool)
		for _,str := range required {
			left[str] = true
		}
		pairs := strings.Fields(passport)
		for _,pair := range pairs {
			vals := strings.Split(pair, ":")
			//fmt.Printf("prop %s\n", vals[0])
			//fmt.Printf("val %s\n", vals[1])
			delete(left, vals[0])
		}
		if len(left) == 0 {
			numValid++
		}
		//panic(nil)
	}
	fmt.Printf("%d valid\n", numValid)
}

func byr(in string) bool {
	val, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return false
	}
	return val >= 1920 && val <= 2002
}

func iyr(in string) bool {
	val, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return false
	}
	return val >= 2010 && val <= 2020
}

func eyr(in string) bool {
	val, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return false
	}
	return val >= 2020 && val <= 2030
}

func hgt(in string) bool {
	if strings.HasSuffix(in, "in") {
		a := strings.TrimSuffix(in, "in")
		val, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return false
		}
		return val >= 59 && val <= 76
	}
	if strings.HasSuffix(in, "cm") {
		a := strings.TrimSuffix(in, "cm")
		val, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			return false
		}
		return val >= 150 && val <= 193
	}
	return false
}
func hcl(in string) bool {
	if in[0] != '#' {
		return false
	}
	if len(in) != 7 {
		return false
	}
	if len(strings.Trim(in[1:], "0123456789abcdef")) != 0 {
		return false
	}
	return true
}

var validECL = map[string]bool {
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func ecl(in string) bool {
	_,ok := validECL[in]
	return ok
}

func pid(in string) bool {
	if len(in) != 9 {
		return false
	}
	_, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return false
	}
	return true
}

func adventDay4B(path string) {
	passports := readPassports(path)

	fmt.Printf("%d passports\n", len(passports))
	required := map[string]func(string)bool{
		"byr":byr,
		"iyr":iyr,
		"eyr":eyr,
		"hgt":hgt,
		"hcl":hcl,
		"ecl":ecl,
		"pid":pid,
	}

	numValid := 0
	for _,passport := range passports {
		left := make(map[string]func(string)bool)
		for key,val := range required {
			left[key] = val
		}
		pairs := strings.Fields(passport)
		for _,pair := range pairs {
			vals := strings.Split(pair, ":")
			//fmt.Printf("prop %s\n", vals[0])
			//fmt.Printf("val %s\n", vals[1])
			f, ok := left[vals[0]]
			if !ok {
				continue
			}
			if f(vals[1]) {
				delete(left, vals[0])
			}
		}
		if len(left) == 0 {
			numValid++
		}
		//panic(nil)
	}
	fmt.Printf("%d valid\n", numValid)
}
