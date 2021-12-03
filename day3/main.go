package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRating(in []string, pos int, oxygen bool) []string {
	var out []string

	var seekedChar rune
	count0 := 0
	count1 := 0

	for i := 0; i < len(in); i++ {
		if in[i][pos] == '0' {
			count0++
		} else {
			count1++
		}
	}

	if count1 > count0 {
		if oxygen {
			seekedChar = '1'
		} else {
			seekedChar = '0'
		}
	} else if count0 > count1 {
		if oxygen {
			seekedChar = '0'
		} else {
			seekedChar = '1'
		}
	} else {
		if oxygen {
			seekedChar = '1'
		} else {
			seekedChar = '0'
		}
	}

	for i := 0; i < len(in); i++ {
		if in[i][pos] == byte(seekedChar) {
			out = append(out, in[i])
		}
	}

	if len(out) > 1 {
		return getRating(out, pos+1, oxygen)
	}

	return out
}

func main() {
	buf, err := os.ReadFile("puzzle.dat")
	if err != nil {
		panic(err)
	}

	nums := strings.Split(string(buf), "\n")

	numSize := len(nums[0])

	var gammaRateS string
	var epsilonRateS string

	for i := 0; i < numSize; i++ {
		count0 := 0
		count1 := 0

		for j := 0; j < len(nums); j++ {
			if nums[j][i] == '0' {
				count0++
			} else {
				count1++
			}
		}

		if count1 > count0 {
			gammaRateS += "1"
			epsilonRateS += "0"
		} else {
			gammaRateS += "0"
			epsilonRateS += "1"
		}
	}

	gammaRate, err := strconv.ParseInt(gammaRateS, 2, 64)
	if err != nil {
		panic(err)
	}

	epsilonRate, err := strconv.ParseInt(epsilonRateS, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Gamma rate: %s (%d)\n", gammaRateS, gammaRate)
	fmt.Printf("Epsilon rate: %s (%d)\n", epsilonRateS, epsilonRate)
	fmt.Printf("Power consumption: %d.\n", gammaRate*epsilonRate)

	// Part 2
	fmt.Println("-----------------------------------")
	oxygenRatingS := getRating(nums, 0, true)
	CO2scrubberRatingS := getRating(nums, 0, false)

	oxygenRating, err := strconv.ParseInt(oxygenRatingS[0], 2, 64)
	if err != nil {
		panic(err)
	}

	CO2scrubberRating, err := strconv.ParseInt(CO2scrubberRatingS[0], 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Oxygen generator rating: %s (%d)\n", oxygenRatingS[0], oxygenRating)
	fmt.Printf("CO2	scrubber rating: %s (%d)\n", CO2scrubberRatingS[0], CO2scrubberRating)
	fmt.Printf("Life support rating: %d.\n", oxygenRating*CO2scrubberRating)
}
