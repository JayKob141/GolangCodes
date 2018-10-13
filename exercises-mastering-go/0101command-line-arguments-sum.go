package main

import (
	"fmt"
	"os"
	"strconv"
)

// A program to find the sum of the numeric values from the command line
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		return
	}

	arguments := os.Args
	sum := 0.0

	for i := 1; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		sum += n
	}

	fmt.Println("Summation of numbers:", sum)
}
