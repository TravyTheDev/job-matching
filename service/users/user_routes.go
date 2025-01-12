package users

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TravyTheDev/job-matching/service/types"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	userStore types.UserStore
}

func NewUserHandler(userStore types.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

// todo i18n error handling
func (h *UserHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *UserHandler) handleRegister(w http.ResponseWriter, r *http.Request) {
	payload := types.RegisterPayload{}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "payload error", http.StatusBadRequest)
	}

	if payload.Password != payload.PasswordConfirm {
		http.Error(w, "password doesn't match", http.StatusBadRequest)
		return
	}

	user := types.User{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
		Role:     payload.Role,
	}

	if err := h.userStore.CreateUser(user); err != nil {
		log.Print(err.Error())
	}
}
