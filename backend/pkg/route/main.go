package route

import (
	"fmt"
	"net/http"
    "database/sql"
    "encoding/json"

    // . "github.com/bebeshen/efrs/pkg/utils"
	. "github.com/bebeshen/efrs/pkg/db"
)

func SetupCORS(w *http.ResponseWriter, req *http.Request) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func CheckMethod(w *http.ResponseWriter, method string, allowed string) {
    // check method
	if method != allowed {
		http.Error((*w), "404 not found.", http.StatusNotFound)
		fmt.Fprintf((*w), "Method Not Allowed")
		return
	}
}

func ParseRequestData(w *http.ResponseWriter, req *http.Request)  {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf((*w), "ParseForm() err: %v", err)
		return
	}
}

/*
	{Get} /hello
*/
func Hello(w http.ResponseWriter, req *http.Request) {
	SetupCORS(&w, req)

	ParseRequestData(&w, req)
	InsertCustomerDB(DB, req.FormValue("username"), req.FormValue("password"), req.FormValue("location"))
	customers := FindAllCustomer(DB)
	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(customers)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
	fmt.Println("hello")
}

func Query(w http.ResponseWriter, req *http.Request) {
	SetupCORS(&w, req)

	// check method
	CheckMethod(&w, req.Method, "POST")
	// parse data
	ParseRequestData(&w, req)
	
	rows, err := QueryBuilder(req.FormValue("sql_query")); 
	if err != nil {
		return 
	} 

	columnTypes, err := rows.ColumnTypes()
	count := len(columnTypes)
    finalRows := []interface{}{};
    // for _,s := range columnTypes {
    // 	fmt.Println("cols type:", s.DatabaseTypeName());
    // }
	for rows.Next() {
        scanArgs := make([]interface{}, count)
        for i, v := range columnTypes {
            switch v.DatabaseTypeName() {
            case "TEXT":
                scanArgs[i] = new(sql.NullString)
                break;
            case "BOOL":
                scanArgs[i] = new(sql.NullBool)
                break;
            case "REAL":
                scanArgs[i] = new(sql.NullFloat64)
                break;
            case "INTEGER":
                scanArgs[i] = new(sql.NullInt64)
                break;
            default:
                scanArgs[i] = new(sql.NullString)
            }
        }
        err := rows.Scan(scanArgs...)
        if err != nil {
            return 
        }

        masterData := map[string]interface{}{}

        for i, v := range columnTypes {
            if z, ok := (scanArgs[i]).(*sql.NullBool); ok  {
                masterData[v.Name()] = z.Bool
                continue;
            }
            if z, ok := (scanArgs[i]).(*sql.NullString); ok  {
                masterData[v.Name()] = z.String
                continue;
            }
            if z, ok := (scanArgs[i]).(*sql.NullInt64); ok  {
                masterData[v.Name()] = z.Int64
                continue;
            }
            if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok  {
                masterData[v.Name()] = z.Float64
                continue;
            }
            if z, ok := (scanArgs[i]).(*sql.NullInt32); ok  {
                masterData[v.Name()] = z.Int32
                continue;
            }
            masterData[v.Name()] = scanArgs[i]
        }
        finalRows = append(finalRows, masterData)
    }

	// fmt.Println(finalRows)
	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(finalRows)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}