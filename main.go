package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id_user    int    `json:"id_user"`
	Pseudo     string `json:"pseudo"`
	Email      string `json:"email"`
	Passwd     string `json:"passwd"`
	Id_imagepp int    `json:"id_imagepp"`
}

type Imagepp struct {
	Id_pp     int    `json:"id_pp"`
	Image_loc string `json:"image_loc"`
}

type Tags struct {
	Id_tags int    `json:"id_tags"`
	Tags    string `json:"tags"`
}

type Topics struct {
	Id_topics        int       `json:"id_topics"`
	Titre            string    `json:"titre"`
	Crea_date        time.Time `json:"crea_date"`
	Format_crea_date string    `json:"format_crea_date"`
	Id_tags          int       `json:"id_tags"`
	Id_user          int       `json:"id_uer"`
}

type Messages struct {
	Id_message        int       `json:"id_message"`
	Message           string    `json:"message"`
	Id_user           int       `json:"id_user"`
	Publi_time        time.Time `json:"publi_time"`
	Format_publi_time string    `json:"format_publi_time"`
	Id_topics         int       `json:"id_topics"`
}

func getUsers(context *gin.Context) {
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

	var users []User

	for rows.Next() {
		var user User

		err = rows.Scan(&user.Id_user, &user.Pseudo, &user.Email, &user.Passwd, &user.Id_imagepp)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	context.IndentedJSON(http.StatusOK, users)
}

func getMessages(context *gin.Context) {
	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows2, err := db.Query("SELECT * FROM messages ")
	if err != nil {
		panic(err.Error())
	}

	defer rows2.Close()

	var messages []Messages
	for rows2.Next() {
		var message Messages

		err = rows2.Scan(&message.Id_message, &message.Message, &message.Id_user, &message.Publi_time, &message.Id_topics)
		if err != nil {
			panic(err.Error())
		}

		message.Format_publi_time = message.Publi_time.Format("15:04:05 02/01/2006")
		messages = append(messages, message)
	}

	context.IndentedJSON(http.StatusOK, messages)
}

func main() {

	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/messages", getMessages)
	router.Run("localhost:9090")
}
