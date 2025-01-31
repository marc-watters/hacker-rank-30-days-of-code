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

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)
	/*
	   while(n > 0):
	       remainder = n%2;
	       n = n/2;
	       Insert remainder to front of a list or push onto a stack
	   Print list or stack
	*/
	var binary []int32
	for n > 0 {
		remainder := n % 2
		n = n / 2
		binary = append(binary, remainder)
	}

	var maxOnes int
	var count int
	for i := 0; i < len(binary); i++ {
		if binary[i] == 0 {
			continue
		}

		count = 1
		for j := i + 1; j < len(binary); j++ {
			if binary[j] == 0 {
				break
			}
			count++
		}

		if count > maxOnes {
			maxOnes = count
		}
	}

	fmt.Println(maxOnes)
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
