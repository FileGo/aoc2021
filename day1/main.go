package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	buf, err := os.ReadFile("test.dat")
	if err != nil {
		panic(err)
	}

	input := string(buf)
	var depths []int

	for _, line := range strings.Split(input, "\n") {
		depth, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		depths = append(depths, depth)
	}

	// Part 1
	count := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}

	fmt.Printf("Depth has increased %d times.\n", count)
	fmt.Println("---------------------------------")

	// Part 2
	counter := make(map[int]int)
	//windowSize := 4
	number := 1
	for i := 0; i < len(depths); i++ {
		fmt.Printf("%d  ", depths[i])

		for j := 0; j < i+1; j++ {
			if counter[number+j] >= 3 {
				fmt.Print("  ")
			} else {
				fmt.Printf("%d ", number+j)
				counter[number+j]++
			}

			if j == 3 {
				if i%j == 0 {
					number = number + j + 1
				}
			}
		}

		fmt.Println()
	}
}
