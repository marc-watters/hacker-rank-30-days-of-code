package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type ABCMeta interface{ display() }

type Book struct {
	ABCMeta

	title  string
	author string
}

type MyBook struct {
	*Book

	price float64
}

func NewMyBook(title, author string, price float64) *MyBook {
	myBook := &MyBook{
		Book:  &Book{title: title, author: author},
		price: price,
	}
	myBook.ABCMeta = myBook
	return myBook
}

func (b *MyBook) display() {
	fmt.Printf("Title: %s\nAuthor: %s\nPrice: %.2f", b.title, b.author, b.price)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var title string
	if ok := scanner.Scan(); ok {
		title = scanner.Text()
	}

	var author string
	if ok := scanner.Scan(); ok {
		author = scanner.Text()
	}

	var price float64
	if ok := scanner.Scan(); ok {
		val, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
		price = val
	}

	newNovel := ABCMeta(NewMyBook(title, author, price))
	newNovel.display()
}
