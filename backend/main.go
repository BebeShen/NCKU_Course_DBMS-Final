package main

import (
	"fmt"
	"net/http"

    . "github.com/bebeshen/efrs/pkg/route"
	. "github.com/bebeshen/efrs/pkg/db"
)

func main() {

	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/queryBuilder", Query)
	// customer
	http.HandleFunc("/addCustomer", AddCustomer)
	http.HandleFunc("/updateCustomer", UpdateCustomer)
	http.HandleFunc("/deleteCustomer", DeleteCustomer)
	http.HandleFunc("/getAllCustomer", GetAllCustomer)
	http.HandleFunc("/searchCustomerById", SearchCustomerById)
	// employee
	http.HandleFunc("/addEmployee", AddEmployee)
	http.HandleFunc("/updateEmployee", UpdateEmployee)
	http.HandleFunc("/deleteEmployee", DeleteEmployee)
	http.HandleFunc("/getAllEmployee", GetAllEmployee)
	http.HandleFunc("/searchEmployeeById", SearchEmployeeById)
	// food
	http.HandleFunc("/addFood", AddFood)
	http.HandleFunc("/updateFood", UpdateFood)
	http.HandleFunc("/deleteFood", DeleteFood)
	http.HandleFunc("/getAllFood", GetAllFood)
	http.HandleFunc("/searchFoodById", SearchFoodById)
	// store
	http.HandleFunc("/getAllStore", GetAllStore)
	// order
	http.HandleFunc("/getAllOrder", GetAllOrder)
	// wasted
	http.HandleFunc("/getAllWasted", GetAllWasted)

	Connect()
	CreateTable()
	InsertDefaultData()

	fmt.Println("Server is running at ", 8789)
	http.ListenAndServe("127.0.0.1:8789", nil)
}