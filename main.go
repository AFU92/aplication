package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("El servidor esta corriendo en http://localhost:8080")

	router.HandleFunc("/", Index)
	router.HandleFunc("/contact", contact)

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func contact(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Esta es la pagina de contacto: 5777-324-334")
}
