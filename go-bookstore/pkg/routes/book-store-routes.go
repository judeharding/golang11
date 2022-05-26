package routes

import (
	"github.com/akhil/go-bookstore/pkg/controllers" // helps import my controllers
	"github.com/gorilla/mux"                        // helps create routes
)

// All 5 routes -- Get all, Get One, Update, Delete
var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
} 