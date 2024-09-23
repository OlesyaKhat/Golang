package main

import (
	"encoding/base64" // Импортируем пакет для работы с кодированием и декодированием Base64
	"fmt"             // Импортируем пакет для форматированного ввода-вывода

	"github.com/gofiber/fiber/v3/client" // Импортируем клиентскую библиотеку Fiber для выполнения HTTP-запросов
)

func main() {
	// Создаем новый экземпляр клиента Fiber
	cc := client.New()

	// Кодируем строку "john:doe" в формат Base64 для использования в заголовке Authorization
	out := base64.StdEncoding.EncodeToString([]byte("john:doe"))

	// Выполняем GET-запрос к локальному серверу на порту 3000
	resp, err := cc.Get("http://localhost:3000", client.Config{
		Header: map[string]string{ // Указываем заголовки запроса
			"Authorization": "Basic " + out, // Добавляем заголовок авторизации с использованием Basic Auth
		},
	})

	// Проверка на наличие ошибки при выполнении запроса
	if err != nil {
		panic(err) // Если произошла ошибка, выводим её и выходим
	}

	// Выводим тело ответа на консоль
	fmt.Print(string(resp.Body()))
}
