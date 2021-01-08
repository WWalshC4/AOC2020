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

func day3a(entries []string) (treesHit int64) {
	xPos := 0
	for _, entry := range entries {
		isTree := (entry[xPos:xPos+1] == "#")
		if isTree {
			treesHit++
		}
		xPos = xPos + 3
		xPos = xPos % len(entry)
	}
	return
}

func day3b(entries []string) (validCount int64) {
	var xPos [5]int
	var treesHit [5]int64

	xDelta := [...]int{1, 3, 5, 7, 1}
	yDelta := [...]int{1, 1, 1, 1, 2}

	for y, entry := range entries {
		for i := 0; i < len(xPos); i++ {
			if y%yDelta[i] == 0 {
				isTree := (entry[xPos[i]:xPos[i]+1] == "#")
				if isTree {
					treesHit[i]++
				}
				xPos[i] = xPos[i] + xDelta[i]
				xPos[i] = xPos[i] % len(entry)
			}
		}
	}
	var totalTreesHit int64 = 1
	for _, x := range treesHit {
		fmt.Println(x)
		totalTreesHit = totalTreesHit * x
	}
	return totalTreesHit
}

func main() {
	entries := readFileToSlice("./day3_input.txt")

	fmt.Println(day3a(entries)) //282
	fmt.Println(day3b(entries)) //958815792

}
