package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Enter your code here. Read input from STDIN. Print output to STDOUT
	input := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	for i := 1; i < len(input); i++ {
		var evenChars, oddChars string
		for j := 0; j < len(input[i]); j++ {
			if j%2 == 0 {
				evenChars += string(input[i][j])
			} else {
				oddChars += string(input[i][j])
			}
		}
		fmt.Printf("%s %s\n", evenChars, oddChars)
	}
}
