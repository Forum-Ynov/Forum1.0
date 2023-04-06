package main

import (
	"fmt"
	"log"
	"net/http"

	"Forum1.0/env"

	"Forum1.0/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	// load
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	router := gin.Default()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT", "PATCH"},
	})

	router.GET("/pp", controllers.GetPps)
	router.GET("/pp/:id", controllers.GetPp)

	router.POST("/login", controllers.Loginuser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.GET("/userpseudo/:pseudo", controllers.GetUserPseudo)
	router.PATCH("/userpp/:id", controllers.Change_imagepp)
	router.PATCH("/usertheme/:id", controllers.Change_theme)
	router.POST("/adduser", controllers.AddUsers)
	router.DELETE("/deleteuser/:id", controllers.DeleteUser)

	router.GET("/tags", controllers.GetTags)
	router.GET("/tags/:id", controllers.GetTag)

	router.GET("/topics", controllers.GetTopics)
	router.GET("/topics/:id", controllers.GetTopic)
	router.POST("/addtopic", controllers.AddTopic)

	router.GET("/messages", controllers.GetMessages)
	router.GET("/messages/:id", controllers.GetMessage)
	router.PATCH("/message/:id", controllers.ChangeMessage)
	router.POST("/addmessage", controllers.AddMessage)
	router.DELETE("/deletemessage/:id", controllers.DeleteMessage)

	handler := c.Handler(router)
	env.SetEnv()
	fmt.Println("the port : " + env.Api_port + " DB open : " + env.Sql_db)
	log.Fatal(http.ListenAndServe(":"+env.Api_port, handler))

}
