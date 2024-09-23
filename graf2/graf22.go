package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func heapify(arr []int, n int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[i] < arr[l] {
		largest = l
	}

	if r < n && arr[largest] < arr[r] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		heapify(arr, n, largest)
	}
}

func heapSort(arr []int, ascending bool) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}

	
	if !ascending {
		reverse(arr)
	}
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите целые числа через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	stringSlice := strings.Split(input, " ")
	var arr []int
	for _, s := range stringSlice {
		number, err := strconv.Atoi(s)
		if err == nil {
			arr = append(arr, number)
		} else {
			fmt.Println("Некорректный ввод. Пожалуйста, введите целые числа.")
			return
		}
	}

	fmt.Print("Выберите порядок сортировки (1 - по возрастанию, 2 - по убыванию): ")
	orderInput, _ := reader.ReadString('\n')
	orderInput = strings.TrimSpace(orderInput)
	var ascending bool

	if orderInput == "1" {
		ascending = true
	} else if orderInput == "2" {
		ascending = false
	} else {
		fmt.Println("Некорректный выбор. Введите 1 или 2.")
		return
	}

	fmt.Println("Неотсортированный массив:", arr)
	heapSort(arr, ascending)
	fmt.Println("Отсортированный массив:", arr)
}
