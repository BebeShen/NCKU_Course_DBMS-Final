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
	{POST} /searchFoodById
*/
func SearchFoodById(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/SearchFoodById")

	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	food_id, _ := strconv.Atoi(req.FormValue("f_id"))
	food, _ := FindFoodById(DB, food_id)
	fmt.Println(food)
	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(food)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
	fmt.Printf("Food_id = %d\n", food_id)
}

/*
	{GET} /getAllFood
*/
func GetAllFood(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/getAllFood")

	data := FindAllFood(DB)

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(data)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}

/*
	{POST} /addFood
*/
func AddFood(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/addFood")
	
	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	price, _ := strconv.ParseFloat(req.FormValue("price"), 64)
	discount, _ := strconv.ParseFloat(req.FormValue("discount"), 64)
	store_at, _ := strconv.Atoi(req.FormValue("store_at"))

	data := InsertFoodDB(DB, req.FormValue("category"), req.FormValue("name"), req.FormValue("expireDate"), price, discount, store_at)

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(data)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}

/*
	{POST} /updateFood

	Description：Only update discount
*/
func UpdateFood(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/updateFood")
	
	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	food_id, _ := strconv.Atoi(req.FormValue("f_id"))
	food, _ := FindFoodById(DB, food_id)
	food.Discount, _ = strconv.ParseFloat(req.FormValue("discount"), 64)

	response := "failure"
	if f := UpdateFoodDB(DB, food) ; f {
		response = "success"
	}

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(response)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}

/*
	{POST} /deleteFood
*/
func DeleteFood(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/deleteFood")
	
	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)

	food_id, _ := strconv.Atoi(req.FormValue("f_id"))

	response := "failure"
	if f := DeleteFoodDB(DB, food_id) ; f {
		response = "success"
	}

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(response)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}