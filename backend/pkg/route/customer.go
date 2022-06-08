package route

import (
	"fmt"
    "strconv"
	"net/http"
    "encoding/json"

    . "github.com/bebeshen/efrs/pkg/utils"
	. "github.com/bebeshen/efrs/pkg/db"
)

/*
	{Post} /searchCustomerById
*/
func SearchCustomerById(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/searchCustomerById")

	// check method
	if req.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// parse data
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	customer_id, _ := strconv.Atoi(req.FormValue("cid"))
	customer, _ := FindCustomerById(DB, customer_id)
	fmt.Println(customer)
	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(customer)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
	fmt.Printf("customer_id = %d\n", customer_id)
}