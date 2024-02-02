package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"io"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

//  r is pointer of request
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// take the data in w and encode into movie data
	// can do it like this json.NewEncoder(w).Encode(movies) or
	movieData, err := json.Marshal(&movies)
	if err != nil {
		log.Fatal("Error marshaling movie data")
	}
	w.Write(movieData)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			// append all movies before and after index to delete the movie
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// marshal data into type movies
	movieData, err := json.Marshal(&movies)
	if err != nil {
		log.Fatal("Error marshalling json")
	}
	w.Write(movieData)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			// can use encoding like this instead of marshall unmarshall
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movie Movie
	body, _ := io.ReadAll(r.Body)
	
	//unmarshal json data from body into movie 
	_ = json.Unmarshal(body, &movie)
	// create random value between 0 and argument value
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	// marshal movies back to []byte for w.Write
	movieData, _ := json.Marshal(&movies)
	w.Write(movieData)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	body, _  := io.ReadAll(r.Body)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	var movie Movie
	_ = json.Unmarshal(body, &movie)
	movie.ID = params["id"]
	movies = append(movies, movie)
	movieData, err := json.Marshal(&movie)
	if err != nil {
		log.Fatal("Error marshaling movie data")
	}
	w.Write(movieData)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID: "1",
		Isbn: "428377",
		Title: "Movie One",
		Director: &Director{Firstname: "John", Lastname: "Doe"},
	})
	movies = append(movies, Movie{
		ID: "2",
		Isbn: "542998",
		Title: "Movie Two",
		Director: &Director{Firstname: "Jane", Lastname: "Loe"},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on port :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}