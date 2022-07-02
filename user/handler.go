package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type userHandler struct {
	service Service
}

type Handler interface {
	GetTotalUser(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(service Service) Handler {
	return &userHandler{
		service: service,
	}
}

func (h userHandler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error reading request body"))
		return
	}

	err = h.service.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("cannot create user: %s\n", err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created"))
}

func (h userHandler) GetTotalUser(w http.ResponseWriter, r *http.Request) {
	total, err := h.service.GetTotalUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("cannot count total user: %s\n", err.Error())))
		return
	}

	response := fmt.Sprintf(`{"total": %d}`, total)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
