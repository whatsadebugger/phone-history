package main

import "github.com/asdine/storm"

type address struct {
	ID        int    `json:"id" storm:"id,increment"`
	FirstName string `json:"firstname" storm:"index"`
	LastName  string `json:"lastname"  storm:"index"`
	Email     string `json:"email"     storm:"index,unique"`
	Phone     string `json:"phone"     storm:"index,unique"`
}

func createDatabase() *storm.DB {
	db, err := storm.Open("address.db")
	PanicIfError(err)
	return db
}
