package route

import (
	"fmt"
    "strconv"
	"net/http"
    "encoding/json"

    // . "github.com/bebeshen/efrs/pkg/utils"
	. "github.com/bebeshen/efrs/pkg/db"
)

/*
	{POST} /searchCustomerById
*/
func SearchCustomerById(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/SearchCustomerById")

	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	customer_id, _ := strconv.Atoi(req.FormValue("cid"))
	customer, _ := FindCustomerById(DB, customer_id)
	fmt.Println(customer)
	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(customer)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
	fmt.Printf("customer_id = %d\n", customer_id)
}

/*
	{GET} /getAllCustomer
*/
func GetAllCustomer(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/getAllCustomer")

	data := FindAllCustomer(DB)

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(data)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}

/*
	{POST} /addCustomer
*/
func AddCustomer(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/addCustomer")
	
	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	c_location := ""
	// 不存在 -> 回傳空字串
	if req.FormValue("c_location") != "" {
		c_location = req.FormValue("c_location")
	}

	data := InsertCustomerDB(DB, req.FormValue("username"), req.FormValue("password"), c_location)

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(data)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}

/*
	{POST} /updateCustomer
*/
func UpdateCustomer(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/updateCustomer")
	
	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	customer_id, _ := strconv.Atoi(req.FormValue("c_id"))
	customer, _ := FindCustomerById(DB, customer_id)
	customer.Location = req.FormValue("c_location")

	response := "failure"
	if f := UpdateCustomerDB(DB, customer) ; f {
		response = "success"
	}

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(response)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}

/*
	{POST} /deleteCustomer
*/
func DeleteCustomer(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/deleteCustomer")
	
	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	customer_id, _ := strconv.Atoi(req.FormValue("c_id"))

	response := "failure"
	if f := DeleteCustomerDB(DB, customer_id) ; f {
		response = "success"
	}

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(response)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}