package main

import (
	"fmt"
	"log"
	"net/http"

	"Forum1.0/controllers"
	"Forum1.0/env"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

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
	router.GET("/pp", controllers.GetPps)
	router.GET("/pp/:id", controllers.GetPp)

	// Routes pour les opérations d'authentification et de gestion des utilisateurs
	router.POST("/login", controllers.Loginuser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.GET("/userpseudo/:pseudo", controllers.GetUserPseudo)
	router.PATCH("/users/:id", controllers.UpdateUser)
	router.POST("/users", controllers.AddUsers)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// Routes pour les opérations sur les tags
	router.GET("/tags", controllers.GetTags)
	router.GET("/tags/:id", controllers.GetTag)

	// Routes pour les opérations sur les topics
	router.GET("/topics", controllers.GetTopics)
	router.GET("/topics/:id", controllers.GetTopic)
	router.GET("/topics/tags/:id_tags", controllers.GetTopicsByTags)
	router.POST("/addtopic", controllers.AddTopic)

	// Routes pour les opérations sur les messages
	router.GET("/messages", controllers.GetMessages)
	router.GET("/messages/:id", controllers.GetMessage)
	router.GET("/messages/topics/:id_topics", controllers.GetMessagesByTopics)
	router.PATCH("/message/:id", controllers.ChangeMessage)
	router.POST("/addmessage", controllers.AddMessage)
	router.DELETE("/deletemessage/:id", controllers.DeleteMessage)

	// Créer un handler avec CORS middleware et le router
	handler := c.Handler(router)

	// Définir les variables d'environnement pour l'API et la base de données
	env.SetEnv()

	// Afficher les variables d'environnement sur la console
	fmt.Println("the port : " + env.Api_port + " DB open : " + env.Sql_db)

	// Lancer le serveur HTTP avec le handler et le port de l'API à partir des variables d'environnement
	log.Fatal(http.ListenAndServe(":"+env.Api_port, handler))
}
