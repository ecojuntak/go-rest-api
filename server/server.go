package server

import (
	"fmt"
	"go-rest-api/root"
	"go-rest-api/user"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type server struct {
	host string
	port string
	db   *gorm.DB
}

type Server interface {
	Run()
}

func NewServer(host, port string, db *gorm.DB) Server {
	return &server{
		host: host,
		port: port,
		db:   db,
	}
}

func (s server) Run() {
	version, err := ioutil.ReadFile("VERSION")
	if err != nil {
		panic(err)
	}

	r := s.registerHandler()

	address := fmt.Sprintf("%s:%s", s.host, s.port)

	log.Printf("go-rest-api %s\n", version)
	log.Printf("Starting on http://%s", address)
	http.ListenAndServe(address, r)
}

func (s server) registerHandler() (r *mux.Router) {

	userRepository := user.NewUserRepository(s.db)
	userService := user.NewUserService(userRepository)

	rootHandler := root.NewRootHandler()
	userHandler := user.NewUserHandler(userService)

	r = mux.NewRouter()
	r.HandleFunc("/", rootHandler.Root).Methods("GET")
	r.HandleFunc("/total_users", userHandler.GetTotalUser).Methods("GET")
	r.HandleFunc("/users", userHandler.Create).Methods("POST")

	return
}
