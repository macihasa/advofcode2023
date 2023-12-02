package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	day1()

}

func day1() {
	lines := readTxtFile("day1input.txt")

	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var sum int

	for _, line := range lines {

		var firstnum int
		var lastnum int
		fmt.Printf("line: %v\n", line)

		for i, char := range line {
			// Numbers represented as 1 2 3...
			if number, err := strconv.Atoi(string(char)); err == nil {
				if firstnum == 0 {
					firstnum = number
				}

				lastnum = number
				continue
			}

			// Numbers represented as one two ...
			for k, v := range numberMap {

				if len(k) > len(line[i:]) {
					continue
				}

				if k == line[i:i+len(k)] {
					if firstnum == 0 {
						firstnum = v
					}
					lastnum = v
					continue
				}
			}
		}
		sum += firstnum * 10
		sum += lastnum
		fmt.Printf("firstnum: %v\n", firstnum)
		fmt.Printf("lastnum: %v\n", lastnum)
	}

	fmt.Println(sum)
}

func readTxtFile(filename string) []string {
	if filepath.Ext(filename) != "txt" {
		fmt.Println("not a txt file!")
	}
	text, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("unable to read file", err)
	}

	xs := strings.Split(string(text), "\r\n")

	return xs

}
