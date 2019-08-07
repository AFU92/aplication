package main

import (
	"fmt"
	"log"
	"net/http"
)

// Movie("The notebook", "2004", "Nick Cassavetes")
// }
func main() {

	fmt.Println("El servidor esta corriendo en http://localhost:8080")

	router := NewRouter()

	server := http.ListenAndServe(":8080", router)

	log.Fatal(server)

}
