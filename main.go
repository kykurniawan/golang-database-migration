package main

import (
	"fmt"
	"kykurniawan/go-restful-api/controller"
	"kykurniawan/go-restful-api/database"
	"kykurniawan/go-restful-api/exception"
	"kykurniawan/go-restful-api/helper"
	"kykurniawan/go-restful-api/repository"
	"kykurniawan/go-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := database.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.GetAll)
	router.GET("/api/categories/:categoryId", categoryController.GetById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: router,
	}

	fmt.Println("server listening on port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
