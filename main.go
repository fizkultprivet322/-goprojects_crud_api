package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {

}

func getMovie(w http.ResponseWriter, r *http.Request) {

}

func createMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn: "000001", Title: "The Life of Pi", Director: &Director{FirstName: "Ang", LastName: "Lee"}})
	movies = append(movies, Movie{ID: "2", Isbn: "000002", Title: "Fight Club", Director: &Director{FirstName: "David", LastName: "Fincher"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/(id)", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/(id)", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/(id)", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
