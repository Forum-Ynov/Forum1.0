package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id_user    int    `json:"id_user"`
	Pseudo     string `json:"pseudo"`
	Email      string `json:"email"`
	Passwd     string `json:"passwd"`
	Id_imagepp int    `json:"id_imagepp"`
	Theme      string `json:"theme"`
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
	Description      string    `json:"description"`
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func loginuser(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	var newUser User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	user, err := getUserByPseudo(newUser.Pseudo)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	} else {
		if !CheckPasswordHash(newUser.Passwd, user.Passwd) {
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "password incorect"})
			return
		}
	}

	context.IndentedJSON(http.StatusOK, user)
}

func getUsers(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

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

		err = rows.Scan(&user.Id_user, &user.Pseudo, &user.Email, &user.Passwd, &user.Id_imagepp, &user.Theme)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	context.IndentedJSON(http.StatusOK, users)
}

func getUserById(id string) (*User, error) {

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user WHERE id_user = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var user User
	var testuser User

	for rows.Next() {

		err = rows.Scan(&user.Id_user, &user.Pseudo, &user.Email, &user.Passwd, &user.Id_imagepp, &user.Theme)
		if err != nil {
			return nil, errors.New("user not found")
		}
	}

	if user == testuser {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func getUser(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	user, err := getUserById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func getUserByPseudo(pseudo string) (*User, error) {

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user WHERE pseudo = '" + pseudo + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var user User
	var testuser User

	for rows.Next() {

		err = rows.Scan(&user.Id_user, &user.Pseudo, &user.Email, &user.Passwd, &user.Id_imagepp, &user.Theme)
		if err != nil {
			return nil, errors.New("user not found")
		}
	}

	if user == testuser {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func getUserPseudo(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	pseudo := context.Param("pseudo")
	user, err := getUserByPseudo(pseudo)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func change_imagepp(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	user, err := getUserById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err := context.BindJSON(&user); err != nil {
		return
	}

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("UPDATE user SET id_imagepp = " + strconv.Itoa(user.Id_imagepp) + "  WHERE id_user = '" + strconv.Itoa(user.Id_user) + "'"); err != nil {
		fmt.Println(err)
	}

	context.IndentedJSON(http.StatusOK, user)

}

func change_theme(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	user, err := getUserById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err := context.BindJSON(&user); err != nil {
		return
	}

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("UPDATE user SET theme = '" + user.Theme + "'  WHERE id_user = '" + strconv.Itoa(user.Id_user) + "'"); err != nil {
		fmt.Println(err)
	}

	context.IndentedJSON(http.StatusOK, user)
}

func addUsers(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	var newUser User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowspseudo, err := db.Query("SELECT * FROM user WHERE pseudo = '" + newUser.Pseudo + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rowspseudo.Close()

	var user_pseudo User
	for rowspseudo.Next() {
		err = rowspseudo.Scan(&user_pseudo.Id_user, &user_pseudo.Pseudo, &user_pseudo.Email, &user_pseudo.Passwd, &user_pseudo.Id_imagepp, &user_pseudo.Theme)
		if err != nil {
			println(errors.New("user not found"))
		}

	}

	var default_user User
	if user_pseudo != default_user {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "pseudo already used "})
		return
	}

	rowsemail, err := db.Query("SELECT * FROM user WHERE email = '" + newUser.Email + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rowsemail.Close()

	var user_email User
	for rowsemail.Next() {
		err = rowsemail.Scan(&user_email.Id_user, &user_email.Pseudo, &user_email.Email, &user_email.Passwd, &user_email.Id_imagepp, &user_email.Theme)
		if err != nil {
			println(errors.New("user not found"))
		}

	}

	if user_email != default_user {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "email already used "})
		return
	}

	newUser.Passwd, _ = HashPassword(newUser.Passwd)

	if _, err := db.Exec("INSERT INTO user (pseudo, email, passwd, id_imagepp, theme) VALUES ('" + newUser.Pseudo + "', '" + newUser.Email + "', '" + newUser.Passwd + "', '" + strconv.Itoa(newUser.Id_imagepp) + "', '" + newUser.Theme + "')"); err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM user WHERE pseudo = '" + newUser.Pseudo + "' AND email = '" + newUser.Email + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var temp_user User
	for rows.Next() {
		err = rows.Scan(&temp_user.Id_user, &temp_user.Pseudo, &temp_user.Email, &temp_user.Passwd, &temp_user.Id_imagepp, &temp_user.Theme)
		if err != nil {
			println(errors.New("user not found"))
		}
	}

	newUser.Id_user = temp_user.Id_user

	context.IndentedJSON(http.StatusCreated, newUser)

}

func deleteUser(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	user, err := getUserById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("DELETE FROM user WHERE id_user = '" + strconv.Itoa(user.Id_user) + "'"); err != nil {
		fmt.Println(err)
	}

	getUsers(context)
}

func getPps(context *gin.Context) {

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM imagepp ")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var pps []Imagepp

	for rows.Next() {
		var pp Imagepp

		err = rows.Scan(&pp.Id_pp, &pp.Image_loc)
		if err != nil {
			panic(err.Error())
		}

		pps = append(pps, pp)
	}

	context.IndentedJSON(http.StatusOK, pps)
}

func getPpById(id string) (*Imagepp, error) {

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM imagepp WHERE id_pp = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var pp Imagepp
	var testpp Imagepp

	for rows.Next() {

		err = rows.Scan(&pp.Id_pp, &pp.Image_loc)
		if err != nil {
			return nil, errors.New("pp not found")
		}
	}

	if pp == testpp {
		return nil, errors.New("pp not found")
	}

	return &pp, nil
}

func getPp(context *gin.Context) {

	id := context.Param("id")
	pp, err := getPpById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "pp not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, pp)
}

func getTags(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_tags, tags FROM tags")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	tags := []Tags{}
	for rows.Next() {
		var tag Tags
		err := rows.Scan(&tag.Id_tags, &tag.Tags)
		if err != nil {
			panic(err.Error())
		}
		tags = append(tags, tag)
	}
	if err != nil {
		panic(err.Error())
	}

	context.IndentedJSON(http.StatusOK, tags)
}

func getTagsById(id string) (*Tags, error) {

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_tags, tags FROM tags WHERE id_tags = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var tags Tags
	var testTags Tags

	for rows.Next() {

		err = rows.Scan(&tags.Id_tags, &tags.Tags)
		if err != nil {
			return nil, errors.New("tags not found")
		}
	}

	if tags == testTags {
		return nil, errors.New("tags not found")
	}

	return &tags, nil
}

func getTag(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	tags, err := getTagsById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "tags not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, tags)
}

func getTopics(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	topics := []Topics{}
	for rows.Next() {
		var topic Topics
		err := rows.Scan(&topic.Id_topics, &topic.Titre, &topic.Crea_date, &topic.Description, &topic.Id_tags, &topic.Id_user)
		if err != nil {
			panic(err.Error())
		}
		topic.Format_crea_date = topic.Crea_date.Format("2006-01-02 15:04:05")
		topics = append(topics, topic)
	}
	if err != nil {
		panic(err.Error())
	}

	context.IndentedJSON(http.StatusOK, topics)

}

func getTopicsById(id string) (*Topics, error) {

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE id_topics = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var topics Topics
	var testTopics Topics

	for rows.Next() {

		err = rows.Scan(&topics.Id_topics, &topics.Titre, &topics.Crea_date, &topics.Description, &topics.Id_tags, &topics.Id_user)
		if err != nil {
			return nil, errors.New("topics not found")
		}
	}

	if topics == testTopics {
		return nil, errors.New("topics not found")
	}

	return &topics, nil
}

func getTopic(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	topics, err := getTopicsById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "topics not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, topics)
}

func addTopic(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	var newTopic Topics

	if err := context.BindJSON(&newTopic); err != nil {
		return
	}

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsTitre, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE titre = '" + newTopic.Titre + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rowsTitre.Close()

	var topic_Titre Topics
	for rowsTitre.Next() {
		err = rowsTitre.Scan(&topic_Titre.Id_topics, &topic_Titre.Titre, &topic_Titre.Crea_date, &topic_Titre.Description, &topic_Titre.Id_tags, &topic_Titre.Id_user)
		if err != nil {
			println(errors.New("Topics not found"))
		}

	}

	var default_topic Topics
	if topic_Titre != default_topic {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Topics already used "})
		return
	}

	rowsdesc, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE description = '" + newTopic.Description + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rowsdesc.Close()

	var topic_desc Topics
	for rowsdesc.Next() {
		err = rowsdesc.Scan(&topic_desc.Id_topics, &topic_desc.Titre, &topic_desc.Crea_date, &topic_desc.Description, &topic_desc.Id_tags, &topic_desc.Id_user)
		if err != nil {
			println(errors.New("topics not found"))
		}

	}

	if topic_desc != default_topic {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "desc already used "})
		return
	}

	currentTime := time.Now()
	newTopic.Crea_date = currentTime
	newTopic.Format_crea_date = newTopic.Crea_date.Format("2006-01-02 15:04:05")

	if _, err := db.Exec("INSERT INTO topics (titre, description, crea_date, id_tags, id_user) VALUES ('" + newTopic.Titre + "', '" + newTopic.Description + "',  NOW() , '" + strconv.Itoa(newTopic.Id_tags) + "', '" + strconv.Itoa(newTopic.Id_user) + "')"); err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE titre = '" + newTopic.Titre + "' AND description = '" + newTopic.Description + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var temp_topic Topics
	for rows.Next() {
		err = rows.Scan(&temp_topic.Id_topics, &temp_topic.Titre, &temp_topic.Crea_date, &temp_topic.Description, &temp_topic.Id_tags, &temp_topic.Id_user)
		if err != nil {
			println(errors.New("topic not found"))
		}

	}

	newTopic.Id_topics = temp_topic.Id_topics

	context.IndentedJSON(http.StatusCreated, newTopic)

}

func getMessages(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

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

func getMessagesById(id string) (*Messages, error) {

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM messages WHERE id_message = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var message Messages
	var testmessage Messages

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

func getMessage(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	messages, err := getMessagesById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, messages)
}

func changeMessage(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	message, err := getMessagesById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	if err := context.BindJSON(&message); err != nil {
		return
	}

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE messages SET message = "` + message.Message + `" , publi_time = NOW()  WHERE id_message = "` + strconv.Itoa(message.Id_message) + `"`); err != nil {
		fmt.Println(err)
	}

	context.IndentedJSON(http.StatusOK, message)

}

func addMessage(context *gin.Context) {

	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	var newMessage Messages

	if err := context.BindJSON(&newMessage); err != nil {
		return
	}

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsMess, err := db.Query("SELECT * FROM messages WHERE message = '" + newMessage.Message + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rowsMess.Close()

	var Message_Mess Messages
	for rowsMess.Next() {
		err = rowsMess.Scan(&Message_Mess.Id_message, &Message_Mess.Message, &Message_Mess.Id_user, &Message_Mess.Publi_time, &Message_Mess.Id_topics)
		if err != nil {
			println(errors.New("Messages not found"))
		}

	}

	var default_message Messages
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

	var temp_message Messages
	for rows.Next() {
		err = rows.Scan(&temp_message.Id_message, &temp_message.Message, &temp_message.Id_user, &temp_message.Publi_time, &temp_message.Id_topics)
		if err != nil {
			println(errors.New("message not found"))
		}

	}

	newMessage.Id_topics = temp_message.Id_topics

	context.IndentedJSON(http.StatusCreated, newMessage)

}

func deleteMessage(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	message, err := getMessagesById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	db, err := sql.Open("mysql", "sql7606458:S4G39HTa1z@tcp(sql7.freesqldatabase.com:3306)/sql7606458?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("DELETE FROM messages WHERE id_message = '" + strconv.Itoa(message.Id_message) + "'"); err != nil {
		fmt.Println(err)
	}

	getMessages(context)
}

func main() {

	router := gin.Default()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT", "PATCH"},
	})

	router.GET("/pp", getPps)
	router.GET("/pp/:id", getPp)

	router.POST("/login", loginuser)
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.GET("/userpseudo/:pseudo", getUserPseudo)
	router.PATCH("/userpp/:id", change_imagepp)
	router.PATCH("/usertheme/:id", change_theme)
	router.POST("/adduser", addUsers)
	router.DELETE("/deleteuser/:id", deleteUser)

	router.GET("/tags", getTags)
	router.GET("/tags/:id", getTag)

	router.GET("/topics", getTopics)
	router.GET("/topics/:id", getTopic)
	router.POST("/addtopic", addTopic)

	router.GET("/messages", getMessages)
	router.GET("/messages/:id", getMessage)
	router.PATCH("/message/:id", changeMessage)
	router.POST("/addmessage", addMessage)
	router.DELETE("/deletemessage/:id", deleteMessage)

	handler := c.Handler(router)
	log.Fatal((http.ListenAndServe(":8000", handler)))

}
