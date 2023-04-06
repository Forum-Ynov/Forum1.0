package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"Forum1.0/env"
	"Forum1.0/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Loginuser(context *gin.Context) {

	var newUser models.User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	user, err := GetUserByPseudo(newUser.Pseudo)
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

func GetUsers(context *gin.Context) {

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user ")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.Id_user, &user.Pseudo, &user.Email, &user.Passwd, &user.Id_imagepp, &user.Theme)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	context.IndentedJSON(http.StatusOK, users)
}

func GetUserById(id string) (*models.User, error) {

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user WHERE id_user = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var user models.User
	var testuser models.User

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

func GetUser(context *gin.Context) {

	id := context.Param("id")
	user, err := GetUserById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func GetUserByPseudo(pseudo string) (*models.User, error) {

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user WHERE pseudo = '" + pseudo + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var user models.User
	var testuser models.User

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

func GetUserPseudo(context *gin.Context) {

	pseudo := context.Param("pseudo")
	user, err := GetUserByPseudo(pseudo)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func Change_imagepp(context *gin.Context) {

	id := context.Param("id")
	user, err := GetUserById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err := context.BindJSON(&user); err != nil {
		return
	}

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("UPDATE user SET id_imagepp = " + strconv.Itoa(user.Id_imagepp) + "  WHERE id_user = '" + strconv.Itoa(user.Id_user) + "'"); err != nil {
		fmt.Println(err)
	}

	context.IndentedJSON(http.StatusOK, user)

}

func Change_theme(context *gin.Context) {

	id := context.Param("id")
	user, err := GetUserById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err := context.BindJSON(&user); err != nil {
		return
	}

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("UPDATE user SET theme = '" + user.Theme + "'  WHERE id_user = '" + strconv.Itoa(user.Id_user) + "'"); err != nil {
		fmt.Println(err)
	}

	context.IndentedJSON(http.StatusOK, user)
}

func AddUsers(context *gin.Context) {

	var newUser models.User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowspseudo, err := db.Query("SELECT * FROM user WHERE pseudo = '" + newUser.Pseudo + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rowspseudo.Close()

	var user_pseudo models.User
	for rowspseudo.Next() {
		err = rowspseudo.Scan(&user_pseudo.Id_user, &user_pseudo.Pseudo, &user_pseudo.Email, &user_pseudo.Passwd, &user_pseudo.Id_imagepp, &user_pseudo.Theme)
		if err != nil {
			println(errors.New("user not found"))
		}

	}

	var default_user models.User
	if user_pseudo != default_user {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "pseudo already used "})
		return
	}

	rowsemail, err := db.Query("SELECT * FROM user WHERE email = '" + newUser.Email + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rowsemail.Close()

	var user_email models.User
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

	var temp_user models.User
	for rows.Next() {
		err = rows.Scan(&temp_user.Id_user, &temp_user.Pseudo, &temp_user.Email, &temp_user.Passwd, &temp_user.Id_imagepp, &temp_user.Theme)
		if err != nil {
			println(errors.New("user not found"))
		}
	}

	newUser.Id_user = temp_user.Id_user

	context.IndentedJSON(http.StatusCreated, newUser)

}

func DeleteUser(context *gin.Context) {

	id := context.Param("id")
	user, err := GetUserById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if _, err := db.Exec("DELETE FROM user WHERE id_user = '" + strconv.Itoa(user.Id_user) + "'"); err != nil {
		fmt.Println(err)
	}

	GetUsers(context)
}
