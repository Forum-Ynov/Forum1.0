package controllers

import (
	"database/sql"
	"errors"
	"html"
	"net/http"

	"API-go/env"
	"API-go/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// GetTags retourne la liste des tags
func GetTags(context *gin.Context) {

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id_tags, tags FROM tags")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Création d'un tableau de tags
	tags := []models.Tags{}
	for rows.Next() {
		// Récupération des données
		var tag models.Tags
		err := rows.Scan(&tag.Id_tags, &tag.Tags)
		if err != nil {
			panic(err.Error())
		}
		// Ajout des données dans le tableau
		tags = append(tags, tag)
	}

	if err != nil {
		panic(err.Error())
	}

	// Envoi des données
	context.IndentedJSON(http.StatusOK, tags)
}

// GetTagsById retourne un tag en fonction de son id
func GetTagsById(id string) (*models.Tags, error) {

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id = html.EscapeString(id)

	// Récupération des tags
	stmt, err := db.Prepare("SELECT id_tags, tags FROM tags WHERE id_tags = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Création d'un tableau de tags
	var tags models.Tags
	var testTags models.Tags

	// Parcours des tags
	for rows.Next() {

		// Récupération des données
		err = rows.Scan(&tags.Id_tags, &tags.Tags)
		if err != nil {
			return nil, errors.New("tags not found")
		}
	}

	if tags == testTags {
		return nil, errors.New("tags not found")
	}

	// Envoi des données
	return &tags, nil
}

// GetTag retourne un tag en fonction de son id
func GetTag(context *gin.Context) {

	// Récupération de l'id
	id := context.Param("id")
	// Récupération des tags
	tags, err := GetTagsById(id)
	// Vérification de l'erreur
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "tags not found"})
		return
	}

	// Envoi des données
	context.IndentedJSON(http.StatusOK, tags)
}
