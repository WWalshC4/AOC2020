package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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

func day6a(entries []string) (count int) {
	groupAnswers := make([]map[rune]bool, 0)

	currentGroup := make(map[rune]bool)
	groupAnswers = append(groupAnswers, currentGroup)

	for _, answer := range entries {
		if answer == "" {
			currentGroup = make(map[rune]bool)
			groupAnswers = append(groupAnswers, currentGroup)
		} else {
			for _, question := range answer {
				currentGroup[question] = true
			}
		}
	}

	for _, groupAnswer := range groupAnswers {
		for _, answer := range groupAnswer {
			if answer == true {
				count++
			}
		}
	}

	return
}

func day6b(entries []string) (count int) {

	currentGroup := make(map[rune]bool)
	for c := 'a'; c <= 'z'; c++ {
		currentGroup[c] = true
	}

	groupAnswers := make([]map[rune]bool, 0)
	groupAnswers = append(groupAnswers, currentGroup)

	for _, answer := range entries {
		if answer == "" {
			currentGroup = make(map[rune]bool)
			for c := 'a'; c <= 'z'; c++ {
				currentGroup[c] = true
			}

			groupAnswers = append(groupAnswers, currentGroup)
			continue
		}
		currentPerson := make(map[rune]bool)
		for _, question := range answer {
			currentPerson[question] = true
		}

		for question, answer := range currentGroup {
			if answer == true {
				if currentPerson[question] != true {
					currentGroup[question] = false
				}
			}
		}

	}

	for _, groupAnswer := range groupAnswers {
		for _, answer := range groupAnswer {
			if answer == true {
				count++
			}
		}
	}

	return
}

func main() {
	entries := readFileToSlice("./day6_input.txt")

	fmt.Println(day6a(entries)) //6714
	fmt.Println(day6b(entries)) //3435

}
