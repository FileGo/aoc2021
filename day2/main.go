package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	buf, err := os.ReadFile("puzzle.dat")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buf), "\n")

	depth := 0
	position := 0

	for _, line := range lines {
		cols := strings.Split(line, " ")

		num, err := strconv.Atoi(cols[1])
		if err != nil {
			panic(err)
		}

		switch cols[0] {
		case "forward":
			position += num
		case "down":
			depth += num
		case "up":
			depth -= num
		}
	}

	fmt.Printf("Final position: %d forward, %d deep.\n", position, depth)
	fmt.Printf("Result: %d\n", position*depth)
	fmt.Println("--------------------------------------")

	// Part 2
	depth = 0
	position = 0
	aim := 0

	for _, line := range lines {
		cols := strings.Split(line, " ")

		num, err := strconv.Atoi(cols[1])
		if err != nil {
			panic(err)
		}

		switch cols[0] {
		case "forward":
			position += num
			depth += aim * num
		case "down":
			aim += num
		case "up":
			aim -= num
		}
	}

	fmt.Printf("Final position: %d forward, %d deep.\n", position, depth)
	fmt.Printf("Result: %d\n", position*depth)

}
