package main

import (
	"encoding/json"
	"net/http"
)

// Ручка (обработчик), который возвращает все паспорта
func getPassportsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(passports) // Кодируем массив в JSON и отправляем
}
