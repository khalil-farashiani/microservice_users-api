package users_db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	Client   *sql.DB
	username = getEnv("my_sql_username", "root")
	password = getEnv("my_sql_password", "root")
	host     = getEnv("my_sql_host", "127.0.0.1:3306")
	schema   = getEnv("my_sql_schema", "users")
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger()
	log.Println("database successfully connected")
}
