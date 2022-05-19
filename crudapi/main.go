package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json: "isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firsname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

var movies []Movie  // a slice of movies

func getMovies(w http.ResponseWriter, r *http.Request){
	fmt.Println("hello from getMovies")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	fmt.Println("hello from deleteMovie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params ["id"]{
			movies = append(movies[:index], (moviesindex + 1)...) // 45:55
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	//testing
	fmt.Println("hello from getMovie")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	//testing
	fmt.Println("hello from createMovie")
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID =strconv.Itoa (rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	//testing
	fmt.Println("hello from updateMovie")

	// set json content type
	w.Header().Set("Content-Type", "application/json")

	// params
	var params = mux.Vars(r)

	// loop over the movies range
	for index, item := range movies{
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index + 1]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID:"1", Isbn:"438227", Title:"Movie One",Director: &Director {Firstname:"John", Lastname:"Doe"}})

	movies = append(movies, Movie{
		ID:"2",
		Isbn:"222222",
		Title:"Movie Two",
		Director: &Director {Firstname:"Steve", Lastname:"Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000...\n  ")
	log.Fatal(http.ListenAndServe(":8000", r))
}





// ============================
// // This is the name of our package
// // Everything with this package name can see everything
// // else inside the same package, regardless of the file they are in
// package main

// // These are the libraries we are going to use
// // Both "fmt" and "net" are part of the Go standard library
// import (
// 	// "fmt" has methods for formatted I/O operations (like printing to the console)
// 	"fmt"
// 	// The "net/http" library has methods to implement HTTP clients and servers
// 	"net/http"

// 	"github.com/gorilla/mux"
// )



// func main() {
//         // Declare a new router
// 	r := mux.NewRouter()

// 	// This is where the router is useful, it allows us to declare methods that
// 	// this path will be valid for
// 	r.HandleFunc("/hello", handler).Methods("GET")

// 	// We can then pass our router (after declaring all our routes) to this method
// 	// (where previously, we were leaving the second argument as nil)
// 	http.ListenAndServe(":8080", r)
// }

// // "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// // as the arguments.
// func handler(w http.ResponseWriter, r *http.Request) {
// 	// For this case, we will always pipe "Hello World" into the response writer
// 	fmt.Fprintf(w, "Hello World!")
// }