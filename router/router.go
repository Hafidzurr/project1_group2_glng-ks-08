package router

import (
	"github.com/Hafidzurr/project1_group2_glng-ks-08/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todos", controller.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", controller.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", controller.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", controller.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controller.DeleteTodo).Methods("DELETE")

	return router
}
