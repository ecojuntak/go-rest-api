package server

import (
	"fmt"
	"go-rest-api/root"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	host string
	port string
}

type Server interface {
	Run()
}

func NewServer(host, port string) Server {
	return &server{
		host: host,
		port: port,
	}
}

func (s server) Run() {
	version, err := ioutil.ReadFile("VERSION")
	if err != nil {
		panic(err)
	}

	r := registerHandler()

	address := fmt.Sprintf("%s:%s", s.host, s.port)

	log.Printf("go-rest-api %s\n", version)
	log.Printf("Starting on http://%s", address)
	http.ListenAndServe(address, r)
}

func registerHandler() (r *mux.Router) {
	rootHandler := root.NewRootHandler()

	r = mux.NewRouter()
	r.HandleFunc("/", rootHandler.Root).Methods("GET")

	return
}
