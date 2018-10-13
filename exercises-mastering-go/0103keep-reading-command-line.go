package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// A program that keeps reading integers until the word "stop" is given
func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	condition := false
	for !condition && scanner.Scan() {
		text := scanner.Text()

		if text == "stop" {
			condition = true
		} else {
			n, err := strconv.ParseInt(text, 10, 32)
			if err != nil {
				fmt.Println("Error reading integer >", scanner.Text())
			} else {
				fmt.Println("Integer readed >", n)
			}
		}

	}
}
