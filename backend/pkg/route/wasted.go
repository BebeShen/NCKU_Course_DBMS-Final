package route

import (
	"fmt"
    // "strconv"
	"net/http"
    "encoding/json"

    // . "github.com/bebeshen/efrs/pkg/utils"
	. "github.com/bebeshen/efrs/pkg/db"
)

/*
	{GET} /getAllWasted
*/
func GetAllWasted(w http.ResponseWriter, req *http.Request)  {
	SetupCORS(&w, req)
	fmt.Println("/route/getAllWasted")

	data := FindAllWasted(DB)

	// convert object to json (byte[])
	foo_marshalled, _ := json.Marshal(data)
	// convert json into string for sending response
	fmt.Fprintf(w, string(foo_marshalled))
}