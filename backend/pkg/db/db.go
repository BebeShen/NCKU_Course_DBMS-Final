package db

import (
	"database/sql"
    "fmt"
    "time"
    "errors"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	User_id uint32
	Token string
	Device string
	Settings string
	Created time.Time
	Updated time.Time
}

var ErrorUserNotExist = errors.New("Not Exist")
var DB *sql.DB

func Connect() (db *sql.DB, err error) { 
	db, err = sql.Open("mysql", "dbms:1234@tcp(localhost:3309)/EFRS")
	// db, err = sql.Open("sqlite", "./../db/data/sqlite.db")
	CheckErr(err)
    // global db connection object
    DB = db

	return db, err
}

func CreateTable(db *sql.DB) {
	// create table
    sql_table := `
    CREATE TABLE IF NOT EXISTS notification(
        user_id INTEGER PRIMARY KEY,
        token VARCHAR(255) NULL,
        device VARCHAR(15) NULL,
        settings VARCHAR(15) DEFAULT '1,1,1,1,1,1,1',
        created DATE DEFAULT CURRENT_TIMESTAMP,
        updated DATE DEFAULT CURRENT_TIMESTAMP
    );
	`
	db.Exec(sql_table)
}

func FindAll(db *sql.DB) []User {
    rows, err := db.Query("SELECT * FROM notification")
    // defer rows.Close()
    var u User
    var userList []User
    for rows.Next() {
        err = rows.Scan(&u.User_id, &u.Token, &u.Device, &u.Settings, &u.Created, &u.Updated)
        CheckErr(err)
        userList = append(userList, u)
        // fmt.Printf("user_id: %d, token: %s, settings: %s\n", u.User_id, u.Token, u.Settings)
    }
    return userList
}

func FindOne(db *sql.DB, user_id uint32) (user *User, err error) {

    result := db.QueryRow("SELECT * FROM notification WHERE user_id=?", user_id)
    fmt.Println("/db/FindOne [Scan]")
    var u = new(User)
    err = result.Scan(&u.User_id, &u.Token, &u.Device, &u.Settings, &u.Created, &u.Updated)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrorUserNotExist
        }
    }
    fmt.Println("/db/FindOne [Finish]")
    return u, err
}

func Insert(db *sql.DB, user *User) string {
    // insert
    stmt, err := db.Prepare("INSERT INTO notification(user_id, token, device, settings, created, updated) values(?,?,?,?,?,?)")
    CheckErr(err)
    res, err := stmt.Exec(user.User_id, user.Token, user.Device, user.Settings, time.Now(), time.Now())
    CheckErr(err)

    id, err := res.LastInsertId()
    CheckErr(err)

    fmt.Println(id)

    return "success"
}

func Update(db *sql.DB, userId int, user User) bool {
    // update
    stmt, err := db.Prepare("update notification set (token, settings, device, updated)=(?,?,?,?) where user_id=?")
    CheckErr(err)
    user.Updated = time.Now()
    res, err := stmt.Exec(user.Token, user.Settings, user.Device, user.Updated, user.User_id)
    CheckErr(err)
    
    affect, err := res.RowsAffected()
    CheckErr(err)

    return affect!=0
}

func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}