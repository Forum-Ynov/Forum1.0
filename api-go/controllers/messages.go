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
)

// GetMessages récupère tous les messages de la base de données
func GetMessages(context *gin.Context) {
	// Connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Requête SQL pour récupérer tous les messages
	rows2, err := db.Query("SELECT * FROM messages ORDER BY publi_time DESC")
	if err != nil {
		panic(err.Error())
	}

	// Fermer les lignes retournées par la requête
	defer rows2.Close()

	// Initialiser un slice de type Message
	var messages []models.Messages
	for rows2.Next() {
		var message models.Messages

		// Associer les colonnes de la base de données à des variables de type message
		err = rows2.Scan(&message.Id_message, &message.Message, &message.Id_user, &message.Publi_time, &message.Id_topics)
		if err != nil {
			panic(err.Error())
		}

		// Ajouter le message au slice de messages
		messages = append(messages, message)
	}

	// Retourner tous les messages en format JSON
	context.IndentedJSON(http.StatusOK, messages)
}

// GetMessagesById récupère un message par son ID
func GetMessagesById(id string) (*models.Messages, error) {
	// Connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id = html.EscapeString(id)

	// Requête SQL pour récupérer un message par son ID
	rows, err := db.Query("SELECT * FROM messages WHERE id_message = ?", id)
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var message models.Messages
	var testmessage models.Messages

	for rows.Next() {
		// Associer les colonnes de la base de données à des variables de type message
		err = rows.Scan(&message.Id_message, &message.Message, &message.Id_user, &message.Publi_time, &message.Id_topics)
		if err != nil {
			// Retourner une erreur si le message n'est pas trouvé
			return nil, errors.New("message not found")
		}
	}

	// Retourner une erreur si le message est vide
	if message == testmessage {
		return nil, errors.New("message not found")
	}

	// Retourner le message en format JSON
	return &message, nil
}

// GetMessage récupère un message par son ID
func GetMessage(context *gin.Context) {
	// Récupérer l'ID du message dans la requête
	id := context.Param("id")
	// Récupérer le message correspondant à l'ID
	messages, err := GetMessagesById(id)
	if err != nil {
		// Retourner une erreur si le message n'est pas trouvé
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	// Retourner le message en format JSON
	context.IndentedJSON(http.StatusOK, messages)
}

// ChangeMessage met à jour un message dans la base de données
func ChangeMessage(context *gin.Context) {

	// Récupérer l'ID du message depuis l'URL
	id := context.Param("id")

	// Appeler la fonction GetMessagesById pour récupérer le message correspondant à l'ID
	message, err := GetMessagesById(id)
	if err != nil {
		// Si la fonction renvoie une erreur, retourner un message JSON indiquant que le message n'a pas été trouvé
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	// Binder les données JSON reçues dans la requête HTTP à la variable message
	if err := context.BindJSON(&message); err != nil {
		return
	}

	// Ouvrir une connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Mettre à jour le message dans la base de données en utilisant une requête SQL
	if _, err := db.Exec("UPDATE messages SET message = ? , publi_time = NOW()  WHERE id_message = ?", message.Message, strconv.Itoa(message.Id_message)); err != nil {
		// Si la requête échoue, afficher l'erreur dans la console
		fmt.Println(err)
	}

	// Retourner le message mis à jour en tant que réponse JSON
	context.IndentedJSON(http.StatusOK, message)
}

func GetMessagesByTopics(context *gin.Context) {
	// Récupération de l'id_topics
	id_topics := context.Param("id_topics")

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Récupération des messages correspondants à l'id_topics donné
	rows, err := db.Query("SELECT * FROM messages WHERE id_topics = ? ORDER BY publi_time DESC", id_topics)
	if err != nil {
		panic(err.Error())
	}

	// Fermeture de la requête
	defer rows.Close()

	// Création d'un tableau de topics
	var messages []models.Messages
	for rows.Next() {
		var message models.Messages

		// Associer les colonnes de la base de données à des variables de type message
		err = rows.Scan(&message.Id_message, &message.Message, &message.Id_user, &message.Publi_time, &message.Id_topics)
		if err != nil {
			panic(err.Error())
		}

		// Ajouter le message au slice de messages
		messages = append(messages, message)
	}

	if len(messages) <= 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, messages) // Renvoie la liste des messages sous forme de JSON
}

// AddMessage ajoute un message dans la base de données
func AddMessage(context *gin.Context) {

	// Lecture de la nouvelle demande de message à partir du corps de la requête
	var newMessage models.Messages

	if err := context.BindJSON(&newMessage); err != nil {
		// Si la lecture échoue, renvoyer une réponse vide
		return
	}

	// Connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		// Si la connexion échoue, renvoyer une erreur
		panic(err)
	}
	defer db.Close()

	// Insertion du message dans la base de données
	if res, err := db.Exec("INSERT INTO messages (message, id_user, publi_time, id_topics) VALUES (?, ?,  NOW() , ?)", html.EscapeString(newMessage.Message), html.EscapeString(strconv.Itoa(newMessage.Id_user)), html.EscapeString(strconv.Itoa(newMessage.Id_topics))); err != nil {
		// Si l'insertion échoue, renvoyer une erreur
		fmt.Println(err)
	} else {
		id, err := res.LastInsertId()
		if err == nil {
			Messageform, _ := GetMessagesById(strconv.FormatInt(id, 10))
			context.IndentedJSON(http.StatusCreated, Messageform)
			return
		}
	}

	// Envoi du nouveau message en réponse
	context.IndentedJSON(http.StatusCreated, newMessage)

}

// DeleteMessage supprime un message en fonction de son ID
func DeleteMessage(context *gin.Context) {

	// Récupère l'ID du message à supprimer dans l'URL
	id := context.Param("id")

	// Vérifie si le message existe
	message, err := GetMessagesById(id)
	if err != nil {
		// Si le message n'existe pas, renvoie une erreur 404
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "message not found"})
		return
	}

	// Connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Supprime le message de la base de données
	if _, err := db.Exec("DELETE FROM messages WHERE id_message = ?", strconv.Itoa(message.Id_message)); err != nil {
		fmt.Println(err)
	}

	// Renvoie tous les messages après la suppression
	GetMessages(context)
}
