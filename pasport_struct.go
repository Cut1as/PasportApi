package main

// Passport - структура для хранения паспортных данных
type Passport struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    string `json:"number"`
}
