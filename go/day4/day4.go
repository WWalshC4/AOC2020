package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFileToString(fileName string) (sdata string) {
	data, err := ioutil.ReadFile(fileName)

	check(err)

	sdata = string(data)

	return
}

func readFileToSlice(fileName string) (lines []string) {
	file, err := os.Open(fileName)

	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	check(scanner.Err())

	return
}

func day4a(entries string) (validPassports int64) {
	r := regexp.MustCompile(`(\S+)(\s+)`)
	kv := regexp.MustCompile(`(...):(.+)`)

	passports := make([]map[string]string, 0)

	currentPassport := make(map[string]string)
	passports = append(passports, currentPassport)

	fields := r.FindAllStringSubmatch(entries, -1)

	for _, field := range fields {
		kvField := field[1]
		whitespace := field[2]

		kvPair := kv.FindStringSubmatch(kvField)

		key := kvPair[1]
		value := kvPair[2]
		currentPassport[key] = value

		if whitespace == "\n\n" {
			currentPassport = make(map[string]string)
			passports = append(passports, currentPassport)
		}
	}

	for _, passport := range passports {
		if passport["byr"] != "" &&
			passport["iyr"] != "" &&
			passport["eyr"] != "" &&
			passport["hgt"] != "" &&
			passport["hcl"] != "" &&
			passport["ecl"] != "" &&
			//passport["cid"] != "" &&
			passport["pid"] != "" {
			validPassports++
		}

	}

	return
}

func day4b(entries string) (validPassports int64) {
	r := regexp.MustCompile(`(\S+)(\s+)`)
	kv := regexp.MustCompile(`(...):(.+)`)

	passports := make([]map[string]string, 0)

	currentPassport := make(map[string]string)
	passports = append(passports, currentPassport)

	fields := r.FindAllStringSubmatch(entries, -1)

	for _, field := range fields {
		kvField := field[1]
		whitespace := field[2]

		kvPair := kv.FindStringSubmatch(kvField)

		key := kvPair[1]
		value := kvPair[2]
		currentPassport[key] = value

		if whitespace == "\n\n" {
			currentPassport = make(map[string]string)
			passports = append(passports, currentPassport)
		}
	}

	for _, passport := range passports {
		if validateBYR(passport["byr"]) == false ||
			validateIYR(passport["iyr"]) == false ||
			validateEYR(passport["eyr"]) == false ||
			validateHGT(passport["hgt"]) == false ||
			validateHCL(passport["hcl"]) == false ||
			validateECL(passport["ecl"]) == false ||
			validatePID(passport["pid"]) == false {
			continue
		}
		validPassports++
	}

	return
}

func validateBYR(toValidate string) (valid bool) {
	match, e := regexp.MatchString(`^\d\d\d\d$`, toValidate)
	if e != nil {
		return false
	}
	if match != true {
		return false
	}

	year, e := strconv.ParseInt(toValidate, 10, 32)
	if e != nil {
		return false
	}
	if year < 1920 || year > 2002 {
		return false
	}
	return true
}

func validateIYR(toValidate string) (valid bool) {
	match, e := regexp.MatchString(`^\d\d\d\d$`, toValidate)
	if e != nil {
		return false
	}
	if match != true {
		return false
	}

	year, e := strconv.ParseInt(toValidate, 10, 32)
	if e != nil {
		return false
	}
	if year < 2010 || year > 2020 {
		return false
	}
	return true
}

func validateEYR(toValidate string) (valid bool) {
	match, e := regexp.MatchString(`^\d\d\d\d$`, toValidate)
	if e != nil {
		return false
	}
	if match != true {
		return false
	}

	year, e := strconv.ParseInt(toValidate, 10, 32)
	if e != nil {
		return false
	}
	if year < 2020 || year > 2030 {
		return false
	}
	return true
}

func validateHGT(toValidate string) (valid bool) {
	r := regexp.MustCompile(`^(\d+)(cm|in)$`)

	height := r.FindStringSubmatch(toValidate)

	if height == nil {
		return false
	}

	size, e := strconv.ParseInt(height[1], 10, 32)
	if e != nil {
		return false
	}
	units := height[2]

	if units == "in" {
		if size < 59 || size > 76 {
			return false
		}
	} else if units == "cm" {
		if size < 150 || size > 193 {
			return false
		}
	}
	return true
}

func validateHCL(toValidate string) (valid bool) {
	match, e := regexp.MatchString(`^#[0-9a-f]{6}$`, toValidate)
	if e != nil {
		return false
	}
	if match != true {
		return false
	}
	return true
}

func validateECL(toValidate string) (valid bool) {
	if toValidate == "amb" ||
		toValidate == "blu" ||
		toValidate == "brn" ||
		toValidate == "gry" ||
		toValidate == "grn" ||
		toValidate == "hzl" ||
		toValidate == "oth" {
		return true
	}
	return false
}

func validatePID(toValidate string) (valid bool) {
	match, e := regexp.MatchString(`^\d{9}$`, toValidate)
	if e != nil {
		return false
	}
	if match != true {
		return false
	}
	return true
}

func main() {
	entries := readFileToString("./day4_input.txt")

	fmt.Println(day4a(entries)) //192
	fmt.Println(day4b(entries)) //101

}
