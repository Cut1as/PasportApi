package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем ручку (обработчик)
	http.HandleFunc("/passport", passportHandler)

	fmt.Println("Сервер запущен на http://localhost:8080, данные доступны по ссылке http://localhost:8080/passport")
	http.ListenAndServe(":8080", nil)
}
