package db

import (
	"fmt"
	"database/sql"

	. "github.com/bebeshen/efrs/pkg/utils"
)

func FindAllEmployee(db *sql.DB) []Employee {
    rows, err := db.Query("SELECT * FROM Employee")
    // defer rows.Close()
    var e Employee
    var employeeList []Employee
    for rows.Next() {
        err = rows.Scan(&e.Id, &e.Username, &e.Password)
        CheckErr(err)
        employeeList = append(employeeList, e)
    }
    return employeeList
}

func FindEmployeeById(db *sql.DB, e_id int) (user *Employee, err error) {

    result := db.QueryRow("SELECT * FROM Employee WHERE e_id=?", e_id)
    fmt.Println("/db/FindEmployeeById [Scan]")
    var e = new(Employee)
    err = result.Scan(&e.Id, &e.Username, &e.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrorUserNotExist
        }
    }
    fmt.Println("/db/FindEmployeeById [Finish]")
    return e, err
}

func InsertEmployee(db *sql.DB, username string, password string) string {
    // insert
    stmt, err := db.Prepare("INSERT INTO Employee(username, password) values(?,?)")
    CheckErr(err)
    res, err := stmt.Exec(username, password)
    CheckErr(err)

    id, err := res.LastInsertId()
    CheckErr(err)

    fmt.Println(id)

    return "success"
}

func UpdateEmployee(db *sql.DB, e_id int, employee Employee) bool {
    // update
    stmt, err := db.Prepare("update Employee set (id, username, password)=(?,?,?) where e_id=?")
    CheckErr(err)
    res, err := stmt.Exec(employee.Id, employee.Username, employee.Password, e_id)
    CheckErr(err)
    
    affect, err := res.RowsAffected()
    CheckErr(err)

    return affect!=0
}