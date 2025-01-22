package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		fmt.Fprintf(os.Stdout, "error scanning: %s", scanner.Err())
	}
	inputString := scanner.Text()

	fmt.Println("Hello, World.")
	fmt.Println(inputString)
}
