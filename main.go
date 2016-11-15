package main

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql" 
	)

	func main() { 
	db, _ := sql.Open("mysql", "root:@/db1") 
	defer db.Close()

	var (
	id int
	name string
	)

	rows, err := db.Query("select id_mhs, name from data")
	if err != nil {
	log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
	err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
	fmt.Println(id,name)
	}

	err = rows.Err()
	if err != nil {
	log.Fatal(err)
	}
}