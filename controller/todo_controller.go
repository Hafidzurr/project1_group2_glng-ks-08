package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Hafidzurr/project1_group2_glng-ks-08/database"
	"github.com/Hafidzurr/project1_group2_glng-ks-08/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// @Summary Create a new Todo
// @Description Create a new Todo item
// @Produce  json
// @Param todo body models.Todo true "Todo object"
// @Success 201 {object} models.Todo
// @Router /todos [post]
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	if err := database.DB.Create(&todo).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, todo)
}

// @Summary Get all Todos
// @Description Get a list of all Todos
// @Produce json
// @Success 200 {object} []models.Todo
// @Router /todos [get]
func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	if err := database.DB.Find(&todos).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, todos)
}

// @Summary Get a Todo by ID
// @Description Get a Todo by its ID
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [get]
func GetTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	var todo models.Todo
	if err := database.DB.First(&todo, todoID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			respondWithError(w, http.StatusNotFound, "To-Do not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, todo)
}

// @Summary Update a Todo
// @Description Update a Todo by its ID
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body models.Todo true "Updated Todo object"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [put]
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	var todo models.Todo
	if err := database.DB.First(&todo, todoID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			respondWithError(w, http.StatusNotFound, "To-Do not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	if err := database.DB.Save(&todo).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, todo)
}

// @Summary Delete a Todo
// @Description Delete a Todo by its ID
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} map[string]string
// @Router /todos/{id} [delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	var todo models.Todo
	if err := database.DB.First(&todo, todoID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			respondWithError(w, http.StatusNotFound, "To-Do not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := database.DB.Delete(&todo).Error; err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// Helper functions for responding with JSON
func respondWithError(w http.ResponseWriter, code int, message string) {
	response, _ := json.Marshal(map[string]string{"error": message})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
