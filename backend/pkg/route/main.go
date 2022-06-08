package route

import (
	"fmt"
	"net/http"
    "encoding/json"

    . "github.com/bebeshen/efrs/pkg/utils"
	. "github.com/bebeshen/efrs/pkg/db"
)

/*
	{Get} /hello
*/
func Hello(w http.ResponseWriter, req *http.Request) {
	SetupCORS(&w, req)

	c := new(Customer)
	c.Username = "test"
	c.Password = "test"
	InsertCustomer(DB, c)
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
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	rows, err := QueryBuilder(req.FormValue("sql_query")); 
	if err != nil {
		return 
	} 

	
	//TODO: switch case upon query table class(for now is customer)
	var c Customer
	var data []Customer
	for rows.Next() {
		if  err := rows.Scan(&c.Id, &c.Username, &c.Password); err != nil {
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