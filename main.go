package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id_user int
	Pseudo  string
}

func main() {

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user ")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var user User

		err = rows.Scan(&user.Id_user, &user.Pseudo)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Id_user, user.Pseudo)
	}

	fmt.Println("hello")
}
