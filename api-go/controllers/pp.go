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

// GetPps retourne la liste des pp
func GetPps(context *gin.Context) {

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM imagepp")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Création d'un tableau de pp
	var pps []models.Imagepp

	// Parcours des pp
	for rows.Next() {
		var pp models.Imagepp

		// Récupération des données
		err = rows.Scan(&pp.Id_pp, &pp.Image_loc)
		if err != nil {
			panic(err.Error())
		}

		// Ajout des données dans le tableau
		pps = append(pps, pp)
	}

	// Envoi des données
	context.IndentedJSON(http.StatusOK, pps)
}

// GetPpById retourne un pp en fonction de son id
func GetPpById(id string) (*models.Imagepp, error) {

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("mysql", env.Sql_db)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id = html.EscapeString(id)

	// Récupération des pp
	stmt, err := db.Prepare("SELECT * FROM imagepp WHERE id_pp = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Création d'un tableau de pp
	var pp models.Imagepp
	var testpp models.Imagepp

	// Parcours des pp
	for rows.Next() {

		// Récupération des données
		err = rows.Scan(&pp.Id_pp, &pp.Image_loc)
		if err != nil {
			return nil, errors.New("pp not found")
		}
	}

	// Vérification de l'existence du pp
	if pp == testpp {
		return nil, errors.New("pp not found")
	}

	// Envoi des données
	return &pp, nil
}

// GetPp retourne un pp en fonction de son id
func GetPp(context *gin.Context) {

	// Récupération de l'id
	id := context.Param("id")
	pp, err := GetPpById(id)
	// Vérification de l'existence du pp
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "pp not found"})
		return
	}

	// Envoi des données
	context.IndentedJSON(http.StatusOK, pp)
}
