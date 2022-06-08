package utils

import "time"

type Customer struct {
	Id int
	Username string
	Password string
	Location string
}

type Employee struct {
	Id int
	Username string
	Password string
}

type Store struct {
	Id int
	Location string
	Name string
	Type string
}

type Food struct {
	Id int
	Category string
	Name string
	ExpireDate time.Time
	Price float32
	Discount float32
}

type Wasted struct {
	Id int
	Category string
	Name string
}


type Order struct {
	Id int
	CustomerId int
	StoreId int
	FoodId int
	// Date time.Time
	Status string
	Message string
}