package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/damedelion/rest_wallet/config"
	"github.com/gorilla/mux"
)

type Server struct {
	config *config.Server
	db     *sql.DB
	mux    *mux.Router
}

func NewServer(config *config.Server, db *sql.DB, mux *mux.Router) *Server {
	return &Server{config, db, mux}
}

func (s *Server) Run() {
	s.HandlerRegister()

	server := http.Server{
		// Addr:    net.JoinHostPort("localhost", "8080"), // unreachable outside container
		Addr:    fmt.Sprintf(":%d", s.config.Port), // i.e. ":3000" - is accessible outside the container
		Handler: s.mux,
	}

	fmt.Printf("server is listening on %s\n", server.Addr)
	server.ListenAndServe()
}
