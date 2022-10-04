package main

import (
	db "go-main/package/data"
	hd "go-main/package/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	db.InitRedis()

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("/go/static/"))))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	r.HandleFunc("/todo", hd.GetAllTodo).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", hd.GetApiTodo).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", hd.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/api/todo/{id}", hd.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/api/todo/{id}", hd.DeleteTodo).Methods(http.MethodDelete)
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
