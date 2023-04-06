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

func GetTopics(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	topics := []models.Topics{}
	for rows.Next() {
		var topic models.Topics
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

func GetTopicsById(id string) (*models.Topics, error) {

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE id_topics = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var topics models.Topics
	var testTopics models.Topics

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

func GetTopic(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	topics, err := GetTopicsById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "topics not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, topics)
}

func AddTopic(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	var newTopic models.Topics

	if err := context.BindJSON(&newTopic); err != nil {
		return
	}

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Topics already used "})
		return
	}

	rowsdesc, err := db.Query("SELECT id_topics, titre, crea_date, description, id_tags, id_user FROM topics WHERE description = '" + newTopic.Description + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rowsdesc.Close()

	var topic_desc models.Topics
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

	var temp_topic models.Topics
	for rows.Next() {
		err = rows.Scan(&temp_topic.Id_topics, &temp_topic.Titre, &temp_topic.Crea_date, &temp_topic.Description, &temp_topic.Id_tags, &temp_topic.Id_user)
		if err != nil {
			println(errors.New("topic not found"))
		}

	}

	newTopic.Id_topics = temp_topic.Id_topics

	context.IndentedJSON(http.StatusCreated, newTopic)

}
