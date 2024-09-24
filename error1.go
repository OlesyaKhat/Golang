package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func (l *Lion) Eat() string {
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

type UnknownAnimal struct {
	Name string
}

func (ua *UnknownAnimal) Sound() string {
	return "??? (Error: Unknown animal sound)"
}

func (ua *UnknownAnimal) Move() string {
	return "??? (Error: Unknown animal movement)"
}

func (ua *UnknownAnimal) Age() int {
	return -1
}

func (ua *UnknownAnimal) Eat() string {
	return "??? (Error: Unknown animal diet)"
}

func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v\n", err)
	}
	return strings.TrimSpace(input)
}

func getIntInput(prompt string) (int, error) {
	input := getUserInput(prompt)
	age, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid age input: %w", err)
	}
	return age, nil
}

func createAnimal(animalType string, age int) (Animal, error) {
	switch animalType {
	case "Лев":
		return &Lion{age: age}, nil
	case "Жираф":
		return &Giraffe{age: age}, nil
	case "Змея":
		return &Snake{age: age}, nil
	default:
		return &UnknownAnimal{Name: animalType}, errors.New("unknown animal type")
	}
}

func printAnimalInfo(animal Animal) {
	fmt.Printf("Животное: %T\n", animal)
	fmt.Println("Звук:", animal.Sound())
	fmt.Println("Движение:", animal.Move())
	fmt.Printf("Возраст: %d\n", animal.Age())
	fmt.Println("Питание:", animal.Eat())
	fmt.Println()
}

func main() {

	logFile, err := os.OpenFile("errors.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v\n", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	var animals []Animal
	var count int

	fmt.Print("Сколько животных вы хотите добавить? ")
	count, err = getIntInput("")
	if err != nil {
		log.Printf("Error getting animal count: %v\n", err)
		return
	}

	for i := 0; i < count; i++ {
		fmt.Printf("Введите данные для животного %d:\n", i+1)

		animalType := getUserInput("Введите тип животного с большой буквы (Лев, Жираф, Змея): ")

		age, err := getIntInput("Введите возраст животного: ")
		if err != nil {
			log.Printf("Error getting animal age: %v\n", err)
			continue
		}

		animal, err := createAnimal(animalType, age)
		if err != nil {
			log.Printf("Error creating animal: %v\n", err)
			continue
		}

		animals = append(animals, animal)
	}

	fmt.Println("\nИнформация о животных:")

	for _, animal := range animals {

		if _, ok := animal.(*UnknownAnimal); ok {

			fmt.Println("Предупреждение: Неизвестный тип животного, возможно, некорректные данные.")
		}
		printAnimalInfo(animal)
	}
}
