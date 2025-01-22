package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	_ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var ii uint64
	var dd float64
	var ss string

	// Read and save an integer, double, and String to your variables.
	var input string
	if ok := scanner.Scan(); !ok {
		fmt.Fprintf(os.Stderr, "error scanning: %v", scanner.Err())
		os.Exit(1)
	}

	input = scanner.Text()
	if val, err := strconv.Atoi(input); err != nil {
		fmt.Fprintf(os.Stderr, "invalid input: %v", scanner.Err())
	} else {
		ii = uint64(val)
	}

	if ok := scanner.Scan(); !ok {
		fmt.Fprintf(os.Stderr, "error scanning: %v", scanner.Err())
		os.Exit(1)
	}
	input = scanner.Text()

	if val, err := strconv.Atoi(input); err != nil {
		fmt.Fprintf(os.Stderr, "invalid input: %v", scanner.Err())
	} else {
		dd = float64(val)
	}

	if ok := scanner.Scan(); !ok {
		fmt.Fprintf(os.Stderr, "error scanning: %v", scanner.Err())
		os.Exit(1)
	}
	ss = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + ii)

	// Print the sum of the double variables on a new line.
	fmt.Println(d + dd)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s, ss)
}
