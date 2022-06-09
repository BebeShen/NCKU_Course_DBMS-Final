package utils

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
	Work_for int
}

type Store struct {
	Id int
	Location string
	Name string
	Type string
	Owner int
}

type Food struct {
	Id int
	Category string
	Name string
	ExpireDate string
	Price float64
	Discount float64
	Store_at int
}

type Wasted struct {
	Id int
	FoodId int
	Category string
	Name string
}


type Order struct {
	CustomerId int
	FoodId int
	StoreId int
	Message string
	Status string
}