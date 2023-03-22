//go:build wireinject

// +build: wireinject

package main

import (
	"kykurniawan/go-restful-api/controller"
	"kykurniawan/go-restful-api/database"
	"kykurniawan/go-restful-api/middleware"
	"kykurniawan/go-restful-api/repository"
	"kykurniawan/go-restful-api/router"
	"kykurniawan/go-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializeServer() *http.Server {
	wire.Build(
		database.NewDB,
		validator.New,
		categorySet,
		router.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
