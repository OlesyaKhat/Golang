package main

import (
	"fmt"
)

type Animal interface {
	Eat() string
	Sound() string
	Move() string
	Age() int
}

type Lion struct {
	age int
}

func (l *Lion) Sound() string {
	return "Лев Рычит"
}

func (l *Lion) Move() string {
	return "Лев шагает"
}

func (l *Lion) Age() int {
	return l.age
}

func(l *Lion) Eat() string {
	return "Других животных"
}

type Giraffe struct {
	age int
}

func (g *Giraffe) Sound() string {
	return "Жираф не знаю"
}

func (g *Giraffe) Move() string {
	return "Жираф шагает"
}

func (g *Giraffe) Age() int {
	return g.age
}

func (g *Giraffe) Eat() string {
	return "Жираф ест листья с дерева"
}

type Snake struct {
	age int
}

func (s *Snake) Sound() string {
	return "Змея Шипит"
}

func (s *Snake) Move() string {
	return "Змея ползет"
}

func (s *Snake) Age() int {
	return s.age
}

func (s *Snake) Eat() string {
	return "Ест мышей"
}

func inputAnimal() Animal {
	var animalType string
	var age int

	fmt.Print("Введите тип животного с большой буквы (Лев, Жираф, Змея): ")
	fmt.Scan(&animalType)
	fmt.Print("Введите возраст животного: ")
	fmt.Scan(&age)

	switch animalType {
	case "Лев":
		return &Lion{age: age}
	case "Жираф":
		return &Giraffe{age: age}
	case "Змея":
		return &Snake{age: age}
	default:
		fmt.Println("Неизвестны тип животного!")
		return nil
	}
}

func main() {
	var animals []Animal
	var count int

	fmt.Print("Сколько животных вы хотите добавить? ")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		fmt.Printf("Введите данные для животного %d:\n", i+1)
		animal := inputAnimal()
		if animal != nil {
			animals = append(animals, animal)
		}
	}

	fmt.Println("\nинформация о животных")
	for _, animal := range animals {
		fmt.Printf("Животное: %T\n", animal)
		fmt.Println("Звук:", animal.Sound())
		fmt.Println("Движение:", animal.Move())
		fmt.Printf("Возраст: %d\n", animal.Age())
		fmt.Println()
	}
}
