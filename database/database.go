package database

import (
	"database/sql"
	"fmt"
	"kykurniawan/go-restful-api/helper"
	"os"
	"time"
)

func NewDB() *sql.DB {
	mysqlUser := os.Getenv("MYSQL_USERNAME")
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOSTNAME")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDB := os.Getenv("MYSQL_DATABASE")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDB))

	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
