package main

// tipo de dato Movie que es una estructura
type Movie struct {
	Name      string `json:"name"`
	Year      string `json:"year"`
	Directror string `json:"director"`
}

// tipo de dato Movies es un array de Movie
type Movies []Movie
