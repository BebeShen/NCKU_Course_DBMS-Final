package db

import (
	// "fmt"
	"database/sql"

	. "github.com/bebeshen/efrs/pkg/utils"
)

func FindAllStore(db *sql.DB) []Store {
    rows, err := db.Query("SELECT * FROM Store")
    var s Store
    var storeList []Store
    for rows.Next() {
        err = rows.Scan(&s.Id, &s.Location, &s.Name, &s.Type, &s.Owner)
        CheckErr(err)
        storeList = append(storeList, s)
    }
    return storeList
}