package db

import (
	// "fmt"
	"database/sql"

	. "github.com/bebeshen/efrs/pkg/utils"
)

func FindAllWasted(db *sql.DB) []Wasted {
    rows, err := db.Query("SELECT * FROM Wasted")
    var w Wasted
    var wastedList []Wasted
    for rows.Next() {
        err = rows.Scan(&w.Id, &w.FoodId, &w.Category, &w.Name)
        CheckErr(err)
        wastedList = append(wastedList, w)
    }
    return wastedList
}