package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Регистрируем ручку (обработчик)
	http.HandleFunc("/passports", getPassportsHandler)

	fmt.Println("Сервер запущен на http://localhost:8080, данные доступны по ссылке http://localhost:8080/passports")
	http.ListenAndServe(":8080", nil)
}
