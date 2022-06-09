package db

import (
	"fmt"
	"database/sql"

	. "github.com/bebeshen/efrs/pkg/utils"
)

func FindAllFood(db *sql.DB) []Food {
    rows, err := db.Query("SELECT * FROM Food")
    // defer rows.Close()
    var f Food
    var foodList []Food
    for rows.Next() {
        err = rows.Scan(&f.Id, &f.Category, &f.Name, &f.ExpireDate, &f.Price, &f.Discount, &f.Store_at)
        CheckErr(err)
        foodList = append(foodList, f)
    }
    return foodList
}

func FindFoodById(db *sql.DB, f_id int) (food *Food, err error) {

    result := db.QueryRow("SELECT * FROM Food WHERE f_id=?", f_id)
    var f = new(Food)
    err = result.Scan(&f.Id, &f.Category, &f.Name, &f.ExpireDate, &f.Price, &f.Discount, &f.Store_at)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrorUserNotExist
        }
    }
    return f, err
}

func InsertFoodDB(db *sql.DB, category string, name string, expireDate string, price float64, discount float64, store_at int) string {
    fmt.Println(category, name, expireDate, price, discount)
    // insert
    stmt, err := db.Prepare("INSERT INTO Food(category, name, expireDate, price, discount, store_at) values(?,?,?,?,?,?)")
    CheckErr(err)
    res, err := stmt.Exec(category, name, expireDate, price, discount, store_at)
    CheckErr(err)

    id, err := res.LastInsertId()
    CheckErr(err)

    fmt.Println(id)

    return "success"
}

func UpdateFoodDB(db *sql.DB, food *Food) bool {
    stmt, err := db.Prepare("update Food set (category, name, expireDate, price, discount, store_at)=(?,?,?,?,?,?) where f_id=?")
    CheckErr(err)
    res, err := stmt.Exec(food.Category, food.Name, food.ExpireDate, food.Price, food.Discount, food.Store_at, food.Id)
    CheckErr(err)
    
    affect, err := res.RowsAffected()
    CheckErr(err)

    return affect!=0
}

func DeleteFoodDB(db *sql.DB, f_id int) bool {
    // delete
    stmt, err := db.Prepare("DELETE FROM Food WHERE f_id=?")
    CheckErr(err)
    res, err := stmt.Exec(f_id)
    CheckErr(err)
    
    affect, err := res.RowsAffected()
    CheckErr(err)

    return affect!=0
}