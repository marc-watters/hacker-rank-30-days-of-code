package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	sumSlice := func(s []int32) int32 {
		var sum int32
		for _, n := range s {
			sum += n
		}
		return sum
	}

	maxSum := int32(-63)
	for i := 0; i < len(arr)-2; i++ {
		var sum int32
		for j := 0; j < len(arr)-2; j++ {
			top := arr[i][j : j+3]
			middle := arr[i+1][j+1]
			bottom := arr[i+2][j : j+3]

			sum = sumSlice(top) + middle + sumSlice(bottom)

			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	fmt.Println(maxSum)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
