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
	http.HandleFunc("/searchCustomerById", SearchCustomerById)

	DB := Connect()
	CreateTable(DB)

	fmt.Println("Server is running at ", 8789)
	http.ListenAndServe("127.0.0.1:8789", nil)
}