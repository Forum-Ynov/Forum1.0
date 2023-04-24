package models

import (
	"time"
)

// User représente un utilisateur
type User struct {
	Id_user    int    `json:"id_user"`
	Pseudo     string `json:"pseudo"`
	Email      string `json:"email"`
	Passwd     string `json:"passwd"`
	Id_imagepp int    `json:"id_imagepp"`
	Theme      string `json:"theme"`
}

// Imagepp représente une image de profil
type Imagepp struct {
	Id_pp     int    `json:"id_pp"`
	Image_loc string `json:"image_loc"`
}

// Tags représente un tag
type Tags struct {
	Id_tags int    `json:"id_tags"`
	Tags    string `json:"tags"`
}

// Topics représente un topic
type Topics struct {
	Id_topics   int       `json:"id_topics"`
	Titre       string    `json:"titre"`
	Description string    `json:"description"`
	Crea_date   time.Time `json:"crea_date"`
	Id_tags     int       `json:"id_tags"`
	Id_user     int       `json:"id_user"`
}

// Messages représente un message
type Messages struct {
	Id_message int       `json:"id_message"`
	Message    string    `json:"message"`
	Id_user    int       `json:"id_user"`
	Publi_time time.Time `json:"publi_time"`
	Id_topics  int       `json:"id_topics"`
}
