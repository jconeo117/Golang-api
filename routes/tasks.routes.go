package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jconeo117/Go_RestApi/db"
	"github.com/jconeo117/Go_RestApi/models"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	db.DB.Find(&tasks)
	writeJSONResponse(w, http.StatusOK, tasks)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	vars := mux.Vars(r)
	id := vars["id"]

	db.DB.Find(&task, id)

	if task.ID == 0 {
		writeErrorResponse(w, http.StatusNotFound, "user is not found")
		return
	}

	writeJSONResponse(w, http.StatusOK, task)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if result := db.DB.Create(&task); result.Error != nil {
		writeErrorResponse(w, http.StatusBadRequest, result.Error.Error())
		return
	}

	writeJSONResponse(w, http.StatusCreated, task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	vars := mux.Vars(r)
	id := vars["id"]

	db.DB.First(&task, id)

	if task.ID == 0 {
		writeErrorResponse(w, http.StatusNotFound, "user is not found")
		return
	}

	db.DB.Delete(&task)
	w.WriteHeader(http.StatusOK)
}
