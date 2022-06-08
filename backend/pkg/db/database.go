package db

import (
	"fmt"
	"database/sql"

    // . "github.com/bebeshen/efrs/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() () {
    // global db connection object
	DB, _ = sql.Open("sqlite3", "../db/sqlite/efrs.db")
}

func CreateTable() {
    /* Customer Table */
    fmt.Println("create Customer table")
    customer_table := `
    CREATE TABLE IF NOT EXISTS Customer(
        c_id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NULL,
        password TEXT NULL,
        c_location TEXT NULL
    );
	`
	_, err := DB.Exec(customer_table)
    if (err!=nil) {
        fmt.Println("Fail creating Customer table",err)
    }

    /* Employee Table */
    fmt.Println("create Employee table")
    employee_table := `
    CREATE TABLE IF NOT EXISTS Employee(
        e_id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NULL,
        password TEXT NULL
    );
	`
	_, err = DB.Exec(employee_table)
    if (err != nil) {
        fmt.Println("Fail creating Employee table", err)
    }

    /* Store Table */
    fmt.Println("create Store table")
    store_table := `
    CREATE TABLE IF NOT EXISTS Store(
        s_id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NULL,
        location TEXT NULL,
        type TEXT NULL
    );
	`
	_, err = DB.Exec(store_table)
    if (err != nil) {
        fmt.Println("Fail creating Store table", err)
    }

    /* Food Table */
    fmt.Println("create Food table")
    food_table := `
    CREATE TABLE IF NOT EXISTS Food(
        f_id INTEGER PRIMARY KEY AUTOINCREMENT,
        category TEXT NULL,
        name TEXT NULL,
        expiredDate TEXT NULL,
        price REAL NULL,
        discount REAL DEFAULT 0.0
    );
	`
	_, err = DB.Exec(food_table)
    if (err != nil) {
        fmt.Println("Fail creating Food table", err)
    }

    /* Wasted Table */
    fmt.Println("create Wasted table")
    wasted_table := `
    CREATE TABLE IF NOT EXISTS Wasted(
        f_id INTEGER,
        category TEXT NULL,
        name TEXT NULL
    );
	`
	_, err = DB.Exec(wasted_table)
    if (err != nil) {
        fmt.Println("Fail creating Wasted table", err)
    }

    /* Order Table */
    fmt.Println("create Orders table")
    // use orders to prevent conflict against sql keyword
    orders_table := `
    CREATE TABLE IF NOT EXISTS Orders(
        c_id INTEGER NOT NULL,
        s_id INTEGER NOT NULL,
        f_id INTEGER NOT NULL,
        message TEXT NULL,
        status TEXT NULL,
        PRIMARY KEY (c_id, s_id, f_id)
    );
	`
	_, err = DB.Exec(orders_table)
    if (err != nil) {
        fmt.Println("Fail creating Orders table", err)
    }
}

func QueryBuilder(sql_query string) (*sql.Rows, error) {
    fmt.Println("SQL Query:", sql_query)
    rows, err := DB.Query(sql_query)
    if (err!=nil) {
        fmt.Println(err)
        return nil, err
    }

    cols, _ := rows.Columns()
    fmt.Println("Columns: ")
	for i := range cols {
		fmt.Print(cols[i])
		fmt.Print("\t")
	}
    fmt.Println()
    
    return rows, nil
}