package main

import (
	"fmt"
	"os"
	"strconv"
)

// A program to find the average of the float values from the command line
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		return
	}

	arguments := os.Args
	sum := 0.0
	N := len(arguments) - 1

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			// discard the value
			N = N - 1
		} else {
			sum += n
		}
	}

	if N > 0 {
		average := sum / float64(N)
		fmt.Println("Average of numbers:", average)
	} else {
		fmt.Println("Please give one or more floats.")
	}
}
