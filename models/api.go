package models

import (
	"time"
)

type User struct {
	Id_user    int    `json:"id_user"`
	Pseudo     string `json:"pseudo"`
	Email      string `json:"email"`
	Passwd     string `json:"passwd"`
	Id_imagepp int    `json:"id_imagepp"`
	Theme      string `json:"theme"`
}

type Imagepp struct {
	Id_pp     int    `json:"id_pp"`
	Image_loc string `json:"image_loc"`
}

type Tags struct {
	Id_tags int    `json:"id_tags"`
	Tags    string `json:"tags"`
}

type Topics struct {
	Id_topics        int       `json:"id_topics"`
	Titre            string    `json:"titre"`
	Description      string    `json:"description"`
	Crea_date        time.Time `json:"crea_date"`
	Format_crea_date string    `json:"format_crea_date"`
	Id_tags          int       `json:"id_tags"`
	Id_user          int       `json:"id_uer"`
}

type Messages struct {
	Id_message        int       `json:"id_message"`
	Message           string    `json:"message"`
	Id_user           int       `json:"id_user"`
	Publi_time        time.Time `json:"publi_time"`
	Format_publi_time string    `json:"format_publi_time"`
	Id_topics         int       `json:"id_topics"`
}
