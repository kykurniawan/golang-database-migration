package router

import (
	"kykurniawan/go-restful-api/controller"
	"kykurniawan/go-restful-api/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}
		w.WriteHeader(http.StatusNoContent)
	})

	router.GET("/api/categories", categoryController.GetAll)
	router.GET("/api/categories/:categoryId", categoryController.GetById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
	// 	log.Println(r.URL.Path, err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }

	router.PanicHandler = exception.ErrorHandler

	return router
}
