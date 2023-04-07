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

// GetTopics retourne la liste des topics
func GetTopics(context *gin.Context) {

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err) // Si erreur lors de l'ouverture de la connexion à la base de données, affiche une erreur et arrête le programme
	}
	defer db.Close() // Ferme la connexion à la base de données à la fin de la fonction

	// Récupération des topics de la table 'topics'
	rows, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics")
	if err != nil {
		panic(err.Error()) // Si erreur lors de la récupération des topics, affiche une erreur et arrête le programme
	}

	// Fermeture de la requête
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
		// Formatage de la date de création
		topic.Format_crea_date = topic.Crea_date.Format("2006-01-02 15:04:05")
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

	// Récupération du topic par son id
	rows, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE id_topics = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	// Fermeture de la requête
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

	// Vérification que le titre n'est pas déjà utilisé
	rowsTitre, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE titre = '" + newTopic.Titre + "'")
	if err != nil {
		panic(err.Error())
	}
	defer rowsTitre.Close()
	var topic_Titre models.Topics
	for rowsTitre.Next() {
		err = rowsTitre.Scan(&topic_Titre.Id_topics, &topic_Titre.Titre, &topic_Titre.Crea_date, &topic_Titre.Description, &topic_Titre.Id_tags, &topic_Titre.Id_user)
		if err != nil {
			println(errors.New("Topics not found"))
		}
	}
	var default_topic models.Topics
	if topic_Titre != default_topic {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Titre déjà utilisé"})
		return
	}

	// Vérification que la description n'est pas déjà utilisée
	rowsdesc, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE description = '" + newTopic.Description + "'")
	if err != nil {
		panic(err.Error())
	}
	defer rowsdesc.Close()
	var topic_desc models.Topics
	for rowsdesc.Next() {
		err = rowsdesc.Scan(&topic_desc.Id_topics, &topic_desc.Titre, &topic_desc.Crea_date, &topic_desc.Description, &topic_desc.Id_tags, &topic_desc.Id_user)
		if err != nil {
			println(errors.New("Topics not found"))
		}
	}
	if topic_desc != default_topic {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Description déjà utilisée"})
		return
	}

	// Ajout du sujet à la base de données
	currentTime := time.Now()
	newTopic.Crea_date = currentTime
	newTopic.Format_crea_date = newTopic.Crea_date.Format("2006-01-02 15:04:05")
	if _, err := db.Exec(`INSERT INTO topics (titre, description, crea_date, id_tags, id_user) VALUES ("` + newTopic.Titre + `", "` + newTopic.Description + `",  NOW() , ` + strconv.Itoa(newTopic.Id_tags) + `, ` + strconv.Itoa(newTopic.Id_user) + `)`); err != nil {
		fmt.Println(err)
	}

	// Récupération du sujet ajouté pour renvoyer une réponse au client
	rows, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE titre = '" + newTopic.Titre + "' AND description = '" + newTopic.Description + "'")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// On initialise une variable temporaire pour stocker les données du topic nouvellement créé
	var temp_topic models.Topics

	// Pour chaque ligne renvoyée par la requête SQL, on scanne les données et on les stocke dans temp_topic
	for rows.Next() {
		err = rows.Scan(&temp_topic.Id_topics, &temp_topic.Titre, &temp_topic.Crea_date, &temp_topic.Description, &temp_topic.Id_tags, &temp_topic.Id_user)
		if err != nil {
			println(errors.New("topic not found"))
		}
	}

	// On met à jour l'ID du nouveau topic avec celui qui vient d'être créé en base de données
	newTopic.Id_topics = temp_topic.Id_topics

	// On renvoie le nouveau topic en JSON avec le code de statut HTTP 201 (Created)
	context.IndentedJSON(http.StatusCreated, newTopic)

}
