package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	db_host     = "DB_HOST"
	db_user     = "DB_USERNAME"
	db_password = "DB_PASSWORD"
	db_name     = "DB_NAME"
)

var (
	Client   *sql.DB
	host     = os.Getenv(db_host)
	user     = os.Getenv(db_user)
	password = os.Getenv(db_password)
	dbname   = os.Getenv(db_name)
)

func init() {
	datasourceName := fmt.Sprintf("host=%s port=5432 user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, user, password, dbname)
	var err error
	Client, err = sql.Open("postgres", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured !")

}
