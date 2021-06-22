package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zeshanwd/go-rest-api/src/api"
)

func main() {
	var port string = "8080"

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/todos", api.ListTodos).Methods("GET")
	apiRouter.HandleFunc("/todos", api.CreateTodo).Methods("POST")
	apiRouter.HandleFunc("/todos/{id}", api.SearchTodo).Methods("GET")
	apiRouter.HandleFunc("/todos/{id}", api.DeleteTodo).Methods("DELETE")
	apiRouter.HandleFunc("/todos/{id}", api.UpdateTodo).Methods("PUT")

	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":"+port, router)

}
