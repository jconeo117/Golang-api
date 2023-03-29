package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(errorMessage))
}

func SetupRoutes(router *mux.Router) {
	usersRouter := router.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("", getUsers).Methods(http.MethodGet)
	usersRouter.HandleFunc("/{id}", getUser).Methods(http.MethodGet)
	usersRouter.HandleFunc("", postUser).Methods(http.MethodPost)
	usersRouter.HandleFunc("/{id}", deleteUsers).Methods(http.MethodDelete)

	tasksRouter := router.PathPrefix("/tasks").Subrouter()

	tasksRouter.HandleFunc("", GetTasks).Methods(http.MethodGet)
	tasksRouter.HandleFunc("/{id}", GetTask).Methods(http.MethodGet)
	tasksRouter.HandleFunc("", PostTask).Methods(http.MethodPost)
	tasksRouter.HandleFunc("", DeleteTask).Methods(http.MethodDelete)
}
