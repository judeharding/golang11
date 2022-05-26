package main

// GORM = A GOLANG Object-Relational Mapper (ORM) is a code
// library that automates the transfer of data stored in
// relational database tables into objects that are more
// commonly used in application code.

import (
	"log"      // logs out any error
	"net/http" // used for routing

	"github.com/akhil/go-bookstore/pkg/routes" // importing MY/HIS routes
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" // space at the front for
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r)) // creates a server
} 
