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

// GetTopics retourne la liste des topics
func GetTopics(context *gin.Context) {

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err) // Si erreur lors de l'ouverture de la connexion à la base de données, affiche une erreur et arrête le programme
	}
	defer db.Close() // Ferme la connexion à la base de données à la fin de la fonction

	// Préparation de la requête SQL
	stmt, err := db.Prepare("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics ORDER BY crea_date DESC")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Exécution de la requête SQL
	rows, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Création d'un tableau de topics
	topics := []models.Topics{}

	// Parcours des topics
	for rows.Next() {
		// Récupération des données pour chaque topic
		var topic models.Topics
		err := rows.Scan(&topic.Id_topics, &topic.Titre, &topic.Crea_date, &topic.Description, &topic.Id_tags, &topic.Id_user)
		if err != nil {
			panic(err.Error())
		}
		topics = append(topics, topic) // Ajout du topic à la liste des topics
	}
	if err != nil {
		panic(err.Error())
	}

	context.IndentedJSON(http.StatusOK, topics) // Renvoie la liste des topics sous forme de JSON
}

// GetTopicsById récupère un topic à partir de son id
func GetTopicsById(id string) (*models.Topics, error) {

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id = html.EscapeString(id)

	// Préparation de la requête SQL
	stmt, err := db.Prepare("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE id_topics = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Exécution de la requête SQL
	rows, err := stmt.Query(id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var topics models.Topics
	var testTopics models.Topics

	// Récupération des données du topic
	for rows.Next() {
		err = rows.Scan(&topics.Id_topics, &topics.Titre, &topics.Crea_date, &topics.Description, &topics.Id_tags, &topics.Id_user)
		if err != nil {
			return nil, errors.New("topics not found")
		}
	}

	// Vérification si le topic existe
	if topics == testTopics {
		return nil, errors.New("topics not found")
	}

	return &topics, nil
}

// GetTopic retourne un topic à partir de son id
func GetTopic(context *gin.Context) {

	// Récupération de l'id du topic
	id := context.Param("id")

	// Récupération du topic correspondant à l'id
	topics, err := GetTopicsById(id)
	if err != nil {
		// Si le topic n'existe pas, retourne un message d'erreur
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "topics not found"})
		return
	}

	// Retourne le topic demandé
	context.IndentedJSON(http.StatusOK, topics)
}

// GetTopicsByTags récupère les topics correspondants à un id_tags donné
func GetTopicsByTags(context *gin.Context) {
	// Récupération de l'id_tags
	id_tags := context.Param("id_tags")
	id_tags = html.EscapeString(id_tags)

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Préparation de la requête SQL
	stmt, err := db.Prepare("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE id_tags = ? ORDER BY crea_date DESC")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Exécution de la requête SQL
	rows, err := stmt.Query(id_tags)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Création d'un tableau de topics
	topics := []models.Topics{}

	// Parcours des topics
	for rows.Next() {
		// Récupération des données pour chaque topic
		var topic models.Topics
		err := rows.Scan(&topic.Id_topics, &topic.Titre, &topic.Crea_date, &topic.Description, &topic.Id_tags, &topic.Id_user)
		if err != nil {
			panic(err.Error())
		}
		topics = append(topics, topic) // Ajout du topic à la liste des topics
	}
	if err != nil {
		panic(err.Error())
	}

	if len(topics) <= 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "topic not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, topics) // Renvoie la liste des topics sous forme de JSON
}

// AddTopic permet d'ajouter un nouveau sujet à la base de données en vérifiant d'abord si le titre et la description ne sont pas déjà utilisés
func AddTopic(context *gin.Context) {

	// Récupération du nouveau sujet envoyé dans la requête JSON
	var newTopic models.Topics
	if err := context.BindJSON(&newTopic); err != nil {
		return
	}

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Préparation de la requête SQL
	stmt, err := db.Prepare("SELECT * FROM topics WHERE titre = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Exécution de la requête SQL
	rowsTitre, err := stmt.Query(html.EscapeString(newTopic.Titre))
	if err != nil {
		panic(err.Error())
	}
	defer rowsTitre.Close()

	var topic_Titre models.Topics
	for rowsTitre.Next() {
		err = rowsTitre.Scan(&topic_Titre.Id_topics, &topic_Titre.Titre, &topic_Titre.Crea_date, &topic_Titre.Description, &topic_Titre.Id_tags, &topic_Titre.Id_user)
		if err != nil {
			println(errors.New("topics not found"))
		}
	}
	var default_topic models.Topics
	if topic_Titre != default_topic {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Titre déjà utilisé"})
		return
	}

	// Vérification que la description n'est pas déjà utilisée
	// Préparation de la requête SQL
	stmts, err := db.Prepare("SELECT * FROM topics WHERE description = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmts.Close()

	// Exécution de la requête SQL
	rowsdesc, err := stmts.Query(html.EscapeString(newTopic.Description))
	if err != nil {
		panic(err.Error())
	}
	defer rowsTitre.Close()
	var topic_desc models.Topics
	for rowsdesc.Next() {
		err = rowsdesc.Scan(&topic_desc.Id_topics, &topic_desc.Titre, &topic_desc.Crea_date, &topic_desc.Description, &topic_desc.Id_tags, &topic_desc.Id_user)
		if err != nil {
			println(errors.New("topics not found"))
		}
	}
	if topic_desc != default_topic {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Description déjà utilisée"})
		return
	}

	stmt2, err := db.Prepare(`INSERT INTO topics (titre, description, crea_date, id_tags, id_user) VALUES (?, ?, NOW(), ?, ?)`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt2.Close()

	// Exécution de la requête préparée avec les paramètres de la struct newTopic
	respo, err := stmt2.Exec(html.EscapeString(newTopic.Titre), html.EscapeString(newTopic.Description), newTopic.Id_tags, newTopic.Id_user)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		id, err := respo.LastInsertId()
		if err == nil {
			newTopicForm, _ := GetMessagesById(strconv.FormatInt(id, 10))
			context.IndentedJSON(http.StatusCreated, newTopicForm)
			return
		}
	}

	// On renvoie le nouveau topic en JSON avec le code de statut HTTP 201 (Created)
	context.IndentedJSON(http.StatusCreated, newTopic)

}
