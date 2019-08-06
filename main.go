package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie("The notebook", "2004", "Nick Cassavetes")
// }
func main() {

	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("El servidor esta corriendo en http://localhost:8080")

	router.HandleFunc("/", index)
	router.HandleFunc("/contact", contact)
	router.HandleFunc("/movies", movieList)
	router.HandleFunc("/movies/{id}", movieShow)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func contact(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Esta es la pagina de contacto: 5777-324-334")
}

func movieList(w http.ResponseWriter, r *http.Request) {
	movies := Movies{
		Movie{"TItanic", "1997", "James Cameron"},
		Movie{"Avengers", "2012", "Anthony Rosso"},
		Movie{"The notebook", "2004", "Nick Cassavetes"},
	}

	// fmt.Fprintf(w, "Building movie list")
	json.NewEncoder(w).Encode(movies)
}

func movieShow(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	movie_id := params["id"]

	// %s sustitule la variable en el string por el %s

	fmt.Fprintf(w, "Has cargado la pelicula numero %s", movie_id)
}
