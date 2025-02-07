package main

import (
	"fmt"
	"math"
)

type Calculator struct{}

func (c Calculator) power(n, p float64) (float64, error) {
	if n < 0 || p < 0 {
		return 0, fmt.Errorf("n and p should be non-negative")
	}
	return math.Pow(n, p), nil
}

func main() {
	var inputs int
	_, err := fmt.Scan(&inputs)
	if err != nil {
		panic(err)
	}

	C := new(Calculator)
	for range inputs {
		var n, p float64
		_, err := fmt.Scan(&n, &p)
		if err != nil {
			panic(err)
		}

		if answer, err := C.power(n, p); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(answer)
		}
	}
}
