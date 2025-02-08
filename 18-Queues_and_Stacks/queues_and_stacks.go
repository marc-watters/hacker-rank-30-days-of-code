package main

import (
	"bufio"
	"fmt"
	"os"
)

type Solution struct {
	stack []string
	queue []string
}

func (s *Solution) pushCharacter(c string) {
	s.stack = append(s.stack, c)
}

func (s *Solution) popCharacter() string {
	var pop string
	pop, s.stack = s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
	return pop
}

func (s *Solution) enqueueCharacter(c string) {
	s.queue = append(s.queue, c)
}

func (s *Solution) dequeueCharacter() string {
	var pop string
	pop, s.queue = s.queue[0], s.queue[1:]
	return pop
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	if ok := scanner.Scan(); ok {
		input = scanner.Text()
	}

	S := new(Solution)
	for i := range input {
		c := string(input[i])
		S.pushCharacter(c)
		S.enqueueCharacter(c)
	}

	isPalindrome := true
	/***
	* pop the top character from the stack
	* dequeu the first character from the queue
	* compare both the characters
	***/
	for range len(input) / 2 {
		if S.popCharacter() != S.dequeueCharacter() {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Printf("The word, %s, is a palindrome.\n", input)
	} else {
		fmt.Printf("The word, %s, is not a palindrome.\n", input)
	}
}
