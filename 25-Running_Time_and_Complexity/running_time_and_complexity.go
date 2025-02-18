package main

import (
	"fmt"
	"math"
)

func main() {
	var inputs int
	if _, err := fmt.Scan(&inputs); err != nil {
		panic(err)
	}

	for range inputs {
		var input int
		if _, err := fmt.Scan(&input); err != nil {
			panic(err)
		}

		if input < 2 {
			fmt.Println("Not prime")
			continue
		}

		prime := true
		sqrt := int(math.Sqrt(float64(input)))
		for i := 2; i <= sqrt; i++ {
			if input%i == 0 {
				prime = false
				break
			}
		}

		if prime {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}

