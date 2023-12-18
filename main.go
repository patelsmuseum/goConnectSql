package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:netcore@123@tcp(localhost:3306)/testdb")

	if err != nil {
		fmt.Println("error in validating sql.Open")
		panic(err.Error())
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Println("error in validating dp.Ping")
		panic(err.Error())
	}

	insert, err := db.Query("INSERT INTO `testdb`.`students` (`id` , `firstname` , `lastname`) VALUES ('2' , 'Ben' , 'Ford');")
	if err != nil {
		panic(err.Error())

	}

	defer insert.Close()
	fmt.Println("Succcesfull connection to db!")
}
