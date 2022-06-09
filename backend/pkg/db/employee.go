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
        err = rows.Scan(&e.Id, &e.Username, &e.Password, &e.Work_for)
        CheckErr(err)
        employeeList = append(employeeList, e)
    }
    return employeeList
}

func FindEmployeeById(db *sql.DB, e_id int) (user *Employee, err error) {

    result := db.QueryRow("SELECT * FROM Employee WHERE e_id=?", e_id)
    var e = new(Employee)
    err = result.Scan(&e.Id, &e.Username, &e.Password, &e.Work_for)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrorUserNotExist
        }
    }
    return e, err
}

func InsertEmployeeDB(db *sql.DB, username string, password string, work_for int) string {
    // insert
    stmt, err := db.Prepare("INSERT INTO Employee(username, password, work_for) values(?,?,?)")
    CheckErr(err)
    res, err := stmt.Exec(username, password, work_for)
    CheckErr(err)

    id, err := res.LastInsertId()
    CheckErr(err)

    fmt.Println(id)

    return "success"
}

func UpdateEmployeeDB(db *sql.DB, employee *Employee) bool {
    // update
    stmt, err := db.Prepare("update Employee set (username, password, work_for)=(?,?,?) where e_id=?")
    CheckErr(err)
    res, err := stmt.Exec(employee.Username, employee.Password, employee.Work_for, employee.Id)
    CheckErr(err)
    
    affect, err := res.RowsAffected()
    CheckErr(err)

    return affect!=0
}

func DeleteEmployeeDB(db *sql.DB, e_id int) bool {
    // delete
    stmt, err := db.Prepare("DELETE FROM Employee WHERE e_id=?")
    CheckErr(err)
    res, err := stmt.Exec(e_id)
    CheckErr(err)
    
    affect, err := res.RowsAffected()
    CheckErr(err)

    return affect!=0
}