package models

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var (
	username string = os.Getenv("MYSQL_USERNAME")
	password string = os.Getenv("MYSQL_PASSWORD")
	protocol string = os.Getenv("MYSQL_PROTOCOL")
	address  string = os.Getenv("MYSQL_ADDRESS")
	dbname   string = os.Getenv("MYSQL_DBNAME")
	params   string = os.Getenv("MYSQL_PARAMS")
)

func Connet() (*sql.DB, error) {
	db, err := sql.Open("mysql", username+password+protocol+address+dbname+params)
	if err != nil {
		return nil, err
	}
	return db, nil
}
