package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"Forum1.0/env"
	"Forum1.0/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetTags(context *gin.Context) {
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

	rows, err := db.Query("SELECT id_tags, tags FROM tags")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	tags := []models.Tags{}
	for rows.Next() {
		var tag models.Tags
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

func GetTagsById(id string) (*models.Tags, error) {

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_tags, tags FROM tags WHERE id_tags = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var tags models.Tags
	var testTags models.Tags

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

func GetTag(context *gin.Context) {
	if context.Request.Method == "OPTIONS" {
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		context.Header("Access-Control-Allow-Origin", "*")
		context.AbortWithStatus(204)
		return
	}

	id := context.Param("id")
	tags, err := GetTagsById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "tags not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, tags)
}
