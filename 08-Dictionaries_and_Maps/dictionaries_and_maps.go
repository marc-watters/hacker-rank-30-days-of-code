package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	phoneBook := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	var numEntries int
	if ok := scanner.Scan(); !ok {
		fmt.Fprintf(os.Stderr, "error scanning entry count: %v", scanner.Err())
	} else {
		if n, err := strconv.Atoi(scanner.Text()); err != nil {
			fmt.Fprintf(os.Stderr, "error parsing entry count: %v", err)
		} else {
			numEntries = n
		}
	}

	for range numEntries {
		if ok := scanner.Scan(); !ok {
			fmt.Fprintf(os.Stderr, "error scanning entry while building phonebook: %v", scanner.Err())
		} else {
			entry := strings.Fields(scanner.Text())
			name := entry[0]
			number := entry[1]
			phoneBook[name] = number
		}
	}

	for range numEntries {
		if ok := scanner.Scan(); !ok {
			fmt.Fprintf(os.Stderr, "error scanning entry during lookup: %v", scanner.Err())
		} else {
			name := scanner.Text()
			if number, ok := phoneBook[name]; ok {
				fmt.Printf("%s=%s\n", name, number)
			} else {
				fmt.Println("Not found")
			}
		}
	}
}
