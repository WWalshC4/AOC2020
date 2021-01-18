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
	return
}

func day6b(entries []string) (count int) {
	return
}

func main() {
	entries := readFileToSlice("./day5_input.txt")

	fmt.Println(day6a(entries)) //919
	fmt.Println(day6b(entries)) //642

}
