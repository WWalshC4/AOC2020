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

func day2a(entries []string) (validCount int64) {
	r, err := regexp.Compile(`(\d+)\-(\d+) ([A-Za-z])\: ([A-Za-z]+)`)

	check(err)

	for _, entry := range entries {
		match := r.FindStringSubmatch(entry)
		if len(match) > 0 {
			min := match[1]
			max := match[2]
			letter := match[3]
			password := match[4]
			if min != "" && max != "" && letter != "" && password != "" {
				min, _ := strconv.ParseInt(min, 10, 32)
				max, _ := strconv.ParseInt(max, 10, 32)
				var count int64 = 0
				for _, char := range password {
					if string(char) == letter {
						count++
					}
				}
				if count >= min && count <= max {
					validCount = validCount + 1
				}
			}
		}
	}
	return
}

func day2b(entries []string) (validCount int64) {
	r, err := regexp.Compile(`(\d+)\-(\d+) ([A-Za-z])\: ([A-Za-z]+)`)

	check(err)

	for _, entry := range entries {
		match := r.FindStringSubmatch(entry)
		if len(match) > 0 {
			a := match[1]
			b := match[2]
			letter := match[3]
			password := match[4]
			if a != "" && b != "" && letter != "" && password != "" {
				a, _ := strconv.ParseInt(a, 10, 32)
				b, _ := strconv.ParseInt(b, 10, 32)

				aMatch := password[a-1:a] == letter
				bMatch := password[b-1:b] == letter

				if aMatch != bMatch {
					validCount++
				}

			}
		}
	}
	return
}

func main() {
	entries := readFileToSlice("./day2_input.txt")

	fmt.Println(day2a(entries)) //393
	fmt.Println(day2b(entries)) //690

}
