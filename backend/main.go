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

	Connect()
	CreateTable()

	fmt.Println("Server is running at ", 8789)
	http.ListenAndServe("127.0.0.1:8789", nil)
}