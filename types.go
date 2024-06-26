package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccount int `json:"to_account"`
	Amount    int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

// We can specify string for the last argument and it will assign it to every argument
func NewAccount(FirstName, LastName string) *Account {
	return &Account{
		FirstName: FirstName,
		LastName:  LastName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}
