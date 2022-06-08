package db

import (
	"fmt"
	"database/sql"

	. "github.com/bebeshen/efrs/pkg/utils"
)

func FindAllCustomer(db *sql.DB) []Customer {
    rows, err := db.Query("SELECT * FROM Customer")
    // defer rows.Close()
    var u Customer
    var customerList []Customer
    for rows.Next() {
        err = rows.Scan(&u.Id, &u.Username, &u.Password, &u.Location)
        CheckErr(err)
        customerList = append(customerList, u)
        // fmt.Printf("user_id: %d, token: %s, settings: %s\n", u.User_id, u.Token, u.Settings)
    }
    return customerList
}

func FindCustomerById(db *sql.DB, c_id int) (user *Customer, err error) {

    result := db.QueryRow("SELECT * FROM Customer WHERE c_id=?", c_id)
    var u = new(Customer)
    err = result.Scan(&u.Id, &u.Username, &u.Password, &u.Location)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrorUserNotExist
        }
    }
    return u, err
}

func InsertCustomerDB(db *sql.DB, username string, password string, location string) string {
    fmt.Println(username, password, location)
    // insert
    stmt, err := db.Prepare("INSERT INTO Customer(username, password, c_location) values(?,?,?)")
    CheckErr(err)
    res, err := stmt.Exec(username, password, location)
    CheckErr(err)

    id, err := res.LastInsertId()
    CheckErr(err)

    fmt.Println(id)

    return "success"
}

func UpdateCustomerDB(db *sql.DB, customer *Customer) bool {
    // update
    stmt, err := db.Prepare("update Customer set (username, password, c_location)=(?,?,?) where c_id=?")
    CheckErr(err)
    res, err := stmt.Exec(customer.Username, customer.Password, customer.Location, customer.Id)
    CheckErr(err)
    
    affect, err := res.RowsAffected()
    CheckErr(err)

    return affect!=0
}

func DeleteCustomerDB(db *sql.DB, c_id int) bool {
    // delete
    stmt, err := db.Prepare("DELETE FROM Customer WHERE c_id=?")
    CheckErr(err)
    res, err := stmt.Exec(c_id)
    CheckErr(err)
    
    affect, err := res.RowsAffected()
    CheckErr(err)

    return affect!=0
}