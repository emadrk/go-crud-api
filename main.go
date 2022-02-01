package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies = []Movie{}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movier := movies[index]
			json.NewEncoder(w).Encode(movier)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// params := mux.Vars(r)
	// movies = append(movies, movies(ID: params["id"], Isbn: params["isbn"], title: params["title"], Director: params["director"]))
	movies = append(movies, movie)
	// movies = movies[-1]
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	// movies = movies[-1]
	json.NewEncoder(w).Encode(movie)

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "1222", Title: "Main Hoon Na", Director: &Director{Firstname: "Farah", Lastname: "Khan"}})
	movies = append(movies, Movie{ID: "2", Isbn: "12224", Title: "Swades", Director: &Director{Firstname: "Ram Gopal", Lastname: "Verma"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("starting port at Server 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
