package controllers

import "github.com/lucianorbr/formWeb_Go/db"

type SubscriptionController struct {
	Name  string
	Email string
}

func Create(name string, email string) error {
	s := SubscriptionController{Name: name, Email: email}

	return db.Insert("newsletter", s)
}
