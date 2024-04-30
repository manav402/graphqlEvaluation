package database

import (
	"database/sql"
	"errors"
	"os"
	_ "github.com/lib/pq"
)

var DB *sql.DB = nil

func InitDatabase()(*sql.DB,error){
	dbConnUrl,ok := os.LookupEnv("DBCONN_URL")
	if !ok {
		return nil,errors.New("can't find db url from env file")
	}
	return sql.Open("postgres",dbConnUrl)
}