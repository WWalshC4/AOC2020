package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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

type req struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func day5a(entries []string) (seatID int) {

	for _, boardingPass := range entries {
		r := bufio.NewReader(strings.NewReader(boardingPass))
		row := 0
		column := 0
		rowBit := 64
		columnBit := 4
		for {
			c, _, err := r.ReadRune()
			if err == io.EOF {
				break
			}
			check(err)
			if c == 'F' {
				rowBit = rowBit / 2
			} else if c == 'B' {
				row = row | rowBit
				rowBit = rowBit / 2
			} else if c == 'L' {
				columnBit = columnBit / 2
			} else if c == 'R' {
				column = column | columnBit
				columnBit = columnBit / 2
			}
		}
		thisSeatID := (row * 8) + column
		if thisSeatID > seatID {
			seatID = thisSeatID
		}
	}

	return
}

func day5b(entries []string) (seatID int) {

	var seats [128][8]bool
	seatIDs := make(map[int]bool)

	for _, boardingPass := range entries {
		r := bufio.NewReader(strings.NewReader(boardingPass))
		row := 0
		column := 0
		rowBit := 64
		columnBit := 4
		for {
			c, _, err := r.ReadRune()
			if err == io.EOF {
				break
			}
			check(err)
			if c == 'F' {
				rowBit = rowBit / 2
			} else if c == 'B' {
				row = row | rowBit
				rowBit = rowBit / 2
			} else if c == 'L' {
				columnBit = columnBit / 2
			} else if c == 'R' {
				column = column | columnBit
				columnBit = columnBit / 2
			}
		}
		seats[row][column] = true
		thisSeatID := (row * 8) + column
		seatIDs[thisSeatID] = true
	}

	for row, columnSeats := range seats {
		for column, occupied := range columnSeats {
			if occupied == false {
				thisSeatID := (row * 8) + column
				if seatIDs[thisSeatID-1] == true && seatIDs[thisSeatID+1] == true {
					return thisSeatID
				}

			}
		}
	}

	return
}

func main() {
	entries := readFileToSlice("./day5_input.txt")

	fmt.Println(day5a(entries)) //919
	fmt.Println(day5b(entries)) //642

}
