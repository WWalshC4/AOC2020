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

func day1a(entries []string) (product int64) {
	for i, v := range entries {
		for _, v2 := range entries[i+1:] {
			int1, _ := strconv.ParseInt(v, 10, 32)
			int2, _ := strconv.ParseInt(v2, 10, 32)
			if int1+int2 == 2020 {
				return int1 * int2
			}
		}
	}
	return
}

func day1b(entries []string) (product int64) {
	for i, v := range entries {
		for i2, v2 := range entries[i+1:] {
			for _, v3 := range entries[i2+1:] {
				int1, _ := strconv.ParseInt(v, 10, 32)
				int2, _ := strconv.ParseInt(v2, 10, 32)
				int3, _ := strconv.ParseInt(v3, 10, 32)
				if int1+int2+int3 == 2020 {
					return int1 * int2 * int3
				}
			}
		}
	}
	return
}

func main() {
	entries := readFileToSlice("./day1_input.txt")

	fmt.Println(day1a(entries))

	fmt.Println(day1b(entries))

}
