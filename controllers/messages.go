package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"Forum1.0/env"
	"Forum1.0/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetMessages(context *gin.Context) {

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows2, err := db.Query("SELECT * FROM messages ")
	if err != nil {
		panic(err.Error())
	}

	defer rows2.Close()

	var messages []models.Messages
	for rows2.Next() {
		var message models.Messages

		err = rows2.Scan(&message.Id_message, &message.Message, &message.Id_user, &message.Publi_time, &message.Id_topics)
		if err != nil {
			panic(err.Error())
		}

		message.Format_publi_time = message.Publi_time.Format("15:04:05 02/01/2006")
		messages = append(messages, message)
	}

	context.IndentedJSON(http.StatusOK, messages)
}

func GetMessagesById(id string) (*models.Messages, error) {

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM messages WHERE id_message = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var message models.Messages
	var testmessage models.Messages

	for rows.Next() {

		err = rows.Scan(&message.Id_message, &message.Message, &message.Id_user, &message.Publi_time, &message.Id_topics)
		if err != nil {
			return nil, errors.New("message not found")
		}
	}

	if message == testmessage {
		return nil, errors.New("message not found")
	}

	return &message, nil
}

func GetMessage(context *gin.Context) {

	id := context.Param("id")
	messages, err := GetMessagesById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, messages)
}

func ChangeMessage(context *gin.Context) {

	id := context.Param("id")
	message, err := GetMessagesById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	if err := context.BindJSON(&message); err != nil {
		return
	}

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE messages SET message = "` + message.Message + `" , publi_time = NOW()  WHERE id_message = "` + strconv.Itoa(message.Id_message) + `"`); err != nil {
		fmt.Println(err)
	}

	context.IndentedJSON(http.StatusOK, message)

}

func AddMessage(context *gin.Context) {

	var newMessage models.Messages

	if err := context.BindJSON(&newMessage); err != nil {
		return
	}

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsMess, err := db.Query("SELECT * FROM messages WHERE message = '" + newMessage.Message + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rowsMess.Close()

	var Message_Mess models.Messages
	for rowsMess.Next() {
		err = rowsMess.Scan(&Message_Mess.Id_message, &Message_Mess.Message, &Message_Mess.Id_user, &Message_Mess.Publi_time, &Message_Mess.Id_topics)
		if err != nil {
			println(errors.New("Messages not found"))
		}

	}

	var default_message models.Messages
	if Message_Mess != default_message {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Messages already used "})
		return
	}

	currentTime := time.Now()
	newMessage.Publi_time = currentTime
	newMessage.Format_publi_time = newMessage.Publi_time.Format("2006-01-02 15:04:05")

	if _, err := db.Exec(`INSERT INTO messages (message, id_user, publi_time, id_topics) VALUES ("` + newMessage.Message + `", ` + strconv.Itoa(newMessage.Id_user) + `",  NOW() , "` + strconv.Itoa(newMessage.Id_topics) + `")`); err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT message, id_user, publi_time, id_topics FROM messages WHERE message = '" + newMessage.Message + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var temp_message models.Messages
	for rows.Next() {
		err = rows.Scan(&temp_message.Id_message, &temp_message.Message, &temp_message.Id_user, &temp_message.Publi_time, &temp_message.Id_topics)
		if err != nil {
			println(errors.New("message not found"))
		}

	}

	newMessage.Id_topics = temp_message.Id_topics

	context.IndentedJSON(http.StatusCreated, newMessage)

}

func DeleteMessage(context *gin.Context) {

	id := context.Param("id")
	message, err := GetMessagesById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("DELETE FROM messages WHERE id_message = '" + strconv.Itoa(message.Id_message) + "'"); err != nil {
		fmt.Println(err)
	}

	GetMessages(context)
}
