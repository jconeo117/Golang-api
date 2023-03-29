package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jconeo117/Go_RestApi/db"
	"github.com/jconeo117/Go_RestApi/models"
	"github.com/jconeo117/Go_RestApi/routes"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	routes.SetupRoutes(router)

	http.ListenAndServe(":4000", router)
}
