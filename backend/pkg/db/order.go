package db

import (
	// "fmt"
	"database/sql"

	. "github.com/bebeshen/efrs/pkg/utils"
)

func FindAllOrder(db *sql.DB) []Order {
    rows, err := db.Query("SELECT * FROM Orders")
    var o Order
    var orderList []Order
    for rows.Next() {
        err = rows.Scan(&o.CustomerId, &o.FoodId, &o.StoreId, &o.Message, &o.Status)
        CheckErr(err)
        orderList = append(orderList, o)
    }
    return orderList
}