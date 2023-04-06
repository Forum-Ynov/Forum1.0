package env

import "os"

// Variables d'environnement
var (
	Api_port string
	Password string
	User     string
	Port     string
	Database string
	Sql_db   string
)

// SetEnv récupère les variables d'environnement
func SetEnv() {
	Api_port = os.Getenv("API_PORT")
	Password = os.Getenv("DB_PASSWORD")
	User = os.Getenv("DB_USER")
	Port = os.Getenv("DB_PORT")
	Database = os.Getenv("DB_DATABASE")
	Sql_db = User + ":" + Password + "@tcp(" + Port + ")/" + Database + "?parseTime=true"
}
