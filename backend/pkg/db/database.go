package db

import (
	"fmt"
	"database/sql"

    . "github.com/bebeshen/efrs/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() (db *sql.DB) {
    // global db connection object
	DB, err := sql.Open("sqlite3", "../db/sqlite/efrs.db")
	CheckErr(err)

	return DB
}

func CreateTable(db *sql.DB) {
    /* Customer Table */
    fmt.Println("create customer table")
    customer_table := `
    CREATE TABLE IF NOT EXISTS customer(
        c_id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NULL,
        password TEXT NULL,
        c_location TEXT NULL
    );
	`
	_, err := db.Exec(customer_table)
    if (err!=nil) {
        fmt.Println("Fail creating customer table",err)
    }

    /* Employee Table */
    fmt.Println("create employee table")
    employee_table := `
    CREATE TABLE IF NOT EXISTS employee(
        e_id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NULL,
        password TEXT NULL
    );
	`
	_, err = db.Exec(employee_table)
    if (err != nil) {
        fmt.Println("Fail creating employee table", err)
    }

    /* Store Table */
    fmt.Println("create store table")
    store_table := `
    CREATE TABLE IF NOT EXISTS store(
        s_id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NULL,
        location TEXT NULL,
        type TEXT NULL
    );
	`
	_, err = db.Exec(store_table)
    if (err != nil) {
        fmt.Println("Fail creating store table", err)
    }

    /* Food Table */
    fmt.Println("create food table")
    food_table := `
    CREATE TABLE IF NOT EXISTS food(
        f_id INTEGER PRIMARY KEY AUTOINCREMENT,
        category TEXT NULL,
        name TEXT NULL,
        expiredDate TEXT NULL,
        price REAL NULL,
        discount REAL DEFAULT 0.0
    );
	`
	_, err = db.Exec(food_table)
    if (err != nil) {
        fmt.Println("Fail creating food table", err)
    }

    /* Wasted Table */
    fmt.Println("create wasted table")
    wasted_table := `
    CREATE TABLE IF NOT EXISTS wasted(
        f_id INTEGER,
        category TEXT NULL,
        name TEXT NULL
    );
	`
	_, err = db.Exec(wasted_table)
    if (err != nil) {
        fmt.Println("Fail creating wasted table", err)
    }

    /* Order Table */
    fmt.Println("create orders table")
    // use orders to prevent conflict against sql keyword
    orders_table := `
    CREATE TABLE IF NOT EXISTS orders(
        c_id INTEGER NOT NULL,
        s_id INTEGER NOT NULL,
        f_id INTEGER NOT NULL,
        message TEXT NULL,
        status TEXT NULL,
        PRIMARY KEY (c_id, s_id, f_id)
    );
	`
	_, err = db.Exec(orders_table)
    if (err != nil) {
        fmt.Println("Fail creating orders table", err)
    }
}

func QueryBuilder(sql_query string) (*sql.Rows, error) {
    fmt.Println("SQL Query", sql_query)
    rows, err := DB.Query(sql_query)
    if (err!=nil) {
        fmt.Println(err)
        return nil, err
    }

    cols, _ := rows.Columns()
	for i := range cols {
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
    
    return rows, nil
}