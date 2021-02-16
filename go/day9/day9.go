package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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

func day9a(entries []string) (index int64) {
	ints := []int64{}

	for _, i := range entries {
		val, err := strconv.ParseInt(i, 10, 0)
		check(err)
		ints = append(ints, val)
	}

	for nextNum := 25; nextNum < len(ints); nextNum++ {
		checkValid := ints[nextNum]
		isValid := false
		for start, i1 := range ints[nextNum-25 : nextNum] {
			for _, i2 := range ints[start+1 : nextNum] {
				if i1+i2 == checkValid {
					isValid = true
					break
				}
			}
			if isValid == true {
				break
			}
		}
		if isValid != true {
			return checkValid
		}
	}

	return
}

func day9b(entries []string) (index int64) {
	ints := []int64{}

	for _, i := range entries {
		val, err := strconv.ParseInt(i, 10, 0)
		check(err)
		ints = append(ints, val)
	}

	target := int64(3199139634)

	for i := 0; i < len(ints); i++ {
		smallest := target
		largest := int64(0)

		total := int64(0)
		current := i

		for total < target {
			thisInt := ints[current]
			total += thisInt
			if thisInt < smallest {
				smallest = thisInt
			}
			if thisInt > largest {
				largest = thisInt
			}
			current++
		}
		if total == target {
			return (smallest + largest)
		}
	}
	return
}

func main() {
	entries := readFileToSlice("./day9_input.txt")

	fmt.Println(day9a(entries)) //3199139634
	fmt.Println(day9b(entries)) //438559930
}
