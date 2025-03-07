package main

import (
	"encoding/json"
	"net/http"
)

// passportHandler – ручка (обработчик) для получения паспортных данных
func passportHandler(w http.ResponseWriter, r *http.Request) {
	passport := Passport{
		FirstName: "Михаил",
		LastName:  "Круглый",
		Number:    "1234 567890",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(passport)
}
