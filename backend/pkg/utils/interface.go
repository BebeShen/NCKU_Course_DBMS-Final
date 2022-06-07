package utils

import "time"

type Store struct {
	Id int
	Location string
	Name string
}

type Customer struct {
	Id int
	Username string
	Password string
}

type Manager struct {
	Id int
	Username string
	Password string
}

type ExpireFood struct {
	Id int
	Category string
	Name string
	ExpireDate time.Time
	Discount float32
}

type Order struct {
	Id int
	CustomerId int
	FoodId int
	Date time.Time
	Status string
}