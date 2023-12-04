package main

import (
	"C"
	"log"
	"os"
	"strconv"
	"strings"
)
import "fmt"

func main() {
	// day1()
	day2part2()

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
		// fmt.Printf("line: %v\n", line)
		// fmt.Printf("firstnum: %v\n", firstnum)
		// fmt.Printf("lastnum: %v\n", lastnum)
	}

	// fmt.Println(sum)
}

func day2() {

	const MAX_NUM_RED = 12
	const MAX_NUM_GREEN = 13
	const MAX_NUM_BLUE = 14

	games := readTxtFile("day2input.txt")

	var idsum int

	for _, game := range games {
		fmt.Printf("game: %v\n", game)
		gameparts := strings.Split(game, ":")
		id, err := strconv.Atoi(strings.Trim(gameparts[0], "Game "))
		if err != nil {
			log.Println("unable to convert game id to integer:", err)
		}
		fmt.Printf("id: %v\n", id)

		gameIsValid := true
		sets := strings.Split(gameparts[1], "; ")
		for _, set := range sets {

			if !gameIsValid {
				break
			}

			// fmt.Printf("set: %v %v\n", i, set)
			pairs := strings.Split(set, ", ")

			for _, pair := range pairs {

				// fmt.Printf("pair: %v\n", pair)

				pair = strings.TrimSpace(pair)
				partsOfPair := strings.Split(pair, " ")
				num, err := strconv.Atoi(partsOfPair[0])
				if err != nil {
					fmt.Println("unable to convert cube count to int", err, partsOfPair)
				}

				if partsOfPair[1] == "red" && num > MAX_NUM_RED {
					gameIsValid = false
					fmt.Printf("Game invalid, %v  %v\n", partsOfPair[1], num)
					break
				}
				if partsOfPair[1] == "green" && num > MAX_NUM_GREEN {
					gameIsValid = false
					fmt.Printf("Game invalid, %v  %v\n", partsOfPair[1], num)
					break
				}
				if partsOfPair[1] == "blue" && num > MAX_NUM_BLUE {
					gameIsValid = false
					fmt.Printf("Game invalid, %v  %v\n", partsOfPair[1], num)
					break
				}

			}
		}

		if gameIsValid {
			idsum += id
		}

	}
	fmt.Printf("idsum: %v\n", idsum)
}

func day2part2() {

	games := readTxtFile("day2input.txt")

	var powerSum int

	for _, game := range games {
		gameparts := strings.Split(game, ":")
		var MAX_RED int
		var MAX_GREEN int
		var MAX_BLUE int

		sets := strings.Split(gameparts[1], "; ")
		for _, set := range sets {

			// fmt.Printf("set: %v %v\n", i, set)
			pairs := strings.Split(set, ", ")

			for _, pair := range pairs {

				// fmt.Printf("pair: %v\n", pair)

				pair = strings.TrimSpace(pair)
				partsOfPair := strings.Split(pair, " ")
				num, err := strconv.Atoi(partsOfPair[0])
				if err != nil {
					fmt.Println("unable to convert cube count to int", err, partsOfPair)
				}

				if partsOfPair[1] == "red" && num > MAX_RED {
					MAX_RED = num
				}
				if partsOfPair[1] == "green" && num > MAX_GREEN {
					MAX_GREEN = num
				}
				if partsOfPair[1] == "blue" && num > MAX_BLUE {
					MAX_BLUE = num
				}

			}
		}
		fmt.Printf("game: %v\npowers: red %v, green %v, blue %v. Multiplied %v\n\n", game, MAX_RED, MAX_GREEN, MAX_BLUE, MAX_RED*MAX_GREEN*MAX_BLUE)

		powerSum += MAX_RED * MAX_GREEN * MAX_BLUE
	}
	fmt.Printf("powerSum: %v\n", powerSum)
}

func readTxtFile(filename string) []string {
	text, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("unable to read file", err)
	}

	xs := strings.Split(string(text), "\r\n")

	return xs

}
