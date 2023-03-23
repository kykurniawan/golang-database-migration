package main

import (
	"fmt"
	"kykurniawan/go-restful-api/helper"
	"kykurniawan/go-restful-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: authMiddleware,
	}
}

func init() {
	godotenv.Load(".env")
}

func main() {

	server := InitializeServer()

	fmt.Println("server listening on port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
