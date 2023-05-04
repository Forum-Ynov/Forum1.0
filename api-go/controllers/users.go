package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"html"
	"net/http"
	"strconv"

	"API-go/env"
	"API-go/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hache un mot de passe
func HashPassword(password string) (string, error) {
	// Utilise la fonction GenerateFromPassword de la librairie bcrypt pour hacher le mot de passe
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// Retourne le hash en string et une éventuelle erreur
	return string(bytes), err
}

// CheckPasswordHash vérifie si le mot de passe fourni correspond au hash stocké
func CheckPasswordHash(password, hash string) bool {
	// Compare le hash stocké avec celui du mot de passe fourni en utilisant la fonction CompareHashAndPassword de la librairie bcrypt
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// Si err est nul, cela signifie que les mots de passe correspondent, donc on renvoie true, sinon on renvoie false
	return err == nil
}

// Loginuser permet à un utilisateur de se connecter
func Loginuser(context *gin.Context) {

	// On récupère le contenu de la requête sous forme d'un nouvel utilisateur
	var newUser models.User
	if err := context.BindJSON(&newUser); err != nil {
		// Si la requête ne peut pas être interprétée comme un JSON, on retourne une erreur
		return
	}

	// On récupère l'utilisateur dans la base de données en utilisant son pseudo
	user, err := GetUserByPseudo(newUser.Pseudo)
	if err != nil {
		// Si aucun utilisateur n'a été trouvé pour ce pseudo, on retourne une erreur
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	} else {
		// On vérifie si le mot de passe fourni correspond à celui stocké dans la base de données en utilisant la fonction CheckPasswordHash
		if !CheckPasswordHash(html.EscapeString(newUser.Passwd), user.Passwd) {
			// Si les mots de passe ne correspondent pas, on retourne une erreur
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "password incorect"})
			return
		}
	}

	// Si tout s'est bien passé, on renvoie l'utilisateur sous forme de JSON
	context.IndentedJSON(http.StatusOK, user)
}

// GetUsers récupère tous les utilisateurs et les renvoie sous forme de JSON avec une réponse HTTP 200 OK.
func GetUsers(context *gin.Context) {

	// Connexion à la base de données MySQL.
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Récupération de tous les utilisateurs de la table 'user'.
	rows, err := db.Query("SELECT * FROM user ")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Stockage des résultats dans une slice de 'models.User'.
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id_user, &user.Pseudo, &user.Email, &user.Passwd, &user.Id_imagepp, &user.Theme)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	// Envoi de la réponse JSON avec les données récupérées.
	context.IndentedJSON(http.StatusOK, users)
}

