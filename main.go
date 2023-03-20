package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id_user    int
	Pseudo     string
	Email      string
	Passwd     string
	Id_imagepp int
}

type Imagepp struct {
	Id_pp     int
	Image_loc string
}

type Tags struct {
	Id_tags int
	Tags    string
}

type Topics struct {
	Id_topics        int
	Titre            string
	Crea_date        time.Time
	Format_crea_date string
	Id_tags          int
	Id_user          int
}

type Messages struct {
	Id_message        int
	Message           string
	Id_user           int
	Publi_time        time.Time
	Format_publi_time string
	Id_topics         int
}

func main() {

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
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

		err = rows.Scan(&user.Id_user, &user.Pseudo, &user.Email, &user.Passwd, &user.Id_imagepp)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user)
	}

	rows2, err := db.Query("SELECT * FROM messages ")
	if err != nil {
		panic(err.Error())
	}

	defer rows2.Close()

	for rows2.Next() {
		var message Messages

		err = rows2.Scan(&message.Id_message, &message.Message, &message.Id_user, &message.Publi_time, &message.Id_topics)
		if err != nil {
			panic(err.Error())
		}

		message.Format_publi_time = message.Publi_time.Format("15:04:05 02/01/2006")
		fmt.Println(message)
	}

	fmt.Println("Done")
}
