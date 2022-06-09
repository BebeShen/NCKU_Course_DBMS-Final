package db

import (
	// "fmt"
	"database/sql"

	// . "github.com/bebeshen/efrs/pkg/utils"
)

func FindAllOrder(db *sql.DB) *sql.Rows {
    rows, _ := db.Query(`
        SELECT c.username AS customer_name, f.name AS food_name, s.name AS store_name, o.message, o.status
        FROM Orders AS o
        LEFT JOIN Customer AS c ON o.c_id = c.c_id
        LEFT JOIN Food AS f on o.f_id = f.f_id
        LEFT JOIN Store AS s on o.s_id = s.s_id`)
    // var o Order
    // var orderList []Order
    // for rows.Next() {
    //     err = rows.Scan(&o.CustomerId, &o.FoodId, &o.StoreId, &o.Message, &o.Status)
    //     CheckErr(err)
    //     orderList = append(orderList, o)
    // }
    return rows
}