package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printArray[T comparable](arr []T) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			break
		}

		inputs, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}

		list := make([]any, 0, inputs)
		for range inputs {
			if ok := scanner.Scan(); ok {
				val := scanner.Text()
				list = append(list, val)
			}
		}

		printArray(list)
	}
}
