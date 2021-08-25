package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
}

func main() {
	version, err := ioutil.ReadFile("VERSION")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	log.Printf("go-rest-api %s\n", version)
	log.Println("Starting on http://0.0.0.0:8080")
	http.ListenAndServe(":8080", r)
}
