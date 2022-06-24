package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	myfile *os.FileInfo
	e      error
)

func txt(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	filename := string(id) + ".txt"
	myfile, e := os.Stat(filename)
	fmt.Println(filename)
	if e != nil {
		if os.IsNotExist(e) {
			log.Println("file not found")
			fmt.Fprintf(w, string("the file you are requesting could not be found"))
		}
	} else {
		log.Println("a user has accessed ", myfile.Name())
		content, err := os.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Fprintf(w, string(content))
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/txt/{id}", txt).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
