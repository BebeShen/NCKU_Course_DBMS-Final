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
	{POST} /searchEmployeeById
*/
func SearchEmployeeById(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/SearchEmployeeById")

	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	employee_id, _ := strconv.Atoi(req.FormValue("e_id"))
	employee, _ := FindEmployeeById(DB, employee_id)
	fmt.Println(employee)
	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(employee)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
	fmt.Printf("employee_id = %d\n", employee_id)
}

/*
	{GET} /getAllEmployee
*/
func GetAllEmployee(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/getAllEmployee")

	data := FindAllEmployee(DB)

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(data)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}

/*
	{POST} /addEmployee
*/
func AddEmployee(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/addEmployee")
	
	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)
	
	work_for, _ := strconv.Atoi(req.FormValue("work_for"))
	data := InsertEmployeeDB(DB, req.FormValue("username"), req.FormValue("password"), work_for)

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(data)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}

/*
	{POST} /updateEmployee
*/
func UpdateEmployee(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/updateEmployee")
	
	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	employee_id, _ := strconv.Atoi(req.FormValue("e_id"))
	employee, _ := FindEmployeeById(DB, employee_id)
	employee.Password = req.FormValue("password")

	response := "failure"
	if f := UpdateEmployeeDB(DB, employee) ; f {
		response = "success"
	}

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(response)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}

/*
	{POST} /deleteEmployee
*/
func DeleteEmployee(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/deleteEmployee")
	
	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	employee_id, _ := strconv.Atoi(req.FormValue("e_id"))

	response := "failure"
	if f := DeleteEmployeeDB(DB, employee_id) ; f {
		response = "success"
	}

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(response)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}