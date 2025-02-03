package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
	idNumber  int
}

func NewPerson(firstName, lastName string, idNumber int) *Person {
	return &Person{firstName, lastName, idNumber}
}

func (p Person) printPerson() {
	fmt.Printf("Name: %s, %s\nID: %d\n", p.lastName, p.firstName, p.idNumber)
}

type Student struct {
	*Person
	scores []int
}

func NewStudent(firstName, lastName string, idNumber int, scores []int) *Student {
	return &Student{
		Person: &Person{firstName, lastName, idNumber},
		scores: scores,
	}
}

func (s *Student) calculate() (grade string) {
	var sumOfScores int
	for i := range s.scores {
		sumOfScores += s.scores[i]
	}

	average := sumOfScores / len(s.scores)
	switch {
	case average >= 90:
		grade = "0"
	case average >= 80:
		grade = "E"
	case average >= 70:
		grade = "A"
	case average >= 55:
		grade = "P"
	case average >= 40:
		grade = "D"
	case average < 40:
		grade = "T"
	}
	return
}

func main() {
	var (
		firstName string
		lastName  string
		idNumber  int
		scores    = make([]int, 0)

		reader = bufio.NewReader(os.Stdin)
	)

	if _, err := fmt.Scan(&firstName, &lastName, &idNumber); err != nil {
		panic(err)
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	vals := strings.Fields(line)
	for i := range vals {
		val, err := strconv.Atoi(vals[i])
		if err != nil {
			continue
		}
		scores = append(scores, val)
	}

	s := NewStudent(firstName, lastName, idNumber, scores)
	s.printPerson()
	fmt.Printf("Grade: %s\n", s.calculate())
}
