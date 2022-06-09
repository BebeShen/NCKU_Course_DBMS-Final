package route

import (
	"fmt"
    // "strconv"
	"net/http"
    "encoding/json"

    // . "github.com/bebeshen/efrs/pkg/utils"
	. "github.com/bebeshen/efrs/pkg/db"
)

/*
	{GET} /getAllOrder
*/
func GetAllOrder(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/getAllOrder")

	data := FindAllOrder(DB)
	type Row struct{
		Customer_name string `json:"顧客暱稱"`
		Food_name string `json:"食物名稱"`
		Store_name string `json:"商店名稱"`
		Message string `json:"message"`
		Status string `json:"status"`
	}
	var r Row
    var orderList []Row
    for data.Next() {
        if err := data.Scan(&r.Customer_name, &r.Food_name, &r.Store_name, &r.Message, &r.Status) ; err != nil {
			fmt.Println("Error Scan", err)
		}
        orderList = append(orderList, r)
    }

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(orderList)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}