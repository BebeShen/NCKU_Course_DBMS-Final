package main

import (
	"fmt"

	. "github.com/BebeShen/NCKU_Course_DBMS-Final/pkg/db"
)

func main()  {
	
	Connect()
	CreateTable(DB)
	fmt.Println("test")
}