// GetUserById récupère un utilisateur à partir de son identifiant
func GetUserById(id string) (*models.User, error) {

	// Ouvre une connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id = html.EscapeString(id)

	// Récupère les utilisateurs correspondant à l'identifiant
	rows, err := db.Query("SELECT * FROM user WHERE id_user = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Parcourt tous les utilisateurs récupérés
	var user models.User
	var testuser models.User

	for rows.Next() {

		// Remplit la structure utilisateur avec les données récupérées
		err = rows.Scan(&user.Id_user, &user.Pseudo, &user.Email, &user.Passwd, &user.Id_imagepp, &user.Theme)
		if err != nil {
			return nil, errors.New("user not found")
		}
	}

	// Vérifie si l'utilisateur a été trouvé
	if user == testuser {
		return nil, errors.New("user not found")
	}

	// Retourne l'utilisateur trouvé
	return &user, nil
}

// GetUser récupère un utilisateur à partir de son ID et renvoie une réponse JSON contenant les informations de l'utilisateur.
func GetUser(context *gin.Context) {
	// Récupération de l'ID de l'utilisateur depuis le paramètre de la requête
	id := context.Param("id")
	// Récupération de l'utilisateur à partir de son ID
	user, err := GetUserById(id)
	// Vérification de la présence d'une erreur lors de la récupération de l'utilisateur
	if err != nil {
		// Si une erreur est rencontrée, une réponse JSON avec le statut HTTP 404 est renvoyée
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	// Si l'utilisateur est trouvé, une réponse JSON avec les informations de l'utilisateur est renvoyée avec le statut HTTP 200
	context.IndentedJSON(http.StatusOK, user)
}

// GetUserByPseudo récupère l'utilisateur avec le pseudo spécifié dans la base de données
// et renvoie un pointeur vers cet utilisateur ainsi qu'une erreur éventuelle
func GetUserByPseudo(pseudo string) (*models.User, error) {

	// ouvrir une connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pseudo = html.EscapeString(pseudo)

	// exécuter la requête pour récupérer l'utilisateur avec le pseudo spécifié
	rows, err := db.Query("SELECT * FROM user WHERE pseudo = ?", pseudo)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var user models.User
	var testuser models.User

	// récupérer les données de l'utilisateur dans la ligne de résultat
	for rows.Next() {

		err = rows.Scan(&user.Id_user, &user.Pseudo, &user.Email, &user.Passwd, &user.Id_imagepp, &user.Theme)
		if err != nil {
			return nil, errors.New("user not found")
		}
	}

	// si l'utilisateur n'est pas trouvé, renvoyer une erreur
	if user == testuser {
		return nil, errors.New("user not found")
	}

	// renvoyer l'utilisateur et nil s'il n'y a pas d'erreur
	return &user, nil
}

// GetUserPseudo récupère l'utilisateur avec le pseudo spécifié dans la base de données
func GetUserPseudo(context *gin.Context) {

	// Récupérer le pseudo dans les paramètres de la requête
	pseudo := context.Param("pseudo")

	// Récupérer l'utilisateur en fonction de son pseudo
	user, err := GetUserByPseudo(pseudo)
	if err != nil {
		// Si l'utilisateur n'est pas trouvé, renvoyer un message d'erreur
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	// Si l'utilisateur est trouvé, renvoyer ses informations en format JSON
	context.IndentedJSON(http.StatusOK, user)
}

// Met à jour le info de l'utilisateur
func UpdateUser(context *gin.Context) {

	// Récupération de l'ID de l'utilisateur depuis le paramètre de la requête
	id := context.Param("id")
	// Récupération de l'utilisateur à partir de son ID
	user, err := GetUserById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	// Récupération des nouvelles infos à partir du corps de la requête
	if err := context.BindJSON(&user); err != nil {
		return
	}

	ispass := context.Param("passwd")

	if ispass != "" {
		user.Passwd, _ = HashPassword(html.EscapeString(user.Passwd))
	}

	// Ouverture d'une connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Exécution de la requête pour changer les infos de l'utilisateur
	if _, err := db.Exec("UPDATE user SET pseudo = ?, email = ?, passwd = ?, id_imagepp= ?, theme = ?  WHERE id_user = ?", html.EscapeString(user.Pseudo), html.EscapeString(user.Email), user.Passwd, strconv.Itoa(user.Id_imagepp), html.EscapeString(user.Theme), strconv.Itoa(user.Id_user)); err != nil {
		fmt.Println(err)
	}

	// Renvoyer l'utilisateur avec les infos modifié
	context.IndentedJSON(http.StatusOK, user)
}

// AddUsers permet d'ajouter un nouvel utilisateur dans la base de données
func AddUsers(context *gin.Context) {

	// Récupération du nouvel utilisateur depuis le corps de la requête
	var newUser models.User
	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	// Connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Vérification si le pseudo existe déjà
	rowspseudo, err := db.Query("SELECT * FROM user WHERE pseudo = ?", html.EscapeString(newUser.Pseudo))
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
	// Si le pseudo existe déjà, retourne une erreur
	var default_user models.User
	if user_pseudo != default_user {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "pseudo already used "})
		return
	}

	// Vérification si l'email existe déjà
	rowsemail, err := db.Query("SELECT * FROM user WHERE email = ?", html.EscapeString(newUser.Email))
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
	// Si l'email existe déjà, retourne une erreur
	if user_email != default_user {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "email already used "})
		return
	}

	// Hashage du mot de passe
	newUser.Passwd, _ = HashPassword(html.EscapeString(newUser.Passwd))

	// Insertion du nouvel utilisateur dans la base de données
	if resus, err := db.Exec("INSERT INTO user (pseudo, email, passwd, id_imagepp, theme) VALUES (?, ?, ?, ?, ?)", html.EscapeString(newUser.Pseudo), html.EscapeString(newUser.Email), newUser.Passwd, strconv.Itoa(newUser.Id_imagepp), html.EscapeString(newUser.Theme)); err != nil {
		fmt.Println(err)
	} else {
		id, err := resus.LastInsertId()
		if err == nil {
			Userform, _ := GetUserById(strconv.FormatInt(id, 10))
			context.IndentedJSON(http.StatusCreated, Userform)
			return
		}
	}

	// Retourne la réponse
	context.IndentedJSON(http.StatusCreated, newUser)

}

// // DeleteUser supprime un utilisateur de la base de données en fonction de son identifiant
func DeleteUser(context *gin.Context) {

	// Récupération de l'id de l'utilisateur dans les paramètres de l'URL
	id := context.Param("id")
	user, err := GetUserById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	// Connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Suppression de l'utilisateur en base de données
	if _, err := db.Exec("DELETE FROM user WHERE id_user = ?", strconv.Itoa(user.Id_user)); err != nil {
		fmt.Println(err)
	}

	// Récupération de la liste des utilisateurs restants en base de données
	GetUsers(context)
}
