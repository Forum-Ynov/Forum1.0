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

func GetPps(context *gin.Context) {

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM imagepp ")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var pps []models.Imagepp

	for rows.Next() {
		var pp models.Imagepp

		err = rows.Scan(&pp.Id_pp, &pp.Image_loc)
		if err != nil {
			panic(err.Error())
		}

		pps = append(pps, pp)
	}

	context.IndentedJSON(http.StatusOK, pps)
}

func GetPpById(id string) (*models.Imagepp, error) {

	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM imagepp WHERE id_pp = '" + id + "'")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var pp models.Imagepp
	var testpp models.Imagepp

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

func GetPp(context *gin.Context) {

	id := context.Param("id")
	pp, err := GetPpById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "pp not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, pp)
}
