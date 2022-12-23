package main

import (
	"fmt"
	"kykurniawan/go-restful-api/controller"
	"kykurniawan/go-restful-api/database"
	"kykurniawan/go-restful-api/helper"
	"kykurniawan/go-restful-api/middleware"
	"kykurniawan/go-restful-api/repository"
	"kykurniawan/go-restful-api/router"
	"kykurniawan/go-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := database.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := router.NewRouter(categoryController)

	server := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("server listening on port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
