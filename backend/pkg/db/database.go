package db

import (
	"fmt"
	"database/sql"

    . "github.com/bebeshen/efrs/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "../db/sqlite/efrs.db")
	CheckErr(err)
    // global db connection object
    DB = db

	return db, err
}

func CreateTable(db *sql.DB) {
	// create table
    sql_table := `
    CREATE TABLE IF NOT EXISTS customer(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NULL,
        password TEXT NULL
    );
	`
	_,err := db.Exec(sql_table)
    if (err!=nil) {
        fmt.Println(err)
    }
    fmt.Println("create table")
}