package env

import "os"

// Variables d'environnement
var (
	Api_port string

	Host     string
	Port     string
	User     string
	Password string
	Database string
	Sql_db   string
)

// SetEnv récupère les variables d'environnement
func SetEnv() {
	Api_port = os.Getenv("API_PORT")

	Host = os.Getenv("DB_HOST")
	Port = os.Getenv("DB_PORT")
	User = os.Getenv("DB_USER")
	Password = os.Getenv("DB_PASSWORD")
	Database = os.Getenv("DB_DATABASE")

	Sql_db = User + ":" + Password + "@tcp(" + Host + ":" + Port + ")/" + Database + "?parseTime=true"
}
