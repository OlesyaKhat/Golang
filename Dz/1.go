package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	book := Book{
		Title:  "Грокаем Алгоритмы",
		Author: "Адитья Бхаргава",
		Year:   2017,
	}
	fmt.Println("Название:", book.Title)
	fmt.Println("Автор:", book.Author)
	fmt.Println("Год издания:", book.Year)
}
