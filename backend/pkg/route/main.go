package route

import (
	"fmt"
	"net/http"
    "encoding/json"

    . "github.com/bebeshen/efrs/pkg/utils"
	. "github.com/bebeshen/efrs/pkg/db"
)

func SetupCORS(w *http.ResponseWriter, req *http.Request) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func CheckMethod(w *http.ResponseWriter, method string, allowed string) {
    // check method
	if method != allowed {
		http.Error((*w), "404 not found.", http.StatusNotFound)
		fmt.Fprintf((*w), "Method Not Allowed")
		return
	}
}

func ParseRequestData(w *http.ResponseWriter, req *http.Request)  {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf((*w), "ParseForm() err: %v", err)
		return
	}
}

/*
	{Get} /hello
*/
func Hello(w http.ResponseWriter, req *http.Request) {
	SetupCORS(&w, req)

	ParseRequestData(&w, req)
	InsertCustomerDB(DB, req.FormValue("username"), req.FormValue("password"), req.FormValue("location"))
	customers := FindAllCustomer(DB)
	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(customers)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
	fmt.Println("hello")
}

func Query(w http.ResponseWriter, req *http.Request) {
	SetupCORS(&w, req)

	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)
	
	rows, err := QueryBuilder(req.FormValue("sql_query")); 
	if err != nil {
		return 
	} 

	
	//TODO: switch case upon query table class(for now is customer)
	var c Customer
	var data []Customer
	for rows.Next() {
		if  err := rows.Scan(&c.Id, &c.Username, &c.Password, &c.Location); err != nil {
			fmt.Println(err)
		}
		data = append(data, c)
		fmt.Println(c)
	}


	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(data)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}