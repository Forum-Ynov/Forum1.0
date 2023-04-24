package main

import (
	"fmt"
	"log"
	"net/http"

	"API-go/controllers"
	"API-go/env"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	apiName := "apiForum"

	// Charger les variables d'environnement à partir du fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erreur lors du chargement: %v", err)
	}

	// Créer un routeur Gin avec la configuration par défaut
	router := gin.Default()

	// Créer une instance de CORS Middleware avec des options configurées
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT", "PATCH"},
	})

	// Routes pour les manipulations des images de profil utilisateur
	router.GET("/"+apiName+"/pp", controllers.GetPps)
	router.GET("/"+apiName+"/pp/:id", controllers.GetPp)

	// Routes pour les opérations d'authentification et de gestion des utilisateurs
	router.POST("/"+apiName+"/login", controllers.Loginuser)
	router.GET("/"+apiName+"/users", controllers.GetUsers)
	router.GET("/"+apiName+"/users/:id", controllers.GetUser)
	router.GET("/"+apiName+"/userpseudo/:pseudo", controllers.GetUserPseudo)
	router.PATCH("/"+apiName+"/users/:id", controllers.UpdateUser)
	router.POST("/"+apiName+"/users", controllers.AddUsers)
	router.DELETE("/"+apiName+"/users/:id", controllers.DeleteUser)

	// Routes pour les opérations sur les tags
	router.GET("/"+apiName+"/tags", controllers.GetTags)
	router.GET("/"+apiName+"/tags/:id", controllers.GetTag)

	// Routes pour les opérations sur les topics
	router.GET("/"+apiName+"/topics", controllers.GetTopics)
	router.GET("/"+apiName+"/topics/:id", controllers.GetTopic)
	router.GET("/"+apiName+"/topicstags/:id_tags", controllers.GetTopicsByTags)
	router.POST("/"+apiName+"/topics", controllers.AddTopic)

	// Routes pour les opérations sur les messages
	router.GET("/"+apiName+"/messages", controllers.GetMessages)
	router.GET("/"+apiName+"/messages/:id", controllers.GetMessage)
	router.GET("/"+apiName+"/messagestopics/:id_topics", controllers.GetMessagesByTopics)
	router.PATCH("/"+apiName+"/messages/:id", controllers.ChangeMessage)
	router.POST("/"+apiName+"/messages", controllers.AddMessage)
	router.DELETE("/"+apiName+"/messages/:id", controllers.DeleteMessage)

	// Créer un handler avec CORS middleware et le router
	handler := c.Handler(router)

	// Définir les variables d'environnement pour l'API et la base de données
	env.SetEnv()

	// Afficher les variables d'environnement sur la console
	fmt.Println("the port : " + env.Api_port + " DB open : " + env.Sql_db)

	// Lancer le serveur HTTP avec le handler et le port de l'API à partir des variables d'environnement
	log.Fatal(http.ListenAndServe(":"+env.Api_port, handler))
}
