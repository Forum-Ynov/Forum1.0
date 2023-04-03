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
	Id_topics   int       `json:"id_topics"`
	Titre       string    `json:"titre"`
	Description string    `json:"description"`
	Crea_date   time.Time `json:"crea_date"`
	Id_tags     int       `json:"id_tags"`
	Id_user     int       `json:"id_uer"`
}

type Messages struct {
	Id_message        int       `json:"id_message"`
	Message           string    `json:"message"`
	Id_user           int       `json:"id_user"`
	Publi_time        time.Time `json:"publi_time"`
	Format_publi_time string    `json:"format_publi_time"`
	Id_topics         int       `json:"id_topics"`
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
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

	router.GET("/messages", getMessages)

	handler := c.Handler(router)
	log.Fatal((http.ListenAndServe(":8000", handler)))
	print("test")
	// http.Handle("/", router)

}
