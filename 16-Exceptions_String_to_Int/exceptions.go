package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); ok {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Bad String")
		}
		fmt.Println(val)
	}
}
