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

func day7a(entries []string) (bagCount int) {
	r := regexp.MustCompile(`(.+) bags contain (.+)\.`)
	r2 := regexp.MustCompile(`(\d+) (.+?) bag`)

	allBags := map[string]map[string]int64{}

	for _, bagDescription := range entries {
		thisBag := r.FindStringSubmatch(bagDescription)
		bagKey := thisBag[1]
		contains := thisBag[2]
		allBags[bagKey] = map[string]int64{}
		if contains == "no other bags" {
		} else {
			containsBags := r2.FindAllStringSubmatch(contains, -1)
			for _, innerBags := range containsBags {
				innerBagCount, _ := strconv.ParseInt(innerBags[1], 10, 32)
				innerBagKey := innerBags[2]
				allBags[bagKey][innerBagKey] = innerBagCount
			}
		}
	}

	canGoIn := map[string]map[string]bool{}

	for outerBag, innerBagsMap := range allBags {
		for innerBag := range innerBagsMap {
			if canGoIn[innerBag] == nil {
				canGoIn[innerBag] = map[string]bool{}
			}
			canGoIn[innerBag][outerBag] = true
		}
	}

	holdsGold := canGoIn["shiny gold"]

	moreToFind := true

	for moreToFind == true {
		moreToFind = false

		for innerColor := range holdsGold {
			for outerColor := range canGoIn[innerColor] {
				if holdsGold[outerColor] == true {

				} else {
					moreToFind = true
					holdsGold[outerColor] = true
				}
			}
		}
	}

	bagCount = len(holdsGold)
	return
}

func day7b(entries []string) (bagCount int64) {
	r := regexp.MustCompile(`(.+) bags contain (.+)\.`)
	r2 := regexp.MustCompile(`(\d+) (.+?) bag`)

	allBags := map[string]map[string]int64{}

	for _, bagDescription := range entries {
		thisBag := r.FindStringSubmatch(bagDescription)
		bagKey := thisBag[1]
		contains := thisBag[2]
		allBags[bagKey] = map[string]int64{}
		if contains == "no other bags" {
		} else {
			containsBags := r2.FindAllStringSubmatch(contains, -1)
			for _, innerBags := range containsBags {
				innerBagCount, _ := strconv.ParseInt(innerBags[1], 10, 32)
				innerBagKey := innerBags[2]
				allBags[bagKey][innerBagKey] = innerBagCount
			}
		}
	}

	var walkBags func(string) int64

	walkBags = func(outerBag string) (walkedBags int64) {
		for innerBag, count := range allBags[outerBag] {
			walkedBags = walkedBags + count + (count * walkBags(innerBag))
		}
		return
	}

	return walkBags("shiny gold")
}

func main() {
	entries := readFileToSlice("./day7_input.txt")

	fmt.Println(day7a(entries)) //332
	fmt.Println(day7b(entries)) //10875

}
