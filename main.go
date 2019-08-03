package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
	})
	fmt.Println("El servidor esta corriendo en http://localhost:8080")
	server := http.ListenAndServe(":8080", nil)

	log.Fatal(server)

}
