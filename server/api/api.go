package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/TravyTheDev/job-matching/service/upload"
	"github.com/TravyTheDev/job-matching/service/users"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4321"},
		AllowCredentials: true,
		AllowedMethods:   []string{"OPTIONS", http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
	})

	handler := c.Handler(router)

	userStore := users.NewUserStore(s.db)
	userHander := users.NewUserHandler(userStore)
	userHander.RegisterRoutes(subRouter)

	uploadHandler := upload.NewUploadHandler()
	uploadHandler.RegisterRoutes(subRouter)

	fmt.Printf("Listening on %s\n", s.addr)
	return http.ListenAndServe(s.addr, handler)
}
