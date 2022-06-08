package utils

import (
    "fmt"
	"net/http"
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
	}
}