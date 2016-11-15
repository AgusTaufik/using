package main 

import (
	"database/sql"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//create the database hadle, confirm driver is present
	db, _ := sql.open("mysql", "root:@/db1")
	defer db.Close()

	//connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.println("connected to:", version)
}

