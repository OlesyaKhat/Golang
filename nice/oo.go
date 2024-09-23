package main

import "fmt"

func main() {
	fmt.Println("Хочешь лучше узнать Олесю? (да/нет)")

	var ansver string
	fmt.Scanln(&ansver)

	if ansver == "да" {
		fmt.Println("Супер! Это прекрасная девушка , программист !")
	} else if ansver == "нет" {
		fmt.Println("Жаль, но ты все равно узнаешь , что она программист!!!!!!)) P.s. когда она взломает тебя")
	} else {
		fmt.Println("Пожалуйста ответь 'да' или 'нет'.")
	}
}
