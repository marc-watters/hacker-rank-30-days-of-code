package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Difference struct {
	elements          []int
	maximumDifference int
}

func NewDifference(elements []int) *Difference {
	return &Difference{elements: elements}
}

func (d *Difference) computeDifference() {
	for i := 0; i < len(d.elements); i++ {
		for j := 0; j < len(d.elements)-1; j++ {
			if d.elements[j] > d.elements[j+1] {
				d.elements[j], d.elements[j+1] = d.elements[j+1], d.elements[j]
			}
		}
	}

	min := d.elements[0]
	max := d.elements[len(d.elements)-1]
	d.maximumDifference = max - min
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var inputs int
	if ok := scanner.Scan(); ok {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		inputs = val
	}

	elements := make([]int, 0)
	for range inputs {
		var element int
		_, err := fmt.Scan(&element)
		if err != nil {
			panic(err)
		}
		elements = append(elements, element)
	}

	d := NewDifference(elements)
	d.computeDifference()
	fmt.Println(d.maximumDifference)
}
