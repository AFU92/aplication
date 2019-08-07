package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movies como variable global
var movies = Movies{
	Movie{"Titanic", "1997", "James Cameron"},
	Movie{"Avengers", "2012", "Anthony Rosso"},
	Movie{"The notebook", "2004", "Nick Cassavetes"},
}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func contact(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Esta es la pagina de contacto: 5777-324-334")
}

func movieList(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "Building movie list")
	json.NewEncoder(w).Encode(movies)
}

func movieShow(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	movie_id := params["id"]

	// %s sustitule la variable en el string por el %s

	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}

func movieAdd(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	log.Println(movie_data)

	movies = append(movies, movie_data)

}
