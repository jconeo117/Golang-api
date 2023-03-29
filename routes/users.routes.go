package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jconeo117/Go_RestApi/db"
	"github.com/jconeo117/Go_RestApi/models"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	for i := range users {
		db.DB.Model(&users[i]).Association("Tasks").Find(&users[i].Tasks)
	}

	writeJSONResponse(w, http.StatusOK, users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User

	db.DB.First(&user, id)

	if user.ID == 0 {
		writeErrorResponse(w, http.StatusNotFound, "user is not found")
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	writeJSONResponse(w, http.StatusOK, user)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if result := db.DB.Create(&user); result.Error != nil {
		writeErrorResponse(w, http.StatusBadRequest, result.Error.Error())
		return
	}

	writeJSONResponse(w, http.StatusCreated, user)
}

func deleteUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User
	vars := mux.Vars(r)
	id := vars["id"]

	db.DB.First(&user, id)

	if user.ID == 0 {
		writeErrorResponse(w, http.StatusNotFound, "user is not found")
		return
	}

	db.DB.Delete(&user)
	w.WriteHeader(http.StatusOK)
}
