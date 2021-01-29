package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
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

func day8a(entries []string) (accVal int64) {
	instructionR := regexp.MustCompile(`(\w+) ([/+/-]\d+)`)

	alreadyVisited := map[int64]bool{}
	opIndex := int64(0)
	for {
		if alreadyVisited[opIndex] == true {
			break
		}
		alreadyVisited[opIndex] = true
		instruction := instructionR.FindStringSubmatch(entries[opIndex])

		op := instruction[1]
		val, _ := strconv.ParseInt(instruction[2], 10, 64)

		switch op {
		case "nop":
			opIndex++
		case "acc":
			accVal += val
			opIndex++
		case "jmp":
			opIndex += val
		}
	}

	return
}

func day8b(entries []string) (accVal int64) {
	modifiedLine := 0
	previousVal := ""
	for {
		ret, success := runProgram(entries)

		if success == 0 {
			return ret
		}

		if previousVal != "" {
			entries[modifiedLine] = previousVal
		}

		for i := modifiedLine + 1; i <= len(entries); i++ {
			if strings.Contains(entries[i], "nop") {
				previousVal = entries[i]
				entries[i] = strings.Replace(entries[i], "nop", "jmp", 1)
				modifiedLine = i
				break
			}
			if strings.Contains(entries[i], "jmp") {
				previousVal = entries[i]
				entries[i] = strings.Replace(entries[i], "jmp", "nop", 1)
				modifiedLine = i
				break
			}
		}
	}
}

func runProgram(program []string) (accVal int64, returnCode int64) {
	instructionR := regexp.MustCompile(`(\w+) ([/+/-]\d+)`)

	alreadyVisited := map[int64]bool{}
	opIndex := int64(0)

	programLength := int64(len(program))

	for {
		if opIndex == programLength { //ran last instruction
			return accVal, 0
		}
		if alreadyVisited[opIndex] == true { //infinite loop
			return accVal, -1
		}
		if opIndex < 0 { //try to run non existent instruction with negative index
			return accVal, -2
		}
		if opIndex > programLength { //try to run non existent instruction past last instruction
			return accVal, -3
		}
		alreadyVisited[opIndex] = true
		instruction := instructionR.FindStringSubmatch(program[opIndex])

		op := instruction[1]
		val, _ := strconv.ParseInt(instruction[2], 10, 64)

		switch op {
		case "nop":
			opIndex++
		case "acc":
			accVal += val
			opIndex++
		case "jmp":
			opIndex += val
		}
	}

}

func main() {
	entries := readFileToSlice("./day8_input.txt")

	fmt.Println(day8a(entries)) //1317
	fmt.Println(day8b(entries)) // 1033

}
