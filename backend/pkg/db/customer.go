package db

import (
	"fmt"
	"database/sql"

	. "github.com/bebeshen/efrs/pkg/utils"
)

func FindAll(db *sql.DB) []Customer {
    rows, err := db.Query("SELECT * FROM customer")
    // defer rows.Close()
    var u Customer
    var customerList []Customer
    for rows.Next() {
        err = rows.Scan(&u.Id, &u.Username, &u.Password)
        CheckErr(err)
        customerList = append(customerList, u)
        // fmt.Printf("user_id: %d, token: %s, settings: %s\n", u.User_id, u.Token, u.Settings)
    }
    return customerList
}

func FindCustomerById(db *sql.DB, customer_id int) (user *Customer, err error) {

    result := db.QueryRow("SELECT * FROM customer WHERE id=?", customer_id)
    fmt.Println("/db/FindOne [Scan]")
    var u = new(Customer)
    err = result.Scan(&u.Id, &u.Username, &u.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrorUserNotExist
        }
    }
    fmt.Println("/db/FindOne [Finish]")
    return u, err
}

func Insert(db *sql.DB, customer *Customer) string {
    // insert
    stmt, err := db.Prepare("INSERT INTO customer(username, password) values(?,?)")
    CheckErr(err)
    res, err := stmt.Exec(customer.Username, customer.Password)
    CheckErr(err)

    id, err := res.LastInsertId()
    CheckErr(err)

    fmt.Println(id)

    return "success"
}

func Update(db *sql.DB, customer_id int, customer Customer) bool {
    // update
    stmt, err := db.Prepare("update customer set (id, username, password)=(?,?,?) where id=?")
    CheckErr(err)
    res, err := stmt.Exec(customer.Id, customer.Username, customer.Password)
    CheckErr(err)
    
    affect, err := res.RowsAffected()
    CheckErr(err)

    return affect!=0
}