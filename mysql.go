package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:test@/test")
	checkErr(err)
	fmt.Println("mysql connected")
	stmt, err := db.Prepare("INSERT testtable SET firstname=?,lastname=?")
	checkErr(err)
	var fName string
	var lName string
	fName = os.Args[1]
	lName = os.Args[2]
	res, err := stmt.Exec(fName, lName)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	rows, err := db.Query("SELECT * FROM testtable")
	checkErr(err)
	for rows.Next() {
		var id int
		var firstname string
		var lastname string
		err = rows.Scan(&id, &firstname, &lastname)
		checkErr(err)
		fmt.Println(id, firstname, lastname)
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
