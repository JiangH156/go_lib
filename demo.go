package main

import "fmt"

type Book struct {
	Name string
}

func main() {
	book := Book{Name: "111111"}
	if book, err := GetBook(); err != nil {
		fmt.Printf("err != nil:  %v", book)
	}
	fmt.Printf("err == nil:  %v", book)
}

func GetBook() (Book, error) {
	return Book{Name: "golang"}, nil
}
