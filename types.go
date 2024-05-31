package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

// We can specify string for the last argument and it will assign it to every argument
func NewAccount(FirstName, LastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: FirstName,
		LastName:  LastName,
		Number:    int64(rand.Intn(1000000)),
	}
}
