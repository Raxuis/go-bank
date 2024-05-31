package main

import "math/rand"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
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
