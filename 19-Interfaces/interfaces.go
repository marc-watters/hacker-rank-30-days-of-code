package main

import (
	"fmt"
	"os"
	"reflect"
)

type AdvancedArithmetic interface {
	divisorSum(int) int
}

type Calculator func(int) int

func (c Calculator) divisorSum(n int) int {
	divisors := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}

	var divisorsSummed int
	for i := range divisors {
		divisorsSummed += divisors[i]
	}

	return divisorsSummed
}

func main() {
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		os.Exit(1)
	}

	var c AdvancedArithmetic = new(Calculator)
	t := reflect.TypeOf(c)
	i := reflect.TypeOf((*AdvancedArithmetic)(nil)).Elem()
	if t.Implements(i) {
		fmt.Printf("I implemented: %s\n", i.Name())
		fmt.Println(c.divisorSum(input)) // input: 6, output: 12
	}
}
