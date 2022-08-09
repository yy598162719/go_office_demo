package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/trinodb/trino-go-client/trino"

func main() {
	dsn := "http://user@aq2:8086?catalog=tpch&schema=sf1"
	db, _ := sql.Open("trino", dsn)
	rows, _ := db.Query("select*from orders limit 1")
	count := 0
	for rows.Next() {
		count++
		fmt.Println("dd")
	}
}
