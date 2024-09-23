package main

import "fmt"

type Friends struct {
	Name     string
	Surname  string
	Year     int
	Familiar bool
}

func main() {
	friend := Friends{
		Name:     "Ангелина",
		Surname:  "Аведян",
		Year:     19,
		Familiar: true,
	}

	fmt.Println("Имя:", friend.Name)
	fmt.Println("Фамилия:", friend.Surname)
	fmt.Println("Лет:", friend.Year)
	fmt.Println("Знакомы:", friend.Familiar)

}